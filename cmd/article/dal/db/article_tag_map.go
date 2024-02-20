package db

import (
	"context"

	"gorm.io/gorm"
)

const (
	ArticleTagMapTableName = "article_tag_map"
)

type ArticleTagMap struct {
	gorm.Model
	ArticleId int64  `json:"article_id" gorm:"column:article_id"`
	TagName   string `json:"tag_name" gorm:"column:tag_name"`
}

func (f *ArticleTagMap) TableName() string {
	return ArticleTagMapTableName
}

// create article tag info
func CreateArticleTags(ctx context.Context, infos []*ArticleTagMap) error {
	return MySQLMixStewDB.WithContext(ctx).Model(&ArticleTagMap{}).Create(&infos).Error
}

func DeleteArticleTags(ctx context.Context, articleIds []int64) error {
	return MySQLMixStewDB.WithContext(ctx).Model(&ArticleTagMap{}).Where("article_id in ?", articleIds).Delete(&ArticleTagMap{}).Error
}

// get articleIds By tag
func GetArticleIdsByTag(ctx context.Context, tag string) (articleIds []int64, err error) {
	err = MySQLMixStewDB.WithContext(ctx).Model(&ArticleTagMap{}).Select("article_id").Where("tag_name = ?", tag).Pluck("article_id", &articleIds).Error
	return
}

// get tags By articleId
func GetTagsByArticleId(ctx context.Context, articleId int64) (tagList []string, err error) {
	err = MySQLMixStewDB.WithContext(ctx).Model(&ArticleTagMap{}).Select("tag_name").Where("article_id = ?", articleId).Pluck("tag_name", &tagList).Error
	return
}
