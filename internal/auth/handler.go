package auth

import (
	"http_server/configs"
	"http_server/pkg/jwt"
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
	return func(w http.ResponseWriter, r *http.Request) {

		body, handleErr := req.HandleBody[LoginRequest](&w, r)
		if handleErr != nil {
			return
		}

		email, err := handler.Service.Login(UserData{
			Email:    body.Email,
			Password: body.Password,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.Data{
			Email: email,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := LoginResponse{
			Token: token,
		}

		response.Json(w, data, http.StatusOK)

	}
}

func (handler *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, handleErr := req.HandleBody[RegisterRequest](&w, r)
		if handleErr != nil {
			return
		}

		email, err := handler.Service.Register(UserData{
			Email:    body.Email,
			Password: body.Password,
			Name:     body.Name,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.Data{Email: email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := RegisterResponse{
			Token: token,
		}
		response.Json(w, data, http.StatusCreated)

	}
}
