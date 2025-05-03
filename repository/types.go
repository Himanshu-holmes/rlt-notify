package repository

import (
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
)

type Article struct {
	ID uint64
	Title string
	Content string
}

type ArticleRepository interface {
	ById(ctx context.Context,id int)(Article,error)
}

type SimpleSummaryArticle struct {
	ID uint64 `json:"id"`
	Title string `json:"title"`
	Summary string `json:"summary"`
	More string `json:"more"`
}

