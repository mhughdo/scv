// Code generated by SQLBoiler 4.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Language is an object representing the database table.
type Language struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Code      string    `boil:"code" json:"code" toml:"code" yaml:"code"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *languageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L languageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LanguageColumns = struct {
	ID        string
	Code      string
	Name      string
	CreatedAt string
}{
	ID:        "id",
	Code:      "code",
	Name:      "name",
	CreatedAt: "created_at",
}

var LanguageTableColumns = struct {
	ID        string
	Code      string
	Name      string
	CreatedAt string
}{
	ID:        "languages.id",
	Code:      "languages.code",
	Name:      "languages.name",
	CreatedAt: "languages.created_at",
}

// Generated where

var LanguageWhere = struct {
	ID        whereHelperint
	Code      whereHelperstring
	Name      whereHelperstring
	CreatedAt whereHelpertime_Time
}{
	ID:        whereHelperint{field: "\"languages\".\"id\""},
	Code:      whereHelperstring{field: "\"languages\".\"code\""},
	Name:      whereHelperstring{field: "\"languages\".\"name\""},
	CreatedAt: whereHelpertime_Time{field: "\"languages\".\"created_at\""},
}

// LanguageRels is where relationship names are stored.
var LanguageRels = struct {
	File string
}{
	File: "File",
}

// languageR is where relationships are stored.
type languageR struct {
	File *File `boil:"File" json:"File" toml:"File" yaml:"File"`
}

// NewStruct creates a new relationship struct
func (*languageR) NewStruct() *languageR {
	return &languageR{}
}

// languageL is where Load methods for each relationship are stored.
type languageL struct{}

var (
	languageAllColumns            = []string{"id", "code", "name", "created_at"}
	languageColumnsWithoutDefault = []string{"code", "name"}
	languageColumnsWithDefault    = []string{"id", "created_at"}
	languagePrimaryKeyColumns     = []string{"id"}
)

type (
	// LanguageSlice is an alias for a slice of pointers to Language.
	// This should almost always be used instead of []Language.
	LanguageSlice []*Language
	// LanguageHook is the signature for custom Language hook methods
	LanguageHook func(context.Context, boil.ContextExecutor, *Language) error

	languageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	languageType                 = reflect.TypeOf(&Language{})
	languageMapping              = queries.MakeStructMapping(languageType)
	languagePrimaryKeyMapping, _ = queries.BindMapping(languageType, languageMapping, languagePrimaryKeyColumns)
	languageInsertCacheMut       sync.RWMutex
	languageInsertCache          = make(map[string]insertCache)
	languageUpdateCacheMut       sync.RWMutex
	languageUpdateCache          = make(map[string]updateCache)
	languageUpsertCacheMut       sync.RWMutex
	languageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var languageBeforeInsertHooks []LanguageHook
var languageBeforeUpdateHooks []LanguageHook
var languageBeforeDeleteHooks []LanguageHook
var languageBeforeUpsertHooks []LanguageHook

var languageAfterInsertHooks []LanguageHook
var languageAfterSelectHooks []LanguageHook
var languageAfterUpdateHooks []LanguageHook
var languageAfterDeleteHooks []LanguageHook
var languageAfterUpsertHooks []LanguageHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Language) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range languageBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Language) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range languageBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Language) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range languageBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Language) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range languageBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Language) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range languageAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Language) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range languageAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Language) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range languageAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Language) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range languageAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Language) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range languageAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLanguageHook registers your hook function for all future operations.
func AddLanguageHook(hookPoint boil.HookPoint, languageHook LanguageHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		languageBeforeInsertHooks = append(languageBeforeInsertHooks, languageHook)
	case boil.BeforeUpdateHook:
		languageBeforeUpdateHooks = append(languageBeforeUpdateHooks, languageHook)
	case boil.BeforeDeleteHook:
		languageBeforeDeleteHooks = append(languageBeforeDeleteHooks, languageHook)
	case boil.BeforeUpsertHook:
		languageBeforeUpsertHooks = append(languageBeforeUpsertHooks, languageHook)
	case boil.AfterInsertHook:
		languageAfterInsertHooks = append(languageAfterInsertHooks, languageHook)
	case boil.AfterSelectHook:
		languageAfterSelectHooks = append(languageAfterSelectHooks, languageHook)
	case boil.AfterUpdateHook:
		languageAfterUpdateHooks = append(languageAfterUpdateHooks, languageHook)
	case boil.AfterDeleteHook:
		languageAfterDeleteHooks = append(languageAfterDeleteHooks, languageHook)
	case boil.AfterUpsertHook:
		languageAfterUpsertHooks = append(languageAfterUpsertHooks, languageHook)
	}
}

// One returns a single language record from the query.
func (q languageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Language, error) {
	o := &Language{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for languages")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Language records from the query.
func (q languageQuery) All(ctx context.Context, exec boil.ContextExecutor) (LanguageSlice, error) {
	var o []*Language

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Language slice")
	}

	if len(languageAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Language records in the query.
func (q languageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count languages rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q languageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if languages exists")
	}

	return count > 0, nil
}

// File pointed to by the foreign key.
func (o *Language) File(mods ...qm.QueryMod) fileQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"language_id\" = ?", o.ID),
	}

	queryMods = append(queryMods, mods...)

	query := Files(queryMods...)
	queries.SetFrom(query.Query, "\"files\"")

	return query
}

// LoadFile allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-1 relationship.
func (languageL) LoadFile(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLanguage interface{}, mods queries.Applicator) error {
	var slice []*Language
	var object *Language

	if singular {
		object = maybeLanguage.(*Language)
	} else {
		slice = *maybeLanguage.(*[]*Language)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &languageR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &languageR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`files`),
		qm.WhereIn(`files.language_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load File")
	}

	var resultSlice []*File
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice File")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for files")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for files")
	}

	if len(languageAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.File = foreign
		if foreign.R == nil {
			foreign.R = &fileR{}
		}
		foreign.R.Language = object
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ID == foreign.LanguageID {
				local.R.File = foreign
				if foreign.R == nil {
					foreign.R = &fileR{}
				}
				foreign.R.Language = local
				break
			}
		}
	}

	return nil
}

// SetFile of the language to the related item.
// Sets o.R.File to related.
// Adds o to related.R.Language.
func (o *Language) SetFile(ctx context.Context, exec boil.ContextExecutor, insert bool, related *File) error {
	var err error

	if insert {
		related.LanguageID = o.ID

		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"files\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"language_id"}),
			strmangle.WhereClause("\"", "\"", 2, filePrimaryKeyColumns),
		)
		values := []interface{}{o.ID, related.ID}

		if boil.IsDebug(ctx) {
			writer := boil.DebugWriterFrom(ctx)
			fmt.Fprintln(writer, updateQuery)
			fmt.Fprintln(writer, values)
		}
		if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.LanguageID = o.ID

	}

	if o.R == nil {
		o.R = &languageR{
			File: related,
		}
	} else {
		o.R.File = related
	}

	if related.R == nil {
		related.R = &fileR{
			Language: o,
		}
	} else {
		related.R.Language = o
	}
	return nil
}

// Languages retrieves all the records using an executor.
func Languages(mods ...qm.QueryMod) languageQuery {
	mods = append(mods, qm.From("\"languages\""))
	return languageQuery{NewQuery(mods...)}
}

// FindLanguage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLanguage(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Language, error) {
	languageObj := &Language{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"languages\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, languageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from languages")
	}

	if err = languageObj.doAfterSelectHooks(ctx, exec); err != nil {
		return languageObj, err
	}

	return languageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Language) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no languages provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(languageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	languageInsertCacheMut.RLock()
	cache, cached := languageInsertCache[key]
	languageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			languageAllColumns,
			languageColumnsWithDefault,
			languageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(languageType, languageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(languageType, languageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"languages\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"languages\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into languages")
	}

	if !cached {
		languageInsertCacheMut.Lock()
		languageInsertCache[key] = cache
		languageInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Language.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Language) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	languageUpdateCacheMut.RLock()
	cache, cached := languageUpdateCache[key]
	languageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			languageAllColumns,
			languagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update languages, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"languages\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, languagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(languageType, languageMapping, append(wl, languagePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update languages row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for languages")
	}

	if !cached {
		languageUpdateCacheMut.Lock()
		languageUpdateCache[key] = cache
		languageUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q languageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for languages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for languages")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LanguageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), languagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"languages\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, languagePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in language slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all language")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Language) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no languages provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(languageColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	languageUpsertCacheMut.RLock()
	cache, cached := languageUpsertCache[key]
	languageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			languageAllColumns,
			languageColumnsWithDefault,
			languageColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			languageAllColumns,
			languagePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert languages, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(languagePrimaryKeyColumns))
			copy(conflict, languagePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"languages\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(languageType, languageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(languageType, languageMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert languages")
	}

	if !cached {
		languageUpsertCacheMut.Lock()
		languageUpsertCache[key] = cache
		languageUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Language record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Language) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Language provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), languagePrimaryKeyMapping)
	sql := "DELETE FROM \"languages\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from languages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for languages")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q languageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no languageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from languages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for languages")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LanguageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(languageBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), languagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"languages\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, languagePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from language slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for languages")
	}

	if len(languageAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Language) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLanguage(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LanguageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LanguageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), languagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"languages\".* FROM \"languages\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, languagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LanguageSlice")
	}

	*o = slice

	return nil
}

// LanguageExists checks if the Language row exists.
func LanguageExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"languages\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if languages exists")
	}

	return exists, nil
}
