package interactor

import (
	"context"

	"github.com/Daaaai0809/bun_prac/pkg/domain/entity"
	"github.com/Daaaai0809/bun_prac/pkg/domain/repository"
)

type IUserInteractor interface {
	FindAll(ctx context.Context) ([]entity.User, error)
	FindById(ctx context.Context, id int64) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int64) error
}

type UserInteractor struct {
	userRepository repository.IUserRepository
}

func NewUserInteractor(userRepository repository.IUserRepository) IUserInteractor {
	return &UserInteractor{
		userRepository: userRepository,
	}
}

func (interactor *UserInteractor) FindAll(ctx context.Context) ([]entity.User, error) {
	return interactor.userRepository.FindAll(ctx)
}

func (interactor *UserInteractor) FindById(ctx context.Context, id int64) (*entity.User, error) {
	return interactor.userRepository.FindById(ctx, id)
}

func (interactor *UserInteractor) Create(ctx context.Context, user *entity.User) error {
	return interactor.userRepository.Create(ctx, user)
}

func (interactor *UserInteractor) Update(ctx context.Context, user *entity.User) error {
	return interactor.userRepository.Update(ctx, user)
}

func (interactor *UserInteractor) Delete(ctx context.Context, id int64) error {
	return interactor.userRepository.Delete(ctx, id)
}