// Code generated by Kitex v0.7.2. DO NOT EDIT.

package articleservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	article "github.com/lizaiganshenmo/mixStew/kitex_gen/article"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateArticle(ctx context.Context, req *article.CreateArticleReq, callOptions ...callopt.Option) (r *article.CreateArticleResp, err error)
	UpdateArticle(ctx context.Context, req *article.UpdateArticleReq, callOptions ...callopt.Option) (r *article.UpdateArticleResp, err error)
	DeleteArticle(ctx context.Context, req *article.DeleteArticleReq, callOptions ...callopt.Option) (r *article.DeleteArticleResp, err error)
	MGetArticle(ctx context.Context, req *article.MGetArticleReq, callOptions ...callopt.Option) (r *article.MGetArticleResp, err error)
	MGetFeedArticle(ctx context.Context, req *article.MGetFeedArticleReq, callOptions ...callopt.Option) (r *article.MGetFeedArticleResp, err error)
	CommentArticle(ctx context.Context, req *article.CommentArticleReq, callOptions ...callopt.Option) (r *article.CommentArticleResp, err error)
	GetArticleComment(ctx context.Context, req *article.GetArticleCommentReq, callOptions ...callopt.Option) (r *article.GetArticleCommentResp, err error)
	DeleteComment(ctx context.Context, req *article.DeleteCommentReq, callOptions ...callopt.Option) (r *article.DeleteCommentResp, err error)
	FavoriteArticle(ctx context.Context, req *article.FavoriteArticleReq, callOptions ...callopt.Option) (r *article.FavoriteArticleResp, err error)
	UnFavoriteArticle(ctx context.Context, req *article.UnFavoriteArticleReq, callOptions ...callopt.Option) (r *article.UnFavoriteArticleResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kArticleServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kArticleServiceClient struct {
	*kClient
}

func (p *kArticleServiceClient) CreateArticle(ctx context.Context, req *article.CreateArticleReq, callOptions ...callopt.Option) (r *article.CreateArticleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateArticle(ctx, req)
}

func (p *kArticleServiceClient) UpdateArticle(ctx context.Context, req *article.UpdateArticleReq, callOptions ...callopt.Option) (r *article.UpdateArticleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateArticle(ctx, req)
}

func (p *kArticleServiceClient) DeleteArticle(ctx context.Context, req *article.DeleteArticleReq, callOptions ...callopt.Option) (r *article.DeleteArticleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteArticle(ctx, req)
}

func (p *kArticleServiceClient) MGetArticle(ctx context.Context, req *article.MGetArticleReq, callOptions ...callopt.Option) (r *article.MGetArticleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MGetArticle(ctx, req)
}

func (p *kArticleServiceClient) MGetFeedArticle(ctx context.Context, req *article.MGetFeedArticleReq, callOptions ...callopt.Option) (r *article.MGetFeedArticleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MGetFeedArticle(ctx, req)
}

func (p *kArticleServiceClient) CommentArticle(ctx context.Context, req *article.CommentArticleReq, callOptions ...callopt.Option) (r *article.CommentArticleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentArticle(ctx, req)
}

func (p *kArticleServiceClient) GetArticleComment(ctx context.Context, req *article.GetArticleCommentReq, callOptions ...callopt.Option) (r *article.GetArticleCommentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetArticleComment(ctx, req)
}

func (p *kArticleServiceClient) DeleteComment(ctx context.Context, req *article.DeleteCommentReq, callOptions ...callopt.Option) (r *article.DeleteCommentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteComment(ctx, req)
}

func (p *kArticleServiceClient) FavoriteArticle(ctx context.Context, req *article.FavoriteArticleReq, callOptions ...callopt.Option) (r *article.FavoriteArticleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteArticle(ctx, req)
}

func (p *kArticleServiceClient) UnFavoriteArticle(ctx context.Context, req *article.UnFavoriteArticleReq, callOptions ...callopt.Option) (r *article.UnFavoriteArticleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UnFavoriteArticle(ctx, req)
}
