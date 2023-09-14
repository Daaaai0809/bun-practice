package repository

import (
	"context"

	"github.com/Daaaai0809/bun_prac/pkg/domain/entity"
)

type ITagRepository interface {
	FindAll(ctx context.Context) ([]entity.Tag, error)
	FindById(ctx context.Context, id int64) (*entity.Tag, error)
	Create(ctx context.Context, tag *entity.Tag) error
	Update(ctx context.Context, tag *entity.Tag) error
	Delete(ctx context.Context, id int64) error
}