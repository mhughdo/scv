package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Languages LanguageModel
	Files FileModel
}

func NewModels (db *sql.DB) Models {
	return Models{
		Files: FileModel{DB: db},
		Languages: LanguageModel{DB: db},
	}
}
