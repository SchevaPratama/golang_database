package golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestConnection(t *testing.T) {
	db, err := sql.Open("mysql", "schs6631_scheva:APZP&J3nOfkh@tcp(203.175.9.41:3306)/schs6631_godatabasetest")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
