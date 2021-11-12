package main

import (
	"context"
	"database/sql"
	"flag"
	_ "github.com/lib/pq"
	"github.com/speps/go-hashids/v2"
	"os"
	"scv/internal/data"
	"scv/internal/jsonlog"
	"sync"
	"time"
)

type config struct {
	port        int
	env         string
	hashidsSalt string
	db          struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type application struct {
	config  config
	logger  *jsonlog.Logger
	wg      sync.WaitGroup
	models  data.Models
	hashids *hashids.HashID
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://postgres:12345678@localhost/scv?sslmode=disable", "PostgreSQL DSN")
	flag.StringVar(&cfg.hashidsSalt, "hashids-salt", "unpredictable secret salt", "Hashids salt")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	defer db.Close()

	logger.PrintInfo("database connection pool established", nil)

	hd := hashids.NewData()
	hd.Salt = cfg.hashidsSalt
	hd.MinLength = 11
	h, err := hashids.NewWithData(hd)

	if err != nil {
		logger.PrintFatal(err, nil)
	}

	app := &application{
		config:  cfg,
		logger:  logger,
		models:  data.NewModels(db),
		hashids: h,
	}

	err = app.serve()

	if err != nil {
		logger.PrintFatal(err, nil)
	}
}

// The openDB() function returns a sql.DB connection pool.
func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	// Set the maximum number of open (in-use + idle) connections in the pool.
	// Passing a value less than or equal to 0 will mean there is no limit.
	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	// Set the maximum number of idle connections in the pool. Again, passing a value
	// less than or equal to 0 will mean there is no limit.
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	// Use the time.ParseDuration() function to convert the idle timeout duration string
	// to a time.Duration type.
	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	// Set the maximum idle timeout.
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
