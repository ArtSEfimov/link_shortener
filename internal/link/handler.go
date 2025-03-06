package link

import (
	"fmt"
	req "http_server/pkg/request"
	"http_server/pkg/response"
	"net/http"
)

type HandlerDeps struct {
	LinkRepository *Repository
}

type Handler struct {
	LinkRepository *Repository
}

func NewLinkHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
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
		createdLink, createErr := h.LinkRepository.Create(link)
		if createErr != nil {
			http.Error(w, createErr.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, createdLink, http.StatusOK)

	}
}
func (h *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func (h *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(id)
	}
}
func (h *Handler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
