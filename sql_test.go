package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()

	_, errsql := db.ExecContext(context, `INSERT INTO customer (id, name, email, balance, rating, married, birth_date)
	VALUES ('92592b9e-7e2f-4407-92d8-90ded70508b1', 'John', NULL, 100000, 5.0, false, NULL)`)
	if errsql != nil {
		panic(errsql)
	}

	fmt.Println("Insert Data Success")
}

func TestQuery(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()

	rows, errsql := db.QueryContext(context, "SELECT id, name FROM customer")
	if errsql != nil {
		panic(errsql)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}

func TestComplexQuery(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()
	query := "SELECT id, name, email, balance, rating, married, birth_date, created_at FROM	customer"

	rows, err := db.QueryContext(context, query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var married bool
		var birth_date sql.NullTime
		var created_at time.Time

		err := rows.Scan(&id, &name, &email, &balance, &rating, &married, &birth_date, &created_at)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		fmt.Println("Married:", married)
		if birth_date.Valid {
			fmt.Println("Birth Data:", birth_date.Time)
		}
		fmt.Println("Created At:", created_at)
	}
}

func TestParameterQuery(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "admin; #"
	password := "admin"

	context := context.Background()
	sqlQuery := "SELECT username FROM user WHERE username = ? and password = ? LIMIT 1"
	rows, err := db.QueryContext(context, sqlQuery, username, password)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()
	username := "zio"
	password := "zio123"
	sqlQuery := "INSERT INTO user (username, password) VALUES(?, ?)"
	_, err := db.ExecContext(context, sqlQuery, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Insert new user success")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()
	email := "zio@gmail.com"
	comment := "Test"
	sqlQuery := "INSERT INTO comment (email, comment) VALUES(?, ?)"
	rows, err := db.ExecContext(context, sqlQuery, email, comment)
	if err != nil {
		panic(err)
	}

	result, errs := rows.LastInsertId()
	if errs != nil {
		panic(errs)
	}

	fmt.Println("Last id:", result)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()
	sqlQuery := "INSERT INTO comment (email, comment) VALUES (?, ?)"
	statement, err := db.PrepareContext(context, sqlQuery)
	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 1; i < 10; i++ {
		email := "zio" + strconv.Itoa(i) + "@gmail.com"
		comment := "Topic: Prepare Statement, Comment ke:" + strconv.Itoa(i)

		rows, err := statement.ExecContext(context, email, comment)

		if err != nil {
			panic(err)
		}

		result, _ := rows.LastInsertId()

		fmt.Println("Id", result)
	}
}

func TestDatabaseTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	context := context.Background()
	sqlQuery := "INSERT INTO comment (email, comment) VALUES (?, ?)"

	for i := 1; i < 10; i++ {
		email := "zio" + strconv.Itoa(i) + "@gmail.com"
		comment := "Topic: Database Transaction, Comment ke:" + strconv.Itoa(i)
		_, err := tx.ExecContext(context, sqlQuery, email, comment)
		if err != nil {
			panic(err)
		}
	}

	tx.Commit()
}
