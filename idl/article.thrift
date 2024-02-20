namespace go article

struct BaseResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct Author {
    1:string username
    2:string bio
    3:string image
    4:bool following
}

struct Article {
    1:i64 article_id
    2:string title
    3:string description
    4:string body
    5:list<string> tag_list
    6:string created_at
    7:string updated_at
    8:bool favorited
    9:i32 favoritesCount
    10:Author author
}

struct Comment {
    1:i64 id
    2:string created_at
    3:string updated_at
    4:string body
    5:Author author
}

struct CreateArticleReq {
    1:i64 uid
    2:string title
    3:string description
    4:string body
    5:optional list<string> tag_list
}

struct CreateArticleResp {
    1:Article article
    2:BaseResp base_resp
}

struct UpdateArticleReq {
    1:i64 uid
    2:i64 article_id
    3:optional string title
    4:optional string description
    5:optional string body
}

struct UpdateArticleResp {
    1:BaseResp base_resp
}

struct DeleteArticleReq {
    1:i64 uid
    2:i64 article_id
}

struct DeleteArticleResp {
    1:BaseResp base_resp
}

struct MGetArticleReq {
    1:i64 uid
    2:optional string tag
    3:optional string author_name
    4:optional bool favorited
    5:optional i32 limit = 20
    6:optional i32 offest = 0
}

struct MGetArticleResp {
    1:list<Article> articles
    2:BaseResp base_resp
}

struct MGetFeedArticleReq {
    1:i64 uid
    2:optional i32 limit = 20
    3:optional i32 offest = 0
}

struct MGetFeedArticleResp {
    1:list<Article> articles
    2:BaseResp base_resp
}

struct CommentArticleReq {
    1:i64 uid
    2:i64 article_id
    3:i64 target_uid
    4:string body
}

struct CommentArticleResp {
    1:Comment comment
    2:BaseResp base_resp
}

struct GetArticleCommentReq {
    1:i64 article_id
}

struct GetArticleCommentResp {
    1:list<Comment> comments
    2:BaseResp base_resp
}

struct DeleteCommentReq {
    1:i64 uid
    2:i64 article_id
    3:i64 comment_id
}

struct DeleteCommentResp {
    1:BaseResp base_resp
}

struct FavoriteArticleReq {
    1:i64 uid
    2:i64 article_id
}

struct FavoriteArticleResp {
    1:BaseResp base_resp
}

struct UnFavoriteArticleReq{
    1:i64 uid
    2:i64 article_id
}

struct UnFavoriteArticleResp {
    1:BaseResp base_resp
}

service ArticleService {
    CreateArticleResp CreateArticle(1:CreateArticleReq req)
    UpdateArticleResp UpdateArticle(1:UpdateArticleReq req)
    DeleteArticleResp DeleteArticle(1:DeleteArticleReq req)
    MGetArticleResp MGetArticle(1:MGetArticleReq req)
    MGetFeedArticleResp MGetFeedArticle(1:MGetFeedArticleReq req)
    CommentArticleResp CommentArticle(1:CommentArticleReq req)
    GetArticleCommentResp GetArticleComment(1:GetArticleCommentReq req)
    DeleteCommentResp DeleteComment(1:DeleteCommentReq req)
    FavoriteArticleResp FavoriteArticle(1:FavoriteArticleReq req)
    UnFavoriteArticleResp UnFavoriteArticle(1:UnFavoriteArticleReq req)
}
