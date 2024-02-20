package rpc

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	kitextracing "github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/article"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/article/articleservice"
	"github.com/lizaiganshenmo/mixStew/library/constants"
	"github.com/lizaiganshenmo/mixStew/library/errno"
	"github.com/lizaiganshenmo/mixStew/library/middleware"
)

var articleClient articleservice.Client

func InitArticleRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAdd})
	if err != nil {
		panic(err)
	}

	c, err := articleservice.NewClient(
		constants.ArticleServiceName,
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

	articleClient = c
}

// create article
func CreateArticle(ctx context.Context, req *article.CreateArticleReq) error {
	resp, err := articleClient.CreateArticle(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return nil
}

// update article
func UpdateArticle(ctx context.Context, req *article.UpdateArticleReq) error {
	resp, err := articleClient.UpdateArticle(ctx, req)

	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return nil

}

// delete article
func DeleteArticle(ctx context.Context, req *article.DeleteArticleReq) error {
	resp, err := articleClient.DeleteArticle(ctx, req)

	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return nil
}

// mget article
func MGetArticle(ctx context.Context, req *article.MGetArticleReq) (articles []*article.Article, err error) {
	resp, err := articleClient.MGetArticle(ctx, req)

	if err != nil {
		return
	}
	if resp.BaseResp.StatusCode != 0 {
		err = errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
		return
	}

	return resp.Articles, nil
}

func MGetFeedArticle(ctx context.Context, req *article.MGetArticleReq) (articles []*article.Article, err error) {
	resp, err := articleClient.MGetArticle(ctx, req)

	if err != nil {
		return
	}
	if resp.BaseResp.StatusCode != 0 {
		err = errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
		return
	}

	return resp.Articles, nil
}

func CommentArticle(ctx context.Context, req *article.CommentArticleReq) (err error) {
	resp, err := articleClient.CommentArticle(ctx, req)

	if err != nil {
		return
	}
	if resp.BaseResp.StatusCode != 0 {
		err = errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
		return
	}

	return

}

func GetArticleComment(ctx context.Context, req *article.GetArticleCommentReq) (comments []*article.Comment, err error) {
	var resp *article.GetArticleCommentResp
	resp, err = articleClient.GetArticleComment(ctx, req)

	if err != nil {
		return
	}
	if resp.BaseResp.StatusCode != 0 {
		err = errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
		return
	}

	return resp.Comments, err
}

func DeleteArticleComment(ctx context.Context, req *article.DeleteCommentReq) (err error) {
	var resp *article.DeleteCommentResp
	resp, err = articleClient.DeleteComment(ctx, req)

	if err != nil {
		return
	}
	if resp.BaseResp.StatusCode != 0 {
		err = errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
		return
	}

	return
}

func FavoriteArticle(ctx context.Context, req *article.FavoriteArticleReq) (err error) {
	resp, err := articleClient.FavoriteArticle(ctx, req)

	if err != nil {
		return
	}
	if resp.BaseResp.StatusCode != 0 {
		err = errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
		return
	}

	return
}

func UnFavoriteArticle(ctx context.Context, req *article.UnFavoriteArticleReq) (err error) {
	resp, err := articleClient.UnFavoriteArticle(ctx, req)

	if err != nil {
		return
	}
	if resp.BaseResp.StatusCode != 0 {
		err = errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
		return
	}

	return
}
