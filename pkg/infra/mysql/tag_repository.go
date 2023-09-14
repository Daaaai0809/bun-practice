package repository

import (
	"context"
	"sync"

	"github.com/uptrace/bun"
	"github.com/Daaaai0809/bun_prac/pkg/domain/entity"
	"github.com/Daaaai0809/bun_prac/pkg/domain/repository"
)

type TagRepository struct {
	db *bun.DB
	mu sync.RWMutex
}

func NewTagRepository(db *bun.DB) repository.ITagRepository {
	return &TagRepository{db: db}
}

func (r *TagRepository) FindAll(ctx context.Context) ([]entity.Tag, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var tags []entity.Tag
	err := r.db.NewSelect().Model(&tags).Relation("Posts").Scan(ctx)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (r *TagRepository) FindById(ctx context.Context, id int64) (*entity.Tag, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var tag entity.Tag
	err := r.db.NewSelect().Model(&tag).Where("id = ?", id).Relation("Posts").Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (r *TagRepository) Create(ctx context.Context, tag *entity.Tag) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tx, _ := r.db.Begin()
	if _, err := tx.NewInsert().Model(tag).Exec(ctx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *TagRepository) Update(ctx context.Context, tag *entity.Tag) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tx, _ := r.db.Begin()
	if _, err := tx.NewUpdate().Model(tag).Where("id = ?", tag.ID).Exec(ctx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *TagRepository) Delete(ctx context.Context, id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tx, _ := r.db.Begin()
	if _, err := tx.NewDelete().Model(&entity.Tag{}).Where("id = ?", id).Exec(ctx); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.NewDelete().Model(&entity.PostTags{}).Where("tag_id = ?", id).Exec(ctx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}