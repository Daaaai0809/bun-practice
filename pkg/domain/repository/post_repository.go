package repository

import (
	"context"

	"github.com/Daaaai0809/bun_prac/pkg/domain/entity"
)

type IPostRepository interface {
	FindAll(ctx context.Context) ([]entity.Post, error)
	FindById(ctx context.Context, id int64) (*entity.Post, error)
	Create(ctx context.Context, post *entity.Post) error
	Update(ctx context.Context, post *entity.Post) error
	Delete(ctx context.Context, id int64) error
}