package repository

import (
	"context"

	"github.com/Daaaai0809/bun_prac/pkg/domain/entity"
)

type IUserRepository interface {
	FindAll(ctx context.Context) ([]entity.User, error)
	FindById(ctx context.Context, id int64) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int64) error
}