package utils

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/lizaiganshenmo/mixStew/library/constants"
)

func GetUid(c *app.RequestContext) (uid int64) {
	t, ok := c.Get(constants.IdentityKey)
	if !ok {
		return
	}

	uid = int64(t.(float64))

	return
}
