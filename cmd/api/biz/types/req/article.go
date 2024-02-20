package req

type CreateArticleReq struct {
	Uid         int64    `json:"uid"`
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Body        string   `json:"body" validate:"required"`
	TagList     []string `json:"tag_list,omitempty"`
}

type UpdateArticleReq struct {
	Uid         int64  `json:"uid"`
	ArticleId   int64  `json:"article_id" validate:"required"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Body        string `json:"body,omitempty"`
}

type DeleteArticleReq struct {
	ArticleId int64 `json:"article_id"`
}

type MGetArticleReq struct {
	Uid        int64  `json:"uid"`
	Tag        string `json:"tag,omitempty"`
	AuthorName string `json:"author,omitempty"`
	Favorited  bool   `json:"favorited,omitempty"`
	Limit      int32  `json:"limit,omitempty" default:"20"`
	Offest     int32  `json:"offest,omitempty" default:"0"`
}

type CommentArticleReq struct {
	Uid       int64  `json:"uid"`
	ArticleId int64  `json:"article_id" validate:"required"`
	TargetUid int64  `json:"target_uid" validate:"required"`
	Body      string `json:"body" validate:"required"`
}

type GetArticleCommentReq struct {
	ArticleId int64 `json:"article_id" validate:"required"`
}

type DeleteArticleCommentReq struct {
	ArticleId int64 `json:"article_id" validate:"required"`
	CommentId int64 `json:"comment_id" validate:"required"`
}

type FavoriteArticleReq struct {
	ArticleId int64 `json:"article_id" validate:"required"`
}

type UnFavoriteArticleReq struct {
	ArticleId int64 `json:"article_id" validate:"required"`
}
