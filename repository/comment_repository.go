package repository

import (
	"context"
	"golang_database/entity"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int16) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}
