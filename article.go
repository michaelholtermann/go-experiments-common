package common

import "github.com/google/uuid"

type Article struct {
	ID         uuid.UUID
	Title      string
	Author     string
	Characters int16
}

type ArticleList map[uuid.UUID]*Article
