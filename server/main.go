package main

import (
	"context"
	"os"
	"net/http"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	_ "github.com/go-sql-driver/mysql"

	entity "github.com/Daaaai0809/bun_prac/pkg/domain/entity"
	repository "github.com/Daaaai0809/bun_prac/pkg/infra/mysql"
	interactor "github.com/Daaaai0809/bun_prac/pkg/usecase"
	handler "github.com/Daaaai0809/bun_prac/pkg/handler"
)

func main() {
	os.Exit(run())
}

func run() int {
	ctx := context.Background()
	db, err := sql.Open("mysql", "root:root@(db:3306)/bun_prac?parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4,utf8")
	if err != nil {
		return 1
	}
	defer db.Close()

	bunDB := bun.NewDB(db, mysqldialect.New())
	bunDB.SetMaxOpenConns(10)
	bunDB.SetMaxIdleConns(10)
	bunDB.NewCreateTable().Model((*entity.User)(nil)).IfNotExists().Exec(ctx)
	bunDB.NewCreateTable().Model((*entity.Post)(nil)).IfNotExists().ForeignKey(`"user_id" references "users" ("id") on delete cascade`).Exec(ctx)

	userRepository := repository.NewUserRepository(bunDB)
	postRepository := repository.NewPostRepository(bunDB)
	userInteractor := interactor.NewUserInteractor(userRepository)
	postInteractor := interactor.NewPostInteractor(postRepository)
	userHandler := handler.NewUserHandler(userInteractor)
	postHandler := handler.NewPostHandler(postInteractor)

	mux := http.NewServeMux()
	mux.HandleFunc("/users", userHandler.Index)
	mux.HandleFunc("/users/", userHandler.Show)
	mux.HandleFunc("/users/create", userHandler.Create)
	mux.HandleFunc("/users/update", userHandler.Update)
	mux.HandleFunc("/users/delete", userHandler.Delete)
	mux.HandleFunc("/posts", postHandler.Index)
	mux.HandleFunc("/posts/", postHandler.Show)
	mux.HandleFunc("/posts/create", postHandler.Create)
	mux.HandleFunc("/posts/update", postHandler.Update)
	mux.HandleFunc("/posts/delete", postHandler.Delete)
	
	if err := http.ListenAndServe(":8080", mux); err != nil {
		return 1
	}

	return 0
}