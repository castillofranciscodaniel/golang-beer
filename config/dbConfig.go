package config

import (
	"database/sql"
)

type DbManager interface {
	DB() *sql.DB
}
