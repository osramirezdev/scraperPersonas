package repository

import (
	"github.com/jmoiron/sqlx"
)

type RucRepository struct {
	*sqlx.DB
}
