package glob

import (
	"context"
	"database/sql"
	_ "embed"
	"path"
	"sync"

	"github.com/cbugk/iamfeelingcody/internal/binpath"
	"github.com/cbugk/iamfeelingcody/internal/sqlc"
	_ "github.com/mattn/go-sqlite3"
)

var q *sqlc.Queries
var once sync.Once

func Q() *sqlc.Queries {
	once.Do(func() {
		initDB()
	})
	return q
}

// /go:embed ../sqlc/schema.sql
var schema string

func initDB() error {
	ctx := context.Background()

	// create tables
	db, err := sql.Open("sqlite3", path.Join(binpath.Dir(), "iamfeelingcody.sqlite"))
	if err != nil {
		return err
	}

	// create tables
	if _, err := db.ExecContext(ctx, schema); err != nil {
		return err
	}

	q = sqlc.New(db)

	return nil
}
