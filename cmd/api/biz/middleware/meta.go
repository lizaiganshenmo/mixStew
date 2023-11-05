package middleware

import (
	"context"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/requestid"
	"github.com/lizaiganshenmo/mixStew/library/constants"
)

func AddMetaInfo() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		ctx = metainfo.WithPersistentValue(ctx, constants.RequestIdKey, requestid.Get(c))
		c.Next(ctx)
	}
}
