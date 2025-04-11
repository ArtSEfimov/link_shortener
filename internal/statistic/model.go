package statistic

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Statistic struct {
	gorm.Model
	LinkID uint           `json:"link_id"`
	Clicks uint           `json:"clicks"`
	Date   datatypes.Date `json:"date"`
}
