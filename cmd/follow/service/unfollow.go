package service

import (
	"github.com/lizaiganshenmo/mixStew/cmd/follow/dal/db"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/follow"
	"github.com/lizaiganshenmo/mixStew/library/errno"
)

func (s *FollowService) Unfollow(req *follow.FollowReq) error {
	if req.Uid == 0 || req.FollowUid == 0 {
		return errno.ParamErr
	}

	err := db.InsertOrUpdate(s.ctx, []*db.FollowInfo{
		{
			Uid:       req.Uid,
			FollowUid: req.FollowUid,
			Status:    db.UnFollowStatus,
		},
	})

	return err

}
