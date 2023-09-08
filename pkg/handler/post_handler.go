package handler

import (
	"net/http"
	"encoding/json"
	"strconv"
	"log"

	"github.com/Daaaai0809/bun_prac/pkg/domain/entity"
	interactor "github.com/Daaaai0809/bun_prac/pkg/usecase"
)

type PostHandler struct {
	postInteractor interactor.IPostInteractor
}

func NewPostHandler(postInteractor interactor.IPostInteractor) *PostHandler {
	return &PostHandler{
		postInteractor: postInteractor,
	}
}

func (handler *PostHandler) Index(w http.ResponseWriter, r *http.Request) {
	posts, err := handler.postInteractor.FindAll(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// convert posts to json
	postsJson, err := json.Marshal(posts)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}	

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(postsJson))
}

func (handler *PostHandler) Show(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	post, err := handler.postInteractor.FindById(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// convert post to json
	postJson, err := json.Marshal(post)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(postJson))
}

func (handler *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	newPost := &entity.Post{}

	if err := json.NewDecoder(r.Body).Decode(newPost); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(newPost)

	if err := handler.postInteractor.Create(r.Context(), newPost); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
	post := &entity.Post{}

	if err := json.NewDecoder(r.Body).Decode(post); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.postInteractor.Update(r.Context(), post); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var id int64

	if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.postInteractor.Delete(r.Context(), id); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}