package link

import (
	"http_server/pkg/db"
)

type Repository struct {
	Database *db.DB
}

func NewLinkRepository(database *db.DB) *Repository {
	return &Repository{Database: database}
}

func (repo *Repository) Create(link *Link) (*Link, error) {
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}
