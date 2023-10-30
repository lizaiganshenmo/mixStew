package handler

import (
	"context"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/lizaiganshenmo/mixStew/cmd/api/biz/rpc"
	"github.com/lizaiganshenmo/mixStew/cmd/api/biz/types/resp"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/follow"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/user"
	"github.com/lizaiganshenmo/mixStew/library/constants"
	"github.com/lizaiganshenmo/mixStew/library/errno"
	"github.com/spf13/cast"
)

func Follow(ctx context.Context, c *app.RequestContext) {
	var uid int64
	t, ok := c.Get(constants.IdentityKey)
	if !ok {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	uid = int64(t.(float64))
	followUidStr := c.Param("follow_uid")
	followUid := cast.ToInt64(followUidStr)

	if uid == 0 || followUid == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := rpc.Follow(ctx, &follow.FollowReq{Uid: uid, FollowUid: followUid})
	if err != nil {
		SendResponse(c, err, nil)
		hlog.Warnf("rpc.Follow err : %v , uid: %d, follow_uid: %d", err, uid, followUid)
		return
	}

	SendResponse(c, errno.Success, nil)

}

func UnFollow(ctx context.Context, c *app.RequestContext) {
	var uid int64
	t, ok := c.Get(constants.IdentityKey)
	if !ok {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	uid = int64(t.(float64))
	followUidStr := c.Param("follow_uid")
	followUid := cast.ToInt64(followUidStr)

	if uid == 0 || followUid == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := rpc.UnFollow(ctx, &follow.FollowReq{Uid: uid, FollowUid: followUid})
	if err != nil {
		SendResponse(c, err, nil)
		hlog.Warnf("rpc.UnFollow err : %v , uid: %d, follow_uid: %d", err, uid, followUid)
		return
	}

	SendResponse(c, errno.Success, nil)

}

func GetProfile(ctx context.Context, c *app.RequestContext) {
	var uid int64
	t, ok := c.Get(constants.IdentityKey)
	// hlog.Warnf("uid ok : %+v, %t", t, ok)
	if !ok {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	uid = int64(t.(float64))
	followUidStr := c.Param("follow_uid")
	followUid := cast.ToInt64(followUidStr)
	// hlog.Warnf("followUid : %d", followUid)
	if followUid == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	var user1 *user.User
	var isFollowing bool
	var err1, err2 error
	var wg sync.WaitGroup
	// s1 查询user信息
	wg.Add(1)
	go func() {
		defer wg.Done()
		user1, err1 = rpc.GetUser(ctx, &user.GetUserReq{Uid: followUid})
	}()

	// s2 查询 follow信息
	wg.Add(1)
	go func() {
		defer wg.Done()
		isFollowing, err2 = rpc.IsFollow(ctx, &follow.FollowReq{Uid: uid, FollowUid: followUid})
	}()

	wg.Wait()
	// hlog.Warnf("user1, err1 : %+v, %+v", user1, err1)
	// hlog.Warnf("isFollowing, err2 : %+v, %+v", isFollowing, err2)

	data := &resp.Profile{}
	if err1 == nil {
		data.Username = user1.UserRegister.Username
		data.Bio = user1.Bio
		data.Image = user1.Image
	} else {
		hlog.Warnf("rpc.GetUser err: %+v", err1)
	}

	if err2 != nil {
		hlog.Warnf("rpc.IsFollow err: %+v", err2)
	} else {
		data.Following = isFollowing
	}

	SendResponse(c, errno.Success, data)

}
