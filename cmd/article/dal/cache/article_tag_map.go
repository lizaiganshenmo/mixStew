package cache

import (
	"context"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

func GetArticleIdsByTag(ctx context.Context, tag string) (articleIds []int64, err error) {
	var strIds []string
	key := GetArticleTagInfoKey(tag)
	strIds, err = RedisClient.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return
	}
	if len(strIds) == 0 {
		return
	}

	articleIds = make([]int64, 0, len(strIds))
	for _, v := range strIds {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			klog.CtxWarnf(ctx, "invalid elem in cache. key: %s, val: %s, err:%+v", key, v, err)
			continue
		}
		articleIds = append(articleIds, id)
	}

	return

}

func SetArticleIdsByTag(ctx context.Context, tag string, articleIds []int64) error {
	idStr := make([]string, 0, len(articleIds))
	for _, v := range articleIds {
		idStr = append(idStr, strconv.FormatInt(v, 10))
	}

	pipe := RedisClient.TxPipeline()
	key := GetArticleTagInfoKey(tag)

	if err := pipe.RPush(ctx, key, idStr).Err(); err != nil {
		return err
	}

	if err := pipe.Expire(ctx, key, ArticleTagInfoKeyExpire*time.Second).Err(); err != nil {
		return err
	}

	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil

}

// get article tags by article_id
func GetArticleTags(ctx context.Context, articleId int64) (tagList []string, err error) {
	key := GetArticleTagsKey(articleId)
	tagList, err = RedisClient.LRange(ctx, key, 0, -1).Result()
	return
}

// saet article tags by article_id
func SetArticleTags(ctx context.Context, articleId int64, tagList []string) error {
	pipe := RedisClient.TxPipeline()
	key := GetArticleTagsKey(articleId)

	if err := pipe.RPush(ctx, key, tagList).Err(); err != nil {
		return err
	}

	if err := pipe.Expire(ctx, key, ArticleTagInfoKeyExpire*time.Second).Err(); err != nil {
		return err
	}

	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil
}
