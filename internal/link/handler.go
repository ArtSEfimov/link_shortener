package link

import (
	"gorm.io/gorm"
	"http_server/pkg/middleware"
	req "http_server/pkg/request"
	"http_server/pkg/response"
	"net/http"
	"strconv"
)

type HandlerDeps struct {
	LinkRepository *Repository
}

type Handler struct {
	Repository *Repository
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		Repository: deps.LinkRepository,
	}
	router.HandleFunc("POST /link", handler.Create())
	router.Handle("PATCH /link/{id}", middleware.IsAuthed(handler.Update()))
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{hash}", handler.GoTo())
}

func (h *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, handleBodyErr := req.HandleBody[CreateRequest](&w, r)
		if handleBodyErr != nil {
			return
		}
		link := NewLink(body.Url)

		for {
			_, err := h.Repository.GetByHash(link.Hash)
			if err != nil {
				break
			}
			link.GenerateHash()
		}

		createdLink, createErr := h.Repository.Create(link)
		if createErr != nil {
			http.Error(w, createErr.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, createdLink, http.StatusOK)

	}
}
func (h *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, handleBodyErr := req.HandleBody[UpdateRequest](&w, r)
		if handleBodyErr != nil {
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		updatedLink, err := h.Repository.Update(&Link{
			Model: gorm.Model{
				ID: uint(id),
			},
			Url:  body.Url,
			Hash: body.Hash,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, updatedLink, http.StatusOK)
	}
}
func (h *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := h.Repository.GetById(uint(id)); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = h.Repository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Json(w, nil, http.StatusOK)

	}
}
func (h *Handler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := h.Repository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}
