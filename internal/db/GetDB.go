package db

import (
	"database/sql"
	"log"
)

func GetDB(path string) *Sqlite {
	if _, ok := dbSet[path]; !ok {
		var err error
		dbFile := new(Sqlite)
		dbFile.path = path

		if path == "" {
			dbFile.db, err = sql.Open("sqlite", ":memory:")
		} else {
			dbFile.db, err = sql.Open("sqlite", path)
		}

		if err != nil {
			log.Fatal(err.Error())
		}

		dbSet[path] = dbFile
	}

	return dbSet[path]
}
