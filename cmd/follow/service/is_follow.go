package service

import (
	"github.com/lizaiganshenmo/mixStew/cmd/follow/dal/db"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/follow"
)

func (s *FollowService) IsFollow(req *follow.FollowReq) (bool, error) {
	infos, err := db.QueryFollowStatus(s.ctx, req.Uid, req.FollowUid)
	if err != nil {
		return false, err
	}

	if len(infos) == 0 || infos[0].Status != db.FollowStatus {
		return false, nil
	}

	return true, nil

}
