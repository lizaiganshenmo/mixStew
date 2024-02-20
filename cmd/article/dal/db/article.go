package db

import (
	"context"

	"gorm.io/gorm"
)

const (
	ArticleInfoTableName = "article"
)

type ArticleInfo struct {
	gorm.Model
	Uid         int64  `json:"uid" gorm:"column:uid"`
	ArticleId   int64  `json:"article_id" gorm:"column:article_id"`
	Titile      string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Body        string `json:"body" gorm:"column:body"`
}

func (a *ArticleInfo) TableName() string {
	return ArticleInfoTableName
}

// create article
func InsertArticle(ctx context.Context, aInfo []*ArticleInfo) error {
	return MySQLMixStewDB.WithContext(ctx).Model(&ArticleInfo{}).Create(&aInfo).Error
}

// update article
func UpdateArticle(ctx context.Context, aInfo *ArticleInfo) error {
	// m := map[string]interface{}{}
	// m["uid"] = aInfo.Uid
	// m["article_id"] = aInfo.ArticleId
	// m["title"] = aInfo.Titile
	// m["description"] = aInfo.Description
	// m["body"] = aInfo.Body

	// data, err := utils.Struct2map(aInfo)
	// if err != nil {
	// 	return err
	// }

	return MySQLMixStewDB.WithContext(ctx).Model(&ArticleInfo{}).
		Where("uid = ? and article_id = ?", aInfo.Uid, aInfo.ArticleId).Updates(aInfo).Error
}

type Option func(db *gorm.DB) *gorm.DB

func WithArticleId(articleId int64) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("article_id = ?", articleId)
	}
}

func WithUid(uid int64) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("uid = ?", uid)
	}
}

func WithUids(uids []int64) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("uid in ?", uids)
	}
}

func WithLimit(limit int) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit)
	}
}

func WithOffest(offset int) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset)
	}
}

func WithOrder(orderStr string) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(orderStr)
	}
}

func WithArticleIds(articleIds []int64) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("article_id in ?", articleIds)
	}
}

// get article
func QueryArticle(ctx context.Context, opts ...Option) ([]*ArticleInfo, error) {
	db := MySQLMixStewDB.WithContext(ctx).Model(&ArticleInfo{})
	for _, opt := range opts {
		opt(db)
	}

	var res []*ArticleInfo
	var err error
	err = db.Find(&res).Error

	return res, err
}

// del article
func DeleteArticle(ctx context.Context, opts ...Option) error {
	db := MySQLMixStewDB.WithContext(ctx).Model(&ArticleInfo{})
	for _, opt := range opts {
		opt(db)
	}

	err := db.Delete(&ArticleInfo{}).Error
	return err
}
