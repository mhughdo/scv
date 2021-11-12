package data

import (
	"context"
	"database/sql"
	"errors"
	"scv/models"
)

type FileModel struct {
	DB *sql.DB
}

func (f FileModel) Get(hash string) (*models.File, error) {
	file, err := models.Files(models.FileWhere.Hash.EQ(hash)).One(context.Background(), f.DB)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return file, nil
}
