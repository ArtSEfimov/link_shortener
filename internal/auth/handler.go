package auth

import (
	"encoding/json"
	"fmt"
	"http_server/configs"
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

		var payload LoginRequest
		decodeErr := json.NewDecoder(request.Body).Decode(&payload)
		if decodeErr != nil {
			response.Json(writer, decodeErr.Error(), http.StatusBadRequest)
		}

		fmt.Println(payload)

		data := LoginResponse{
			Token: "123",
		}

		response.Json(writer, data, http.StatusOK)

	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Register")
	}
}
