package data

import (
	"database/sql"
	"scv/models"
)

type FileModel struct {
	DB *sql.DB
}

type File models.File
