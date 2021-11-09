package common

import "github.com/google/uuid"

type Article struct {
	ID         uuid.UUID
	Title      string
	Author     string
	Characters int16
}

type ArticleList map[uuid.UUID]*Article

func NewArticle(title, author string, characters int16) *Article {
	return &Article{
		Title:      title,
		Author:     author,
		Characters: characters,
	}
}

func (a *Article) SetID(ID uuid.UUID) {
	a.ID = ID
}
