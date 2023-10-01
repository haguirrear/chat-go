package db

import (
	"chat/settings"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func ConnectDb() *Database {
	db, err := sql.Open("postgres", settings.GetSettings().DatabaseUrl)

	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
	}

	return &Database{db: db}
}

func (db *Database) CloseDb() {
	err := db.db.Close()
	if err != nil {
		log.Println("Error closing db conn: %v", err)
	}
}

func (db *Database) GetDb() *sql.DB {
	return db.db
}
