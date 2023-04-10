package pkg

import (
	"database/sql"
)

type DB interface {
	GetDbName() string
	GetDB() *sql.DB
	Close()
}
