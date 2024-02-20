package cache

import (
	"context"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

// redis has or the article favorite key
func IsExistArticle(ctx context.Context, articleId int64) (int64, error) {
	return RedisClient.Exists(ctx, GetArticleFavoriteKey(articleId)).Result()
}

func IsFavorited(ctx context.Context, articleId, uid int64) (bool, error) {
	return RedisClient.SIsMember(ctx, GetArticleFavoriteKey(articleId), uid).Result()
}

func AddFavorite(ctx context.Context, articleId, uid int64) error {
	pipe := RedisClient.TxPipeline()

	// add faourite count
	if err := pipe.Incr(ctx, GetArticleFavoriteCountKey(articleId)).Err(); err != nil {
		return err
	}

	// set faourite count expire time
	if err := pipe.Expire(ctx, GetArticleFavoriteCountKey(articleId), ArticleFavouriteCountKeyExpire*time.Second).Err(); err != nil {
		return err
	}

	// add article faourite info
	if err := pipe.SAdd(ctx, GetArticleFavoriteKey(articleId), uid).Err(); err != nil {
		return err
	}

	// add uid faourite info
	if err := pipe.SAdd(ctx, GetPersonFavoriteKey(uid), articleId).Err(); err != nil {
		return err
	}

	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil

}

func ReduceFavorite(ctx context.Context, articleId, uid int64) error {
	pipe := RedisClient.TxPipeline()

	// reduce faourite count
	if err := pipe.Decr(ctx, GetArticleFavoriteCountKey(articleId)).Err(); err != nil {
		return err
	}

	// remove faourite info
	if err := pipe.SRem(ctx, GetArticleFavoriteKey(articleId), uid).Err(); err != nil {
		return err
	}

	// remove uid faourite info
	if err := pipe.SRem(ctx, GetPersonFavoriteKey(uid), articleId).Err(); err != nil {
		return err
	}

	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil

}

func DelArticleFavoriteInfo(ctx context.Context, articleId, uid int64) error {
	pipe := RedisClient.TxPipeline()

	// del favorite count
	if err := pipe.Del(ctx, GetArticleFavoriteCountKey(articleId)).Err(); err != nil {
		return err
	}

	// del favorite info
	if err := pipe.Unlink(ctx, GetArticleFavoriteKey(articleId)).Err(); err != nil {
		return err
	}

	// del uid faourite info
	if err := pipe.Unlink(ctx, GetPersonFavoriteKey(uid)).Err(); err != nil {
		return err
	}

	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func GetArticleFavoriteList(ctx context.Context, uid int64) (articleIds []int64, exist int64, err error) {
	key := GetPersonFavoriteKey(uid)
	var idsStr []string

	pipe := RedisClient.TxPipeline()

	exist, err = pipe.Exists(ctx, key).Result()
	if err != nil {
		return
	}

	idsStr, err = pipe.SMembers(ctx, key).Result()
	if err != nil {
		return
	}

	if _, err = pipe.Exec(ctx); err != nil {
		return
	}

	if exist != 1 || len(idsStr) == 0 {
		return
	}

	articleIds = make([]int64, 0, len(idsStr))
	for _, v := range idsStr {
		id, err1 := strconv.ParseInt(v, 10, 64)
		if err1 != nil {
			klog.CtxWarnf(ctx, "invalid elem in cache.key: %s,val:%s,err:%v", key, v, err1)
			continue
		}
		articleIds = append(articleIds, id)

	}

	return

}

func SetArticleFavoriteList(ctx context.Context, uid int64, articleIds []int64) error {
	idStr := make([]string, 0, len(articleIds))
	for _, v := range articleIds {
		idStr = append(idStr, strconv.FormatInt(v, 10))
	}

	pipe := RedisClient.TxPipeline()
	key := GetPersonFavoriteKey(uid)
	if err := pipe.RPush(ctx, key, idStr).Err(); err != nil {
		return err
	}

	if err := pipe.Expire(ctx, key, ArticleUidFavouriteInfoExpire*time.Second).Err(); err != nil {
		return err
	}

	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil

}

func SetArticleFavouriteCount(ctx context.Context, articleId, count int64) error {
	return RedisClient.Set(ctx, GetArticleFavoriteCountKey(articleId), count, ArticleFavouriteCountKeyExpire*time.Second).Err()
}

func GetArticleFavoriteCount(ctx context.Context, articleId int64) (int64, error) {
	cntStr, err := RedisClient.Get(ctx, GetArticleFavoriteCountKey(articleId)).Result()
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(cntStr, 10, 64)
}
