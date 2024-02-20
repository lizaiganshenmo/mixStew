package db

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	SoSoStatus = iota
	LikeStatus
	UnLikeStatus

	ArticleFavoriteInfoTableName = "article_favourite_info"
)

type ArticleFavoriteInfo struct {
	gorm.Model
	Uid       int64 `json:"uid" gorm:"column:uid"`
	ArticleId int64 `json:"article_id" gorm:"column:article_id"`
	Status    int   `json:"status" gorm:"column:status"`
}

func (f *ArticleFavoriteInfo) TableName() string {
	return ArticleFavoriteInfoTableName
}

func CreateFavorite(ctx context.Context, articleId, uid int64) error {
	return MySQLMixStewDB.WithContext(ctx).
		Clauses(clause.OnConflict{
			DoUpdates: clause.AssignmentColumns([]string{"status"}),
		}).
		Create(&ArticleFavoriteInfo{
			Uid:       uid,
			ArticleId: articleId,
			Status:    LikeStatus,
		},
		).Error
}

func UnFavorite(ctx context.Context, articleId, uid int64) error {
	return MySQLMixStewDB.WithContext(ctx).
		Model(&ArticleFavoriteInfo{}).
		Where("uid = ? and article_id = ?", uid, articleId).
		Updates(&ArticleFavoriteInfo{
			Uid:       uid,
			ArticleId: articleId,
			Status:    UnLikeStatus,
		}).Error
}

func GetFavoriteListByUid(ctx context.Context, uid int64, opts ...Option) (articleIds []int64, err error) {
	db := MySQLMixStewDB.WithContext(ctx).Model(&ArticleFavoriteInfo{})

	db.Select("article_id").Where("uid = ? and status = ?", uid, LikeStatus)
	for _, opt := range opts {
		opt(db)
	}

	var infos []*ArticleFavoriteInfo
	err = db.Find(&infos).Error
	if err != nil {
		return
	}

	articleIds = make([]int64, len(infos))
	for i, v := range infos {
		articleIds[i] = v.ArticleId
	}

	return
}

func GetFavoriteCountByArticleId(ctx context.Context, articleId int64, opts ...Option) (cnt int64, err error) {
	db := MySQLMixStewDB.WithContext(ctx).Model(&ArticleFavoriteInfo{})

	db.Where("article_id = ? and status = ?", articleId, LikeStatus)
	for _, opt := range opts {
		opt(db)
	}

	err = db.Count(&cnt).Error
	if err != nil {
		return
	}

	return
}
