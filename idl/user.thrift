namespace go user

struct BaseResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct User{
    1:CreateUserReq user_register
    2:string bio
    3:string image
}

struct UserAuth{
    1:string email
    2:string password
}

struct CreateUserReq{
    1:UserAuth user_auth
    2:string username
}

struct CreateUserResp{
    1:BaseResp base_resp
}

struct UpdateUserReq{
    1:string email
    2:string bio
    3:string image
    4:string username
    5:string password
}

struct UpdateUserResp{
    1:BaseResp base_resp
}

struct GetUserReq{
    1:i64 uid
}

struct GetUserResp{
    1:User user
    2:BaseResp base_resp
}

struct CheckUserReq{
    1:UserAuth user_Auth
}

struct CheckUserResp{
    1:i64 uid
    2:BaseResp base_resp
}

service UserService {
    CreateUserResp CreateUser(1:CreateUserReq req)
    GetUserResp GetUser(1:GetUserReq req)
    CheckUserResp CheckUser(1:CheckUserReq req)
    UpdateUserResp UpdateUser(1:UpdateUserReq req)
}
