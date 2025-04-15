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

func (repo *Repository) GetStatistic(by string, from, to time.Time) []GetResponse {
	var statistics []GetResponse
	var selectQuery string
	switch by {
	case GroupByDay:
		selectQuery = "to_char(date, 'YYYY-MM-DD') as period, sum(clicks)"
	case GroupByMonth:
		selectQuery = "to_char(date, 'YYYY-MM') as period, sum(clicks)"
	}

	repo.DB.Table("statistics").
		Select(selectQuery).
		Where("date BETWEEN ? AND ?", from, to).
		Group("period").
		Order("period").
		Scan(&statistics)

	return statistics
}
