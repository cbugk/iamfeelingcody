package db

import (
	_ "database/sql"
	"log"

	_ "github.com/glebarez/go-sqlite"
)

func CreateTable(pSqlite *Sqlite) {
	pSqlite.mutex.Lock()

	createTable, err := pSqlite.db.Prepare(`CREATE TABLE IF NOT EXISTS ? (
id INTEGER PRIMARY KEY AUTOINCREMENT,
? TEXT NOT NULL UNIQUE,	
);`)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = createTable.Exec("GithubUsers", "name")
	if err != nil {
		log.Fatal(err.Error())
	}

	pSqlite.mutex.Unlock()
}
