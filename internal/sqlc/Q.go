package sqlc

import (
	"context"
	"database/sql"
	_ "embed"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var q *Queries
var once sync.Once

func Q() *Queries {
	once.Do(func() {
		initDB()
	})
	return q
}

//go:embed schema.sql
var schema string

func initDB() error {
	ctx := context.Background()

	// create tables
	//db, err := sql.Open("sqlite3", path.Join(binpath.Dir(), "iamfeelingcody.sqlite"))
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return err
	}

	// create tables
	if _, err := db.ExecContext(ctx, schema); err != nil {
		return err
	}

	q = New(db)

	return nil
}
