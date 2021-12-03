// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Modified by Hugh Do in 12/2021

package sanbox

import (
	"bytes"
	"context"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"text/template"
	"time"
	//"github.com/bradfitz/gomemcache/memcache"
)

const (
	maxRunTime = 2 * time.Second

	// progName is the implicit program name written to the temp
	// dir and used in compiler and vet errors.
	progName = "prog.go"
)

// Responses that contain these strings will not be cached due to
// their non-deterministic nature.
var nonCachingErrors = []string{"out of memory", "cannot allocate memory"}

type response struct {
	Errors      string  `json:"errors"`
	Events      []Event `json:"events"`
	Status      int     `json:"status"`
	IsTest      bool    `json:"isTest"`
	TestsFailed int     `json:"testsFailed"`
}

// isTestFunc tells whether fn has the type of a testing function.
func isTestFunc(fn *ast.FuncDecl) bool {
	if fn.Type.Results != nil && len(fn.Type.Results.List) > 0 ||
		fn.Type.Params.List == nil ||
		len(fn.Type.Params.List) != 1 ||
		len(fn.Type.Params.List[0].Names) > 1 {
		return false
	}
	ptr, ok := fn.Type.Params.List[0].Type.(*ast.StarExpr)
	if !ok {
		return false
	}
	// We can't easily check that the type is *testing.T
	// because we don't know how testing has been imported,
	// but at least check that it's *T or *something.T.
	if name, ok := ptr.X.(*ast.Ident); ok && name.Name == "T" {
		return true
	}
	if sel, ok := ptr.X.(*ast.SelectorExpr); ok && sel.Sel.Name == "T" {
		return true
	}
	return false
}

// isTest tells whether name looks like a test (or benchmark, according to prefix).
// It is a Test (say) if there is a character after Test that is not a lower-case letter.
// We don't want TesticularCancer.
func isTest(name, prefix string) bool {
	if !strings.HasPrefix(name, prefix) {
		return false
	}
	if len(name) == len(prefix) { // "Test" is ok
		return true
	}
	return ast.IsExported(name[len(prefix):])
}

// getTestProg returns source code that executes all valid tests and examples in src.
// If the main function is present or there are no tests or examples, it returns nil.
// getTestProg emulates the "go test" command as closely as possible.
// Benchmarks are not supported because of sandboxing.
func getTestProg(src []byte) []byte {
	fset := token.NewFileSet()
	// Early bail for most cases.
	f, err := parser.ParseFile(fset, progName, src, parser.ImportsOnly)
	if err != nil || f.Name.Name != "main" {
		return nil
	}

	// importPos stores the position to inject the "testing" import declaration, if needed.
	importPos := fset.Position(f.Name.End()).Offset

	var testingImported bool
	for _, s := range f.Imports {
		if s.Path.Value == `"testing"` && s.Name == nil {
			testingImported = true
			break
		}
	}

	// Parse everything and extract test names.
	f, err = parser.ParseFile(fset, progName, src, parser.ParseComments)
	if err != nil {
		return nil
	}

	var tests []string
	for _, d := range f.Decls {
		n, ok := d.(*ast.FuncDecl)
		if !ok {
			continue
		}
		name := n.Name.Name
		switch {
		case name == "main":
			// main declared as a method will not obstruct creation of our main function.
			if n.Recv == nil {
				return nil
			}
		case isTest(name, "Test") && isTestFunc(n):
			tests = append(tests, name)
		}
	}

	// Tests imply imported "testing" package in the code.
	// If there is no import, bail to let the compiler produce an error.
	if !testingImported && len(tests) > 0 {
		return nil
	}

	// We emulate "go test". An example with no "Output" comment is compiled,
	// but not executed. An example with no text after "Output:" is compiled,
	// executed, and expected to produce no output.
	var ex []*doc.Example
	// exNoOutput indicates whether an example with no output is found.
	// We need to compile the program containing such an example even if there are no
	// other tests or examples.
	exNoOutput := false
	for _, e := range doc.Examples(f) {
		if e.Output != "" || e.EmptyOutput {
			ex = append(ex, e)
		}
		if e.Output == "" && !e.EmptyOutput {
			exNoOutput = true
		}
	}

	if len(tests) == 0 && len(ex) == 0 && !exNoOutput {
		return nil
	}

	if !testingImported && (len(ex) > 0 || exNoOutput) {
		// In case of the program with examples and no "testing" package imported,
		// add import after "package main" without modifying line numbers.
		importDecl := []byte(`;import "testing";`)
		src = bytes.Join([][]byte{src[:importPos], importDecl, src[importPos:]}, nil)
	}

	data := struct {
		Tests    []string
		Examples []*doc.Example
	}{
		tests,
		ex,
	}
	code := new(bytes.Buffer)
	if err := testTmpl.Execute(code, data); err != nil {
		panic(err)
	}
	src = append(src, code.Bytes()...)
	return src
}

var testTmpl = template.Must(template.New("main").Parse(`
func main() {
	matchAll := func(t string, pat string) (bool, error) { return true, nil }
	tests := []testing.InternalTest{
{{range .Tests}}
		{"{{.}}", {{.}}},
{{end}}
	}
	examples := []testing.InternalExample{
{{range .Examples}}
		{"Example{{.Name}}", Example{{.Name}}, {{printf "%q" .Output}}, {{.Unordered}}},
{{end}}
	}
	testing.Main(matchAll, tests, nil, examples)
}
`))

var failedTestPattern = "--- FAIL"

// compileAndRun tries to build and run a user program.
// The output of successfully ran program is returned in *response.Events.
// If a program cannot be built or has timed out,
// *response.Errors contains an explanation for a user.
func CompileAndRun(content string) (*response, error) {
	// TODO(andybons): Add semaphore to limit number of running programs at once.
	tmpDir, err := ioutil.TempDir("", "gosandbox")
	if err != nil {
		return nil, fmt.Errorf("error creating temp directory: %v", err)
	}

	//defer os.RemoveAll(tmpDir)

	files, err := splitFiles([]byte(content))
	if err != nil {
		return &response{Errors: err.Error()}, nil
	}

	var testParam string
	var buildPkgArg = "."
	if files.Num() == 1 && len(files.Data(progName)) > 0 {
		buildPkgArg = progName
		src := files.Data(progName)
		if code := getTestProg(src); code != nil {
			testParam = "-test.v"
			files.AddFile(progName, code)
		}
	}

	if !files.Contains("go.mod") {
		files.AddFile("go.mod", []byte("module play\n"))
	}

	for f, src := range files.m {
		// Before multi-file support we required that the
		// program be in package main, so continue to do that
		// for now. But permit anything in subdirectories to have other
		// packages.
		if !strings.Contains(f, "/") {
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, f, src, parser.PackageClauseOnly)
			if err == nil && f.Name.Name != "main" {
				return &response{Errors: "package name must be main"}, nil
			}
		}

		in := filepath.Join(tmpDir, f)
		if strings.Contains(f, "/") {
			if err := os.MkdirAll(filepath.Dir(in), 0755); err != nil {
				return nil, err
			}
		}
		if err := ioutil.WriteFile(in, src, 0644); err != nil {
			return nil, fmt.Errorf("error creating temp file %q: %v", in, err)
		}
	}

	exe := filepath.Join(tmpDir, "out")
	goCache := filepath.Join(tmpDir, "gocache")
	//goCache := filepath.Join(tmpDir, "gocache")
	cmd := exec.Command("go", "build", "-o", exe)
	cmd.Dir = tmpDir
	var goPath string
	//cmd.Env = []string{"GOOS=nacl", "GOARCH=amd64p32", "GOCACHE=" + goCache}

	// Modification: Support go modules by default

	// Create a GOPATH just for modules to be downloaded
	// into GOPATH/pkg/mod.
	goPath, err = ioutil.TempDir("", "gopath")
	if err != nil {
		return nil, fmt.Errorf("error creating temp directory: %v", err)
	}
	defer os.RemoveAll(goPath)
	cmd.Env = append(cmd.Env, "GO111MODULE=on", "GOPROXY="+playgroundGoproxy())
	cmd.Args = append(cmd.Args, "-mod=mod")
	cmd.Args = append(cmd.Args, "-modcacherw")
	cmd.Args = append(cmd.Args, buildPkgArg)
	cmd.Env = append(cmd.Env, "GOPATH="+goPath)
	cmd.Env = append(cmd.Env, "GOOS="+runtime.GOOS, "GOARCH="+runtime.GOARCH)
	//cmd.Env = append(cmd.Env, "GOOS=linux", "GOARCH=arm64")
	cmd.Env = append(cmd.Env, "GOCACHE="+goCache)
	if out, err := cmd.CombinedOutput(); err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			// Return compile errors to the user.

			// Rewrite compiler errors to strip the tmpDir name.
			errs := strings.Replace(string(out), tmpDir+"/", "", -1)

			// "go build", invoked with a file name, puts this odd
			// message before any compile errors; strip it.
			errs = strings.Replace(errs, "# command-line-arguments\n", "", 1)

			return &response{Errors: errs}, nil
		}
		return nil, fmt.Errorf("error building go source: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), maxRunTime)
	defer cancel()
	exec.Command("chmod", "+x", exe).Run()
	fmt.Println(exe)
	cmd = exec.CommandContext(ctx, exe, "&>", "/dev/null", testParam)
	rec := new(Recorder)
	cmd.Stdout = rec.Stdout()
	cmd.Stderr = rec.Stderr()
	var status int
	if err := cmd.Run(); err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			// Send what was captured before the timeout.
			events, err := rec.Events()
			if err != nil {
				return nil, fmt.Errorf("error decoding events: %v", err)
			}
			return &response{Errors: "process took too long", Events: events}, nil
		}
		exitErr, ok := err.(*exec.ExitError)
		if !ok {
			return nil, fmt.Errorf("error running sandbox: %v", err)
		}
		if ws, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			status = ws.ExitStatus()
		}
	}
	events, err := rec.Events()
	if err != nil {
		return nil, fmt.Errorf("error decoding events: %v", err)
	}
	var fails int
	if testParam != "" {
		// In case of testing the TestsFailed field contains how many tests have failed.
		for _, e := range events {
			fails += strings.Count(e.Message, failedTestPattern)
		}
	}
	return &response{
		Events:      events,
		Status:      status,
		IsTest:      testParam != "",
		TestsFailed: fails,
	}, nil
}

// playgroundGoproxy returns the GOPROXY environment config the playground should use.
// It is fetched from the environment variable PLAY_GOPROXY. A missing or empty
// value for PLAY_GOPROXY returns the default value of https://proxy.golang.org.
func playgroundGoproxy() string {
	proxypath := os.Getenv("PLAY_GOPROXY")
	if proxypath != "" {
		return proxypath
	}
	return "https://proxy.golang.org"
}
