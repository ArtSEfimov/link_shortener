package auth

import (
	"fmt"
	"http_server/configs"
	req "http_server/pkg/request"
	"http_server/pkg/response"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}
type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (h *AuthHandler) Login() http.HandlerFunc {
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

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, handleErr := req.HandleBody[RegisterRequest](&writer, request)
		if handleErr != nil {
			return
		}

		fmt.Println(body)
	}
}
