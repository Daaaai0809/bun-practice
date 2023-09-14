package interactor

import (
	"context"

	"github.com/Daaaai0809/bun_prac/pkg/domain/entity"
	"github.com/Daaaai0809/bun_prac/pkg/domain/repository"
)

type ITagInteractor interface {
	FindAll(ctx context.Context) ([]entity.Tag, error)
	FindById(ctx context.Context, id int64) (*entity.Tag, error)
	Create(ctx context.Context, tag *entity.Tag) error
	Update(ctx context.Context, tag *entity.Tag) error
	Delete(ctx context.Context, id int64) error
}

type TagInteractor struct {
	repository repository.ITagRepository
}

func NewTagInteractor(repository repository.ITagRepository) ITagInteractor {
	return &TagInteractor{repository: repository}
}

func (i *TagInteractor) FindAll(ctx context.Context) ([]entity.Tag, error) {
	return i.repository.FindAll(ctx)
}

func (i *TagInteractor) FindById(ctx context.Context, id int64) (*entity.Tag, error) {
	return i.repository.FindById(ctx, id)
}

func (i *TagInteractor) Create(ctx context.Context, tag *entity.Tag) error {
	return i.repository.Create(ctx, tag)
}

func (i *TagInteractor) Update(ctx context.Context, tag *entity.Tag) error {
	return i.repository.Update(ctx, tag)
}

func (i *TagInteractor) Delete(ctx context.Context, id int64) error {
	return i.repository.Delete(ctx, id)
}