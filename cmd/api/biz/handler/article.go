package handler

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/lizaiganshenmo/mixStew/cmd/api/biz/rpc"
	"github.com/lizaiganshenmo/mixStew/cmd/api/biz/types/req"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/article"
	"github.com/lizaiganshenmo/mixStew/library/errno"
	"github.com/lizaiganshenmo/mixStew/library/utils"
)

// create article
func CreateArticle(ctx context.Context, c *app.RequestContext) {
	uid := utils.GetUid(c)
	if uid == 0 {
		SendResponse(c, errno.UserNotLoginErr, nil)
		return
	}

	var req req.CreateArticleReq
	if err := c.Bind(&req); err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := ValidateFunc(c, req)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err = rpc.CreateArticle(ctx, &article.CreateArticleReq{
		Uid:         uid,
		Title:       req.Title,
		Description: req.Description,
		Body:        req.Body,
		TagList:     req.TagList,
	})
	if err != nil {
		SendResponse(c, err, nil)
		hlog.CtxWarnf(ctx, "rpc.CreateArticle fail. err: %+v. req:%+v", err, req)
		return
	}

	SendResponse(c, errno.Success, nil)

}

// update article
func UpdateArticle(ctx context.Context, c *app.RequestContext) {
	uid := utils.GetUid(c)
	if uid == 0 {
		SendResponse(c, errno.UserNotLoginErr, nil)
		return
	}

	var req req.UpdateArticleReq
	if err := c.Bind(&req); err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := ValidateFunc(c, req)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err = rpc.UpdateArticle(ctx, &article.UpdateArticleReq{
		Uid:         uid,
		ArticleId:   req.ArticleId,
		Title:       &req.Title,
		Description: &req.Description,
		Body:        &req.Body,
	})

	if err != nil {
		SendResponse(c, err, nil)
		hlog.CtxWarnf(ctx, "rpc.UpdateArticle fail. err: %+v. req:%+v", err, req)
		return
	}

	SendResponse(c, errno.Success, nil)

}

// delete article
func DeleteArticle(ctx context.Context, c *app.RequestContext) {
	uid := utils.GetUid(c)
	if uid == 0 {
		SendResponse(c, errno.UserNotLoginErr, nil)
		return
	}

	var req req.DeleteArticleReq
	if err := c.Bind(&req); err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := ValidateFunc(c, req)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err = rpc.DeleteArticle(ctx, &article.DeleteArticleReq{
		ArticleId: req.ArticleId,
		Uid:       uid,
	})
	if err != nil {
		SendResponse(c, err, nil)
		hlog.CtxWarnf(ctx, "rpc.CreateArticle fail. err: %+v. req:%+v", err, req)
		return
	}

	SendResponse(c, errno.Success, nil)

}

// get articles list
func MGetArticle(ctx context.Context, c *app.RequestContext) {
	uid := utils.GetUid(c)

	var req req.MGetArticleReq
	if err := c.BindQuery(&req); err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := ValidateFunc(c, req)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	articleReq := article.NewMGetArticleReq()
	articleReq.Uid = uid
	articleReq.Tag = &req.Tag
	articleReq.AuthorName = &req.AuthorName
	articleReq.Favorited = &req.Favorited
	if req.Limit != 0 {
		articleReq.Limit = req.Limit
	}
	if req.Offest != 0 {
		articleReq.Offest = req.Offest
	}
	articles, err := rpc.MGetArticle(ctx, articleReq)

	if err != nil {
		SendResponse(c, err, nil)
		hlog.CtxWarnf(ctx, "rpc.MGetArticle fail. err: %+v. req:%+v", err, req)
		return
	}

	SendResponse(c, errno.Success, articles)

}

// get feeed articles list
func GetFeedArticle(ctx context.Context, c *app.RequestContext) {
	uid := utils.GetUid(c)

	var req req.MGetArticleReq
	if err := c.Bind(&req); err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := ValidateFunc(c, req)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	articleReq := article.NewMGetArticleReq()
	articleReq.Uid = uid
	articleReq.Tag = &req.Tag
	articleReq.AuthorName = &req.AuthorName
	articleReq.Favorited = &req.Favorited
	if req.Limit != 0 {
		articleReq.Limit = req.Limit
	}
	if req.Offest != 0 {
		articleReq.Offest = req.Offest
	}
	articles, err := rpc.MGetFeedArticle(ctx, articleReq)

	if err != nil {
		SendResponse(c, err, nil)
		hlog.CtxWarnf(ctx, "rpc.GetFeedArticle fail. err: %+v. req:%+v", err, req)
		return
	}

	SendResponse(c, errno.Success, articles)

}

// comment article
func CommentArticle(ctx context.Context, c *app.RequestContext) {
	uid := utils.GetUid(c)

	var req req.CommentArticleReq
	if err := c.Bind(&req); err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := ValidateFunc(c, req)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err = rpc.CommentArticle(ctx, &article.CommentArticleReq{
		Uid:       uid,
		ArticleId: req.ArticleId,
		TargetUid: req.TargetUid,
		Body:      req.Body,
	})

	if err != nil {
		SendResponse(c, err, nil)
		hlog.CtxWarnf(ctx, "rpc.CommentArticle fail. err: %+v. req:%+v", err, req)
		return
	}

	SendResponse(c, errno.Success, nil)

}

// get article comment
func GetArticleComment(ctx context.Context, c *app.RequestContext) {
	// var req req.GetArticleCommentReq
	articleIdStr := c.Param("article_id")
	articleId, _ := strconv.ParseInt(articleIdStr, 10, 64)

	comments, err := rpc.GetArticleComment(ctx, &article.GetArticleCommentReq{
		ArticleId: articleId,
	})

	if err != nil {
		SendResponse(c, err, nil)
		hlog.CtxWarnf(ctx, "rpc.GetArticleComment fail. err: %+v. articleId:%+v", err, articleId)
		return
	}

	SendResponse(c, errno.Success, comments)

}

// delete article comment
func DeleteArticleComment(ctx context.Context, c *app.RequestContext) {
	uid := utils.GetUid(c)
	if uid == 0 {
		SendResponse(c, errno.UserNotLoginErr, nil)
		return
	}

	var req req.DeleteArticleCommentReq
	if err := c.Bind(&req); err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := ValidateFunc(c, req)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err = rpc.DeleteArticleComment(ctx, &article.DeleteCommentReq{
		Uid:       uid,
		ArticleId: req.ArticleId,
		CommentId: req.CommentId,
	})

	if err != nil {
		SendResponse(c, err, nil)
		hlog.CtxWarnf(ctx, "rpc.DeleteArticleComment fail. err: %+v. req:%+v", err, req)
		return
	}

	SendResponse(c, errno.Success, nil)

}

// favorite article
func FavoriteArticle(ctx context.Context, c *app.RequestContext) {
	uid := utils.GetUid(c)
	if uid == 0 {
		SendResponse(c, errno.UserNotLoginErr, nil)
		return
	}

	var req req.FavoriteArticleReq
	if err := c.Bind(&req); err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := ValidateFunc(c, req)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err = rpc.FavoriteArticle(ctx, &article.FavoriteArticleReq{
		ArticleId: req.ArticleId,
		Uid:       uid,
	})

	if err != nil {
		SendResponse(c, err, nil)
		hlog.CtxWarnf(ctx, "rpc.FavoriteArticle fail. err: %+v. req:%+v", err, req)
		return
	}

	SendResponse(c, errno.Success, nil)

}

// unfavorite article
func UnFavoriteArticle(ctx context.Context, c *app.RequestContext) {
	uid := utils.GetUid(c)
	if uid == 0 {
		SendResponse(c, errno.UserNotLoginErr, nil)
		return
	}

	var req req.UnFavoriteArticleReq
	if err := c.Bind(&req); err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := ValidateFunc(c, req)
	if err != nil {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err = rpc.UnFavoriteArticle(ctx, &article.UnFavoriteArticleReq{
		ArticleId: req.ArticleId,
		Uid:       uid,
	})

	if err != nil {
		SendResponse(c, err, nil)
		hlog.CtxWarnf(ctx, "rpc.UnFavoriteArticle fail. err: %+v. req:%+v", err, req)
		return
	}

	SendResponse(c, errno.Success, nil)

}
