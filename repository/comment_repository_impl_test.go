package repository

import (
	"context"
	"fmt"
	"golang_database"
	"golang_database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepo := GetCommentRepository(golang_database.GetConnection())
	context := context.Background()

	comment := entity.Comment{
		Email:   "scheva@gmail.com",
		Comment: "Test Repository",
	}

	result, err := commentRepo.Insert(context, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepo := GetCommentRepository(golang_database.GetConnection())
	context := context.Background()

	result, err := commentRepo.FindById(context, 12)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindAll(t *testing.T) {
	commentRepo := GetCommentRepository(golang_database.GetConnection())
	context := context.Background()

	result, err := commentRepo.FindAll(context)
	if err != nil {
		panic(err)
	}

	for _, comment := range result {
		fmt.Println(comment)
	}
}
