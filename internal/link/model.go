package link

import (
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,	
	}
	link.GenHash()
	return link
}

func (link *Link) GenHash() {
	link.Hash = RandHashGen(10)
}

var latterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") 

func RandHashGen(n int) string {
	b := make([]rune, n)
	for i := range b{
		b[i] = latterRunes[rand.Intn(len(latterRunes))]
	}
	return string(b)
}