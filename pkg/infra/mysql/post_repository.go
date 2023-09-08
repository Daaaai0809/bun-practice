package repository

import (
	"context"
	"sync"

	"github.com/uptrace/bun"

	"github.com/Daaaai0809/bun_prac/pkg/domain/entity"
	"github.com/Daaaai0809/bun_prac/pkg/domain/repository"
)

type PostRepository struct {
	Conn *bun.DB
	mu   sync.RWMutex
}

func NewPostRepository(conn *bun.DB) repository.IPostRepository {
	return &PostRepository{Conn: conn}
}

func (r *PostRepository) FindAll(ctx context.Context) ([]entity.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tx, _ := r.Conn.Begin()

	var posts []entity.Post
	if err := tx.NewSelect().Model(&posts).Scan(ctx); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepository) FindById(ctx context.Context, id int64) (*entity.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tx, _ := r.Conn.Begin()

	var post entity.Post
	if err := tx.NewSelect().Model(&post).Where("id = ?", id).Scan(ctx); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *PostRepository) Create(ctx context.Context, post *entity.Post) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tx, _ := r.Conn.Begin()

	_, err := tx.NewInsert().Model(post).Table("users").Exec(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) Update(ctx context.Context, post *entity.Post) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tx, _ := r.Conn.Begin()

	_, err := tx.NewUpdate().Model(post).Table("users").Where("id = ?", post.ID).Exec(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) Delete(ctx context.Context, id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tx, _ := r.Conn.Begin()

	_, err := tx.NewDelete().Model(&entity.Post{}).Where("id = ?", id).Exec(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}