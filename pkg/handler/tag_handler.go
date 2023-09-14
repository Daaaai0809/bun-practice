package handler

import (
	"net/http"
	"strconv"
	"log"
	"encoding/json"

	"github.com/Daaaai0809/bun_prac/pkg/domain/entity"
	interactor "github.com/Daaaai0809/bun_prac/pkg/usecase"
)

type TagHandler struct {
	tagInteractor interactor.ITagInteractor
}

func NewTagHandler(tagInteractor interactor.ITagInteractor) *TagHandler {
	return &TagHandler{
		tagInteractor: tagInteractor,
	}
}

func (handler *TagHandler) Index(w http.ResponseWriter, r *http.Request) {
	tags, err := handler.tagInteractor.FindAll(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// convert tags to json
	tagsJson, err := json.Marshal(tags)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tagsJson))
}

func (handler *TagHandler) Show(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tag, err := handler.tagInteractor.FindById(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// convert tag to json
	tagJson, err := json.Marshal(tag)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tagJson))
}

func (handler *TagHandler) Create(w http.ResponseWriter, r *http.Request) {
	var tag entity.Tag
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.tagInteractor.Create(r.Context(), &tag); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *TagHandler) Update(w http.ResponseWriter, r *http.Request) {
	var tag entity.Tag
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.tagInteractor.Update(r.Context(), &tag); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *TagHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handler.tagInteractor.Delete(r.Context(), id); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}