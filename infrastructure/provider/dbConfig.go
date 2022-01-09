package provider

import (
	"database/sql"
)

type DbManager interface {
	DB() *sql.DB
}
