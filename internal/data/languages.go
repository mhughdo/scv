package data

import (
	"context"
	"database/sql"
	"scv/models"
)

type LanguageModel struct {
	DB *sql.DB
}

type Language models.Language

func (l LanguageModel) GetAll() ([]*models.Language, error) {
	return models.Languages().All(context.Background(), l.DB)
}
