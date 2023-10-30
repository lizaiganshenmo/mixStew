namespace go follow

struct BaseResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct FollowReq {
    1:i64 uid
    2:i64 follow_uid
}

struct FollowResp {
    1:BaseResp base_resp
}

struct IsFollorResp {
    1:bool is_following
    2:BaseResp base_resp
}

struct FollowListReq {
    1:i64 uid
}

struct FollowListResp {
    1:list<i64> uids
    2:BaseResp base_resp
}

service FollowService{
    FollowResp Follow(1:FollowReq req)
    FollowResp UnFollow(1:FollowReq req)
    IsFollorResp IsFollow(1:FollowReq req)
    FollowListResp FollowList(1:FollowListReq req)
}