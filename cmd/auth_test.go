package main

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"http_server/internal/auth"
	"http_server/internal/user"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func initDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, openErr := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if openErr != nil {
		panic(openErr)
	}

	return db
}

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "test@test.com",
		Password: "$2a$10$9.sVok04vXgifIgrVWA8buoVsbwgqRZ3hxDVEXG20fAk.DY4duFzC",
		Name:     "Art1",
	})
}
func removeData(db *gorm.DB) {
	db.Unscoped().Where("email = ?", "test@test.com").Delete(&user.User{})
}
func TestLoginSuccess(t *testing.T) {
	// Prepare
	db := initDB()
	initData(db)
	defer removeData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "test@test.com",
		Password: "123",
	})

	response, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != 200 {
		t.Fatalf("Expected %d got %d", 200, response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	var responseData auth.LoginResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		t.Fatal(err)
	}
	if responseData.Token == "" {
		t.Fatal("Token is empty")
	}

}
func TestLoginFail(t *testing.T) {
	ts := httptest.NewServer(App())
	db := initDB()
	initData(db)
	defer ts.Close()
	defer removeData(db)

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "art1@m.com",
		Password: "2123",
	})

	response, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != 401 {
		t.Fatalf("Expected %d got %d", 401, response.StatusCode)
	}
}
