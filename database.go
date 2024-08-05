package golang_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	source := "root:@tcp(localhost:3306)/golangdatabase?parseTime=true"
	db, err := sql.Open("mysql", source)

	if err != nil {
		panic(err)
	}

	// Database pooling
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
