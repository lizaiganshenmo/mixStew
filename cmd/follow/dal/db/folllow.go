package db

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	UnFollowStatus = iota
	FollowStatus
	BlackStatus

	FollowInfoTableName = "follow_info"

	maxValidStatus = 2
	minValidStatus = 0
)

type FollowInfo struct {
	gorm.Model
	Uid       int64 `json:"uid" gorm:"column:uid"`
	FollowUid int64 `json:"follow_uid" gorm:"column:follow_uid"`
	Status    int   `json:"status" gorm:"column:status"`
}

func (f *FollowInfo) TableName() string {
	return FollowInfoTableName
}

// create or update follow info
func InsertOrUpdate(ctx context.Context, fInfo []*FollowInfo) error {
	return MySQLMixStewDB.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uid_follow_uid"}},
		DoUpdates: clause.AssignmentColumns([]string{"status"}),
	}).Create(&fInfo).Error
}

// query follow info by uid
func QueryFollowInfo(ctx context.Context, uid int64) ([]*FollowInfo, error) {
	res := make([]*FollowInfo, 0)
	if err := MySQLMixStewDB.WithContext(ctx).Where("uid = ?", uid).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// query  follow uid->follow_uid
func QueryFollowStatus(ctx context.Context, uid, followUid int64) ([]*FollowInfo, error) {
	res := make([]*FollowInfo, 0)
	if err := MySQLMixStewDB.WithContext(ctx).Where("uid = ? and follow_uid = ?", uid, followUid).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func IsValidStatus(status int) bool {
	return status >= minValidStatus && status <= maxValidStatus
}
