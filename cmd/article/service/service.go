package service

import "context"

type ArticleService struct {
	ctx context.Context
}

// NewArticleService new ArticleService
func NewArticleService(ctx context.Context) *ArticleService {
	return &ArticleService{ctx: ctx}
}
