package rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	kitextracing "github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/interaction"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/interaction/interactionservice"
	"github.com/lizaiganshenmo/mixStew/library/constants"
	"github.com/lizaiganshenmo/mixStew/library/errno"
	"github.com/lizaiganshenmo/mixStew/library/middleware"
)

var interctionClient interactionservice.Client

func InitInteractionRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAdd})
	if err != nil {
		panic(err)
	}

	c, err := interactionservice.NewClient(
		constants.InteractionServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(kitextracing.NewClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
		client.WithLoadBalancer(loadbalance.NewWeightedRoundRobinBalancer()),
	)
	if err != nil {
		panic(err)
	}

	interctionClient = c
}

// comment article
func CommentArticle(ctx context.Context, req *interaction.CreateCommentReq) error {
	resp, err := interctionClient.CreateComment(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return nil
}

// get comment by article_id
func GetArticleComment(ctx context.Context, req *interaction.GetCommentReq) ([]*interaction.Comment, error) {
	resp, err := interctionClient.GetComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.Comments, nil
}

// delete comment by comment_id
func DeleteComment(ctx context.Context, req *interaction.DeleteCommentReq) error {
	resp, err := interctionClient.DeleteComment(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return nil
}
