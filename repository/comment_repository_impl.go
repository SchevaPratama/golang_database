package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_database/entity"
	"strconv"
)

type CommentRepositoryImplementation struct {
	DB *sql.DB
}

func GetCommentRepository(db *sql.DB) CommentRepository {
	return &CommentRepositoryImplementation{DB: db}
}

func (commentRepo *CommentRepositoryImplementation) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlQuery := "INSERT INTO comment (email, comment) VALUES(?, ?)"
	rows, err := commentRepo.DB.ExecContext(ctx, sqlQuery, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	result, errs := rows.LastInsertId()
	if errs != nil {
		return comment, errs
	}
	comment.Id = int16(result)
	return comment, nil
}

func (commentRepo *CommentRepositoryImplementation) FindById(ctx context.Context, id int16) (entity.Comment, error) {
	sqlQuery := "SELECT id, email, comment FROM comment WHERE id = ? LIMIT 1"
	rows, err := commentRepo.DB.QueryContext(ctx, sqlQuery, id)

	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}

	defer rows.Close()
	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("Id:" + strconv.Itoa(int(id)) + "not found")
	}
}

func (commentRepo *CommentRepositoryImplementation) FindAll(ctx context.Context) ([]entity.Comment, error) {
	sqlQuery := "SELECT id, email, comment FROM comment"
	result, err := commentRepo.DB.QueryContext(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	defer result.Close()
	var comments []entity.Comment
	for result.Next() {
		comment := entity.Comment{}
		result.Scan(&comment.Id, &comment.Email, &comment.Id)
		comments = append(comments, comment)
	}

	return comments, nil
}
