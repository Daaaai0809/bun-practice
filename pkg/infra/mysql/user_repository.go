package repository

import (
	"context"
	"sync"

	"github.com/uptrace/bun"
	"github.com/Daaaai0809/bun_prac/pkg/domain/entity"
	"github.com/Daaaai0809/bun_prac/pkg/domain/repository"
)

type UserRepository struct {
	Conn *bun.DB
	mu   sync.RWMutex
}

func NewUserRepository(conn *bun.DB) repository.IUserRepository {
	return &UserRepository{Conn: conn}
}

func (r *UserRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tx, _ := r.Conn.Begin()

	var users []entity.User
	if err := tx.NewSelect().Model(&users).Relation("Posts").Scan(ctx); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) FindById(ctx context.Context, id int64) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tx, _ := r.Conn.Begin()

	var user entity.User
	if err := tx.NewSelect().Model(&user).Where("id = ?", id).Relation("Posts").Scan(ctx); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tx, _ := r.Conn.Begin()

	_, err := tx.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tx, _ := r.Conn.Begin()

	_, err := tx.NewUpdate().Model(user).Where("id = ?", user.ID).Exec(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tx, _ := r.Conn.Begin()

	_, err := tx.NewDelete().Model((*entity.User)(nil)).Where("id = ?", id).Exec(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}