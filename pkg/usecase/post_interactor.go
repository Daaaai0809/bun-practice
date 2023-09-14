package interactor

import (
	"context"

	"github.com/Daaaai0809/bun_prac/pkg/domain/entity"
	"github.com/Daaaai0809/bun_prac/pkg/domain/repository"
)

type IPostInteractor interface {
	FindAll(ctx context.Context) ([]entity.Post, error)
	FindById(ctx context.Context, id int64) (*entity.Post, error)
	Create(ctx context.Context, post *entity.Post, tagIds []int64) error
	Update(ctx context.Context, post *entity.Post, tagIds []int64) error
	Delete(ctx context.Context, id int64) error
}

type PostInteractor struct {
	postRepository repository.IPostRepository
}

func NewPostInteractor(postRepository repository.IPostRepository) IPostInteractor {
	return &PostInteractor{
		postRepository: postRepository,
	}
}

func (interactor *PostInteractor) FindAll(ctx context.Context) ([]entity.Post, error) {
	return interactor.postRepository.FindAll(ctx)
}

func (interactor *PostInteractor) FindById(ctx context.Context, id int64) (*entity.Post, error) {
	return interactor.postRepository.FindById(ctx, id)
}

func (interactor *PostInteractor) Create(ctx context.Context, post *entity.Post, tagIds []int64) error {
	return interactor.postRepository.Create(ctx, post, tagIds)
}

func (interactor *PostInteractor) Update(ctx context.Context, post *entity.Post, tagIds []int64) error {
	return interactor.postRepository.Update(ctx, post, tagIds)
}

func (interactor *PostInteractor) Delete(ctx context.Context, id int64) error {
	return interactor.postRepository.Delete(ctx, id)
}