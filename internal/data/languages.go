package data

import (
	"database/sql"
	"scv/models"
)

type LanguageModel struct {
	DB *sql.DB
}

type Language models.Language
