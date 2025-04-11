package statistic

import (
	"gorm.io/datatypes"
	"http_server/pkg/db"
	"time"
)

type Repository struct {
	*db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{db}
}

func (repo *Repository) AddClick(linkID uint) {
	var statistic Statistic
	currentDate := datatypes.Date(time.Now())
	repo.DB.Find(&statistic, "link_id = ? and date = ?", linkID, currentDate)
	if statistic.ID == 0 {
		repo.DB.Create(&Statistic{
			LinkID: linkID,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		statistic.Clicks++
		repo.DB.Save(&statistic)
	}
}
