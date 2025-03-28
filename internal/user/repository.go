package user

import "http_server/pkg/db"

type Repository struct {
	database *db.DB
}

func NewRepository(database *db.DB) *Repository {
	return &Repository{database: database}
}

func (repository *Repository) Create(user *User) (*User, error) {
	result := repository.database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repository *Repository) FindByEmail(email string) (*User, error) {
	var user User
	result := repository.database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
