package data

import (
	"context"
	"database/sql"
	"errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"scv/models"
	"strings"
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

func (f FileModel) Insert(file models.File) error {
	err := file.Insert(context.Background(), f.DB, boil.Infer())
	if err != nil {
		switch {
		case strings.HasPrefix(err.Error(), "models: unable to insert into files: pq: duplicate key value violates unique constraint"):
			return ErrUniqueViolation
		default:
			return err
		}
	}

	return nil
}
