package link

import (
	"fmt"
	"gorm.io/gorm"
	"http_server/configs"
	"http_server/pkg/di"
	"http_server/pkg/middleware"
	"http_server/pkg/request"
	"http_server/pkg/response"
	"net/http"
	"strconv"
)

type HandlerDeps struct {
	LinkRepository      *Repository
	StatisticRepository di.StatisticRepositoryInterface
	Config              *configs.Config
}

type Handler struct {
	Repository          *Repository
	StatisticRepository di.StatisticRepositoryInterface
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		Repository:          deps.LinkRepository,
		StatisticRepository: deps.StatisticRepository,
	}
	router.HandleFunc("POST /link", handler.Create())
	router.Handle("PATCH /link/{id}", middleware.IsAuthed(handler.Update(), deps.Config))
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{hash}", handler.GoTo())
	router.Handle("GET /link", middleware.IsAuthed(handler.GetAll(), deps.Config))
}

func (handler *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, handleBodyErr := request.HandleBody[CreateRequest](&w, r)
		if handleBodyErr != nil {
			return
		}
		link := NewLink(body.Url)

		for {
			_, err := handler.Repository.GetByHash(link.Hash)
			if err != nil {
				break
			}
			link.GenerateHash()
		}

		createdLink, createErr := handler.Repository.Create(link)
		if createErr != nil {
			http.Error(w, createErr.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, createdLink, http.StatusOK)

	}
}
func (handler *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if email, ok := r.Context().Value(middleware.ContextEmailKey).(string); ok {
			fmt.Println(email)
		}

		body, handleBodyErr := request.HandleBody[UpdateRequest](&w, r)
		if handleBodyErr != nil {
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		updatedLink, err := handler.Repository.Update(&Link{
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
func (handler *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := handler.Repository.GetById(uint(id)); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = handler.Repository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Json(w, nil, http.StatusOK)

	}
}
func (handler *Handler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := handler.Repository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handler.StatisticRepository.AddClick(link.ID)
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *Handler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
		}

		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			http.Error(w, "Invalid offset parameter", http.StatusBadRequest)
		}

		links := handler.Repository.GetAll(limit, offset)
		count := handler.Repository.CountAll()

		response.Json(w, GetAllLinksResponse{
			Links: links,
			Count: count,
		}, http.StatusOK,
		)
	}
}
