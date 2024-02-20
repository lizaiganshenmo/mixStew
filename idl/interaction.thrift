namespace go interaction

struct BaseResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct Comment {
    1:i64 comment_id
    2:i64 uid
    3:i64 targert_article_id
    4:i64 targert_uid
    5:string comment_body
    6:string created_at
    7:string updated_at
}

struct CommentReply {
    1:i64 comment_id
    2:i64 uid
    3:i64 targert_article_id
    4:i64 targert_comment_id
    5:i64 targert_uid
    6:string comment_body
    7:string created_at
    8:string updated_at
}

struct CreateCommentReq {
    1:i64 uid
    2:i64 targert_article_id
    3:i64 targert_uid
    4:string comment_body
}

struct CreateCommentResp {
    1:BaseResp base_resp
}

struct DeleteCommentReq {
    1:i64 uid
    2:i64 comment_id
}

struct DeleteCommentResp {
    1:BaseResp base_resp
}

struct GetCommentReq {
    1:i64 uid
    2:i64 targert_article_id
}

struct GetCommentResp {
    1:list<Comment> comments
    2:BaseResp base_resp
}

struct CreateCommentReplyReq {
    1:i64 uid
    2:i64 targert_comment_id
    3:i64 targert_article_id
    4:i64 targert_uid
    5:string comment_body
}

struct CreateCommentReplyResp {
    1:BaseResp base_resp
}

struct DeleteCommentReplyReq {
    1:i64 uid
    2:i64 comment_id
}

struct DeleteCommentReplyResp {
    1:BaseResp base_resp
}

struct GetCommentReplyReq {
    1:i64 uid
    2:i64 targert_comment_id
}

struct GetCommentReplyResp {
    1:list<Comment> comment_replys
    2:BaseResp base_resp
}


service InteractionService {
    CreateCommentResp CreateComment(1:CreateCommentReq req)
    DeleteCommentResp DeleteComment(1:DeleteCommentReq req)
    GetCommentResp GetComment(1:GetCommentReq req)
    CreateCommentReplyResp CreateCommentReply(1:CreateCommentReplyReq req)
    DeleteCommentReplyResp DeleteCommentReply(1:DeleteCommentReplyReq req)
    GetCommentReplyResp GetCommentReply(1:GetCommentReplyReq req)
}