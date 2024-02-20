package cache

import "fmt"

const (
	// article_favorite_info
	ArticleFavouriteCountKeyExpire = 600 // 10 min
	ArticleUidFavouriteInfoExpire  = 600 // 10 min
	ArticleFavouriteKeyPre         = "article_favourite_"
	ArticleFavouriteCountKeyPre    = "article_favourite_count_"
	PersonFavouriteKeyPre          = "person_favourite_"

	// article_tag_info
	ArticleTagInfoKeyExpire = 600
	ArticleTagInfoKeyPre    = "article_tag_map_info_"
	ArticleTagsPre          = "article_tags_"
)

func GetArticleFavoriteKey(articleId int64) string {
	return fmt.Sprintf("%s%d", ArticleFavouriteKeyPre, articleId)
}
