package service

import (
	"github.com/lizaiganshenmo/mixStew/cmd/follow/dal/db"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/follow"
)

func (s *FollowService) FollowList(req *follow.FollowListReq) ([]int64, error) {
	fInfos, err := db.QueryFollowInfo(s.ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	if len(fInfos) == 0 {
		return []int64{}, nil
	}

	uids := make([]int64, 0, len(fInfos))
	for _, v := range fInfos {
		uids = append(uids, v.FollowUid)
	}

	return uids, nil

}
