package main

import (
	"context"
	"os"
	"time"
	"log"
	"net/http"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	entity "github.com/Daaaai0809/bun_prac/pkg/domain/entity"
	repository "github.com/Daaaai0809/bun_prac/pkg/infra/mysql"
	interactor "github.com/Daaaai0809/bun_prac/pkg/usecase"
	handler "github.com/Daaaai0809/bun_prac/pkg/handler"
	infra "github.com/Daaaai0809/bun_prac/pkg/infra"
)

var mysqlConfig = mysql.Config{}

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println(err)
	}

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Println(err)
	}

	mysqlConfig = mysql.Config{
		User: os.Getenv("MYSQL_ROOT_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Addr: os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT"),
		DBName: os.Getenv("MYSQL_DATABASE"),
		Net: "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc: jst,
	}
}

func main() {
	os.Exit(run())
}

func run() int {
	ctx := context.Background()
	db, err := infra.Connection(mysqlConfig)
	if err != nil {
		log.Println(err)
		return 1
	}
	defer db.Close()

	bunDB := bun.NewDB(db, mysqldialect.New())
	bunDB.SetMaxOpenConns(10)
	bunDB.SetMaxIdleConns(10)
	_, err = bunDB.NewCreateTable().Model((*entity.PostTags)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Println(err)
		return 1
	}
	_, err = bunDB.NewCreateTable().Model((*entity.User)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Println(err)
		return 1
	}
	_, err = bunDB.NewCreateTable().Model((*entity.Post)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Println(err)
		return 1
	}
	_, err = bunDB.NewCreateTable().Model((*entity.Tag)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Println(err)
		return 1
	}

	userRepository := repository.NewUserRepository(bunDB)
	postRepository := repository.NewPostRepository(bunDB)
	tagRepository := repository.NewTagRepository(bunDB)
	userInteractor := interactor.NewUserInteractor(userRepository)
	postInteractor := interactor.NewPostInteractor(postRepository)
	tagInteractor := interactor.NewTagInteractor(tagRepository)
	userHandler := handler.NewUserHandler(userInteractor)
	postHandler := handler.NewPostHandler(postInteractor)
	tagHandler := handler.NewTagHandler(tagInteractor)

	mux := http.NewServeMux()
	mux.HandleFunc("/users", userHandler.Index)
	mux.HandleFunc("/users/", userHandler.Show)
	mux.HandleFunc("/users/create", userHandler.Create)
	mux.HandleFunc("/users/update", userHandler.Update)
	mux.HandleFunc("/users/delete/", userHandler.Delete)
	mux.HandleFunc("/posts", postHandler.Index)
	mux.HandleFunc("/posts/", postHandler.Show)
	mux.HandleFunc("/posts/create", postHandler.Create)
	mux.HandleFunc("/posts/update", postHandler.Update)
	mux.HandleFunc("/posts/delete/", postHandler.Delete)
	mux.HandleFunc("/tags", tagHandler.Index)
	mux.HandleFunc("/tags/", tagHandler.Show)
	mux.HandleFunc("/tags/create", tagHandler.Create)
	mux.HandleFunc("/tags/update", tagHandler.Update)
	mux.HandleFunc("/tags/delete/", tagHandler.Delete)
	
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Println(err)
		return 1
	}

	return 0
}