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

func (l LanguageModel) Get(id int) (*models.Language, error) {
	return models.FindLanguage(context.Background(), l.DB, id)
}
