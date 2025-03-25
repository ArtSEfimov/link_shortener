package link

import (
	"gorm.io/gorm"
	"math/rand"
	"time"
)

const (
	MIN = 5
	MAX = 10
)

var allowableRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(urlString string) *Link {

	link := &Link{
		Url: urlString,
	}

	link.GenerateHash()

	return link
}

func (link *Link) GenerateHash() {
	randSource := rand.NewSource(time.Now().Unix())
	randGen := rand.New(randSource)
	hashSlice := make([]rune, randGen.Intn(MAX-MIN+1)+MIN)
	for i := range hashSlice {
		hashSlice[i] = allowableRunes[randGen.Intn(len(allowableRunes))]
	}

	link.Hash = string(hashSlice)
}
