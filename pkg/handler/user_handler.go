package handler

import (
	"log"
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/Daaaai0809/bun_prac/pkg/domain/entity"
	interactor "github.com/Daaaai0809/bun_prac/pkg/usecase"
)

type UserHandler struct {
	userInteractor interactor.IUserInteractor
}

func NewUserHandler(userInteractor interactor.IUserInteractor) *UserHandler {
	return &UserHandler{
		userInteractor: userInteractor,
	}
}

func (handler *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	users, err := handler.userInteractor.FindAll(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// convert users to json
	usersJson, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(usersJson))
}

func (handler *UserHandler) Show(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := handler.userInteractor.FindById(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// convert user to json
	userJson, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userJson))
}

func (handler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	newUser := &entity.User{}

	if err := json.NewDecoder(r.Body).Decode(newUser); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(newUser)

	if err := handler.userInteractor.Create(r.Context(), newUser); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	user:= entity.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.userInteractor.Update(r.Context(), &user); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.userInteractor.Delete(r.Context(), id); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}