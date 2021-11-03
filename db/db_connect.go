package db

import (
	"database/sql"
	"fmt"
	"os"
)

var db *sql.DB

func ConnectToDB() (err error) {
	envPath := os.Getenv("DATABASE_URL")
	if len(envPath) == 0 {
		envPath = "sqlserver://sa:p@55DressMyTrip@localhost:1433?database=master&encrypt=disable"
	}

	db, err = sql.Open("sqlserver", envPath)
	if err != nil {
		return err
	}
	return
}

func Close() (err error) {
	if db != nil {
		db.Close()
		return nil
	}
	return fmt.Errorf("No  se puede  cerrar la conexión con la DB porque es nula")
}
