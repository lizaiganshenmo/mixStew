package service

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lizaiganshenmo/mixStew/cmd/article/dal/cache"
	"github.com/lizaiganshenmo/mixStew/cmd/article/dal/db"
)

func (s *ArticleService) FavoriteList(uid int64) (articleIds []int64, err error) {
	var ids []int64
	var exist int64
	ids, exist, err = cache.GetArticleFavoriteList(s.ctx, uid)

	if err != nil {
		klog.CtxWarnf(s.ctx, "cache.GetArticleFavoriteList fail.err:%+v", err)
		return
	}
	if exist == 1 {
		articleIds = ids
		return
	} else {
		ids, err = db.GetFavoriteListByUid(s.ctx, uid, db.WithLimit(100), db.WithOrder("created_at desc"))
		if err != nil {
			klog.CtxWarnf(s.ctx, "db.GetFavoriteListByUid fail.err:%+v", err)

			return
		}

		err = cache.SetArticleFavoriteList(s.ctx, uid, articleIds)
		if err != nil {
			klog.CtxWarnf(s.ctx, "cache.SetArticleFavoriteList fail.err:%+v", err)
			return
		}

	}

	return

}
