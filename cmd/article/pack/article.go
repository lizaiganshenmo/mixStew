package pack

import (
	"github.com/lizaiganshenmo/mixStew/cmd/article/dal/db"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/article"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/user"
)

func Author(data *user.User, following bool) *article.Author {
	return &article.Author{
		Username:  data.UserRegister.Username,
		Bio:       data.Bio,
		Image:     data.Image,
		Following: following,
	}
}

func Article(articleInfo db.ArticleInfo, data *user.User, following, favorited bool, favCount int64, tagList []string) *article.Article {
	return &article.Article{
		ArticleId:      articleInfo.ArticleId,
		Title:          articleInfo.Titile,
		Body:           articleInfo.Body,
		Description:    articleInfo.Description,
		CreatedAt:      articleInfo.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:      articleInfo.UpdatedAt.Format("2006-01-02 15:04:05"),
		Favorited:      favorited,
		FavoritesCount: int32(favCount),
		Author:         Author(data, following),
		TagList:        tagList,
	}

}
