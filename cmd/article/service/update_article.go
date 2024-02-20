package service

import (
	"github.com/lizaiganshenmo/mixStew/cmd/article/dal/db"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/article"
)

func (s *ArticleService) UpdateArticle(req *article.UpdateArticleReq) error {
	err := db.UpdateArticle(s.ctx, &db.ArticleInfo{
		Uid:         req.Uid,
		ArticleId:   req.ArticleId,
		Titile:      *req.Title,
		Description: *req.Description,
		Body:        *req.Body,
	})

	return err
}
