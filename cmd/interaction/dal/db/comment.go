package db

import (
	"context"

	"gorm.io/gorm"
)

const (
	PendingStatus = iota
	PassStatus
	NotPasstatus

	CommentTableName = "comment"

	CommentTypeOne = 0 // 一级评论
	CommentTypeTwo = 1 // 二级评论
)

type CommentInfo struct {
	gorm.Model
	Uid             int64  `json:"uid" gorm:"column:uid"`
	TargetUid       int64  `json:"target_uid" gorm:"column:target_uid"`
	CommentId       int64  `json:"comment_id" gorm:"column:comment_id"`
	TargetArticleId int64  `json:"target_article_id" gorm:"column:target_article_id"`
	TargetCommentId int64  `json:"target_comment_id" gorm:"column:target_comment_id"`
	CommentType     int    `json:"comment_type" gorm:"comment_type"`
	Status          int    `json:"status" gorm:"column:status"`
	Body            string `json:"body" gorm:"column:body"`
}

func (c *CommentInfo) TableName() string {
	return CommentTableName
}

func CreateComment(ctx context.Context, cInfo *CommentInfo) error {
	return MySQLMixStewDB.WithContext(ctx).Create(cInfo).Error
}

func DeleteComment(ctx context.Context, commentId int64) error {
	return MySQLMixStewDB.WithContext(ctx).Where("comment_id = ?", commentId).Delete(&CommentInfo{}).Error
}

func GetCommentsByArticleId(ctx context.Context, articleId int64) (commentList []*CommentInfo, err error) {
	err = MySQLMixStewDB.WithContext(ctx).Where("target_article_id = ? and status = ?", articleId, PassStatus).Find(&commentList).Error
	return
}

func GetCommentsByCommentId(ctx context.Context, commentId int64) (commentList []*CommentInfo, err error) {
	err = MySQLMixStewDB.WithContext(ctx).Where("comment_id = ? and status = ?", commentId, PassStatus).Find(&commentList).Error
	return
}
