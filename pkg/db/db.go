package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

const schema = `
CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
	title VARCHAR(128),
	comment TEXT,
	repeat VARCHAR(128)
);
CREATE INDEX dateInd ON scheduler(date);
	`

var db *sql.DB

func Init(dbFile string) error {
	_, err := os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	db, err = sql.Open("sqlite", dbFile)

	if err != nil {
		return fmt.Errorf("Error while opening database: %w", err)
	}

	if install {
		_, err = db.Exec(schema)

		if err != nil {
			return fmt.Errorf("Error while creating database: %w", err)
		}
	}

	return nil
}