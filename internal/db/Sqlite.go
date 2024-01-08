package db

import (
	"database/sql"
	"sync"

	_ "github.com/glebarez/go-sqlite"
)

type Sqlite struct {
	path  string
	db    *sql.DB
	mutex sync.Mutex
}
