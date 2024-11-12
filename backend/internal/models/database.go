package models

import "database/sql"

type DB struct {
	Db *sql.DB
}
