package auth

import (
	"fmt"
	"http_server/configs"
	req "http_server/pkg/request"
	"http_server/pkg/response"
	"net/http"
)

type HandlerDeps struct {
	*configs.Config
	*Service
}
type Handler struct {
	*configs.Config
	*Service
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		deps.Config,
		deps.Service,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *Handler) Login() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		body, handleErr := req.HandleBody[LoginRequest](&writer, request)
		if handleErr != nil {
			return
		}
		fmt.Println(body)

		data := LoginResponse{
			Token: "123",
		}

		response.Json(writer, data, http.StatusOK)

	}
}

func (handler *Handler) Register() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, handleErr := req.HandleBody[RegisterRequest](&writer, request)
		if handleErr != nil {
			return
		}

		_, err := handler.Service.Register(body.Email, body.Password, body.Name)
		if err != nil {
			return
		}
	}
}
