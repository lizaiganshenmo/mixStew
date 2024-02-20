package service

import (
	"sync"

	"github.com/lizaiganshenmo/mixStew/cmd/article/dal/db"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/article"
)

func (s *ArticleService) DeleteArticle(req *article.DeleteArticleReq) error {
	var wg sync.WaitGroup
	var err1, err2 error
	wg.Add(1)
	go func() {
		defer wg.Done()
		err1 = db.DeleteArticle(s.ctx, db.WithArticleId(req.ArticleId), db.WithUid(req.Uid))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err2 = db.DeleteArticleTags(s.ctx, []int64{req.ArticleId})
	}()

	wg.Wait()

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	return nil

}
