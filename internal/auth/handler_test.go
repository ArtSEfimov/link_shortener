package auth_test

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"http_server/configs"
	"http_server/internal/auth"
	"http_server/internal/user"
	"http_server/pkg/db"
	"net/http"
	"net/http/httptest"
	"testing"
)

func bootstrap() (*auth.Handler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database}))
	if err != nil {
		return nil, nil, err
	}

	userRepository := user.NewRepository(&db.DB{gormDB})

	handler := auth.Handler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		Service: auth.NewService(userRepository),
	}
	return &handler, mock, nil
}

func TestLoginHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("art1@m.com", "$2a$10$9.sVok04vXgifIgrVWA8buoVsbwgqRZ3hxDVEXG20fAk.DY4duFzC")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	if err != nil {
		t.Fatal(err)
		return
	}
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "art1@m.com",
		Password: "123",
	})

	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()

	r, err := http.NewRequest("POST", "/auth/login", reader)
	if err != nil {
		t.Fatal(err)
	}
	handler.Login()(w, r)
	if w.Code != http.StatusOK {
		t.Error("Wrong status code: expected", http.StatusOK, "got", w.Code)
	}

}

func TestRegisterHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	rows := sqlmock.NewRows([]string{"email", "password", "name"})
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
	mock.ExpectCommit()

	if err != nil {
		t.Fatal(err)
		return
	}
	data, _ := json.Marshal(&auth.RegisterRequest{
		Email:    "art1@m.com",
		Password: "123",
		Name:     "Art1",
	})

	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()

	r, err := http.NewRequest("POST", "/auth/login", reader)
	if err != nil {
		t.Fatal(err)
	}
	handler.Register()(w, r)
	if w.Code != http.StatusCreated {
		t.Error("Wrong status code: expected", http.StatusCreated, "got", w.Code)
	}

}
