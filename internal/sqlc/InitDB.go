package sqlc

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"path/filepath"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var once sync.Once

//go:embed schema.sql
var schema string

func initDB(path string) error {
	ctx := context.Background()

	// create tables
	//db, err := sql.Open("sqlite3", path.Join(binpath.Dir(), "iamfeelingcody.sqlite"))
	db, err := sql.Open("sqlite3", path)
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

func InitDB(path string) {
	once.Do(func() {
		if err := initDB(filepath.Clean(path)); err != nil {
			log.Panic(err.Error())
		}
	})
}
