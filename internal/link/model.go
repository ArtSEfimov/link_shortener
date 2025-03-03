package link

import (
	"gorm.io/gorm"
	"math/rand"
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(urlString string) *Link {
	return &Link{
		Url:  urlString,
		Hash: RandStringRunes(6),
	}
}

var allowableRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")

func RandStringRunes(n int) string {
	hashSlice := make([]rune, n)
	for i := range hashSlice {
		hashSlice[i] = allowableRunes[rand.Intn(len(allowableRunes))]
	}
	return string(hashSlice)
}
