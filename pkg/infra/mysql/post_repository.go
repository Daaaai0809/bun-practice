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
	if err := tx.NewSelect().Model(&posts).Relation("Tags").Scan(ctx); err != nil {
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
	if err := tx.NewSelect().Model(&post).Relation("Tags").Where("id = ?", id).Scan(ctx); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *PostRepository) Create(ctx context.Context, post *entity.Post, tagIds []int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tx, _ := r.Conn.Begin()

	_, err := tx.NewInsert().Model(post).Exec(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if len(tagIds) > 0 {
		rels := make([]entity.PostTags, 0, len(tagIds))

		for _, tagId := range tagIds {
			rels = append(rels, entity.PostTags{
				PostId: post.ID,
				TagId:  tagId,
			})
		}

		// Many2Many relation by bulk insert
		_, err = tx.NewInsert().Model(&rels).Exec(ctx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) Update(ctx context.Context, post *entity.Post, tagIds []int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tx, _ := r.Conn.Begin()

	_, err := tx.NewUpdate().Model(post).Where("id = ?", post.ID).Exec(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if len(tagIds) > 0 {
		rels := make([]*entity.PostTags, 0, len(tagIds))
		for _, tagId := range tagIds {
			rels = append(rels, &entity.PostTags{
				PostId: post.ID,
				TagId:  tagId,
			})
		}

		_, err := tx.NewDelete().Model((*entity.PostTags)(nil)).Where("post_id = ?", post.ID).Exec(ctx)
		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = tx.NewInsert().Model(&rels).Exec(ctx)
		if err != nil {
			tx.Rollback()
			return err
		}
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

	_, err := tx.NewDelete().Model((*entity.Post)(nil)).Where("id = ?", id).Exec(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.NewDelete().Model((*entity.PostTags)(nil)).Where("post_id = ?", id).Exec(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}