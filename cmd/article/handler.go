package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lizaiganshenmo/mixStew/cmd/article/pack"
	"github.com/lizaiganshenmo/mixStew/cmd/article/service"

	article "github.com/lizaiganshenmo/mixStew/kitex_gen/article"
	"github.com/lizaiganshenmo/mixStew/library/errno"
)

// ArticleServiceImpl implements the last service interface defined in the IDL.
type ArticleServiceImpl struct{}

// CreateArticle implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) CreateArticle(ctx context.Context, req *article.CreateArticleReq) (resp *article.CreateArticleResp, err error) {
	resp = new(article.CreateArticleResp)

	if req.Uid == 0 || req.Title == "" || req.Description == "" || req.Body == "" {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewArticleService(ctx).CreateArticle(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "article.CreateArticle fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// UpdateArticle implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) UpdateArticle(ctx context.Context, req *article.UpdateArticleReq) (resp *article.UpdateArticleResp, err error) {
	resp = new(article.UpdateArticleResp)

	if req.Uid == 0 || req.ArticleId == 0 || (req.Title == nil && req.Description == nil || req.Body == nil) {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewArticleService(ctx).UpdateArticle(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "article.UpdateArticle fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// DeleteArticle implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) DeleteArticle(ctx context.Context, req *article.DeleteArticleReq) (resp *article.DeleteArticleResp, err error) {
	resp = new(article.DeleteArticleResp)
	if req.Uid == 0 || req.ArticleId == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewArticleService(ctx).DeleteArticle(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "article.DeleteArticle fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// MGetArticle implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) MGetArticle(ctx context.Context, req *article.MGetArticleReq) (resp *article.MGetArticleResp, err error) {
	// Returns most recent articles globally by default, provide tag, author or favorited query parameter to filter results
	resp = new(article.MGetArticleResp)
	articleList, err := service.NewArticleService(ctx).MGetArticle(req)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "article.MGetArticle fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Articles = articleList
	return
}

// MGetFeedArticle implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) MGetFeedArticle(ctx context.Context, req *article.MGetFeedArticleReq) (resp *article.MGetFeedArticleResp, err error) {
	// return multiple articles created by followed users, ordered by most recent first.
	resp = new(article.MGetFeedArticleResp)
	articleList, err := service.NewArticleService(ctx).MGetFeedArticle(req)

	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "article.MGetFeedArticle fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Articles = articleList
	return
}

// CommentArticle implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) CommentArticle(ctx context.Context, req *article.CommentArticleReq) (resp *article.CommentArticleResp, err error) {
	resp = new(article.CommentArticleResp)

	err = service.NewArticleService(ctx).CommentArticle(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "article.CommentArticle fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// GetArticleComment implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) GetArticleComment(ctx context.Context, req *article.GetArticleCommentReq) (resp *article.GetArticleCommentResp, err error) {
	resp = new(article.GetArticleCommentResp)

	comments, err := service.NewArticleService(ctx).GetArticleComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "article.GetArticleComment fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Comments = comments
	return
}

// DeleteComment implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) DeleteComment(ctx context.Context, req *article.DeleteCommentReq) (resp *article.DeleteCommentResp, err error) {
	resp = new(article.DeleteCommentResp)

	err = service.NewArticleService(ctx).DeleteComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "article.DeleteComment fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// FavoriteArticle implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) FavoriteArticle(ctx context.Context, req *article.FavoriteArticleReq) (resp *article.FavoriteArticleResp, err error) {
	resp = new(article.FavoriteArticleResp)
	if req.Uid == 0 || req.ArticleId == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}

	err = service.NewArticleService(ctx).Favourite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "article.FavoriteArticle fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// UnFavoriteArticle implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) UnFavoriteArticle(ctx context.Context, req *article.UnFavoriteArticleReq) (resp *article.UnFavoriteArticleResp, err error) {
	resp = new(article.UnFavoriteArticleResp)
	if req.Uid == 0 || req.ArticleId == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}

	err = service.NewArticleService(ctx).UnFavourite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.CtxWarnf(ctx, "article.UnFavoriteArticle fail. err : %+v, req: %+v", err, req)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}
