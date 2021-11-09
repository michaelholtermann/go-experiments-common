package common

import "github.com/google/uuid"

type article struct {
	ID         uuid.UUID
	Title      string
	Author     string
	Characters int16
}

type ArticleList map[uuid.UUID]*article

func NewArticle(title, author string, characters int16) *article {
	return &article{
		Title:      title,
		Author:     author,
		Characters: characters,
	}
}

func (a *article) SetID(ID uuid.UUID) {
	a.ID = ID
}
