package service

import (
	"github.com/lizaiganshenmo/mixStew/cmd/article/dal/db"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/article"
)

func (s *ArticleService) CreateArticle(req *article.CreateArticleReq) error {
	// s1 生成 article_id
	articleID := db.GenArticleId()

	// s2 入库 article 表
	err := db.InsertArticle(s.ctx, []*db.ArticleInfo{
		{
			Uid:         req.Uid,
			ArticleId:   articleID,
			Titile:      req.Title,
			Description: req.Description,
			Body:        req.Body,
		},
	})
	if err != nil {
		return err
	}

	// s3 入库tag 表
	if len(req.TagList) != 0 {
		tagInfos := make([]*db.ArticleTagMap, 0, len(req.TagList))
		for _, v := range req.TagList {
			tagInfos = append(tagInfos, &db.ArticleTagMap{
				ArticleId: articleID,
				TagName:   v,
			})
		}

		err := db.CreateArticleTags(s.ctx, tagInfos)
		if err != nil {
			return err
		}

	}

	return nil

}
