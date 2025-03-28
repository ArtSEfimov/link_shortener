package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique,index"`
	Password string
	Name     string
}
