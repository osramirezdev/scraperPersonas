package repository

import (
	"github.com/jmoiron/sqlx"
)

type DocenteRepositories struct {
	*sqlx.DB
}
