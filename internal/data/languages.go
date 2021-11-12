package data

import (
	"context"
	"database/sql"
	"errors"
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
	language, err := models.FindLanguage(context.Background(), l.DB, id)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return language, nil
}
