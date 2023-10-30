// Code generated by Kitex v0.7.2. DO NOT EDIT.

package followservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	follow "github.com/lizaiganshenmo/mixStew/kitex_gen/follow"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Follow(ctx context.Context, req *follow.FollowReq, callOptions ...callopt.Option) (r *follow.FollowResp, err error)
	UnFollow(ctx context.Context, req *follow.FollowReq, callOptions ...callopt.Option) (r *follow.FollowResp, err error)
	IsFollow(ctx context.Context, req *follow.FollowReq, callOptions ...callopt.Option) (r *follow.IsFollorResp, err error)
	FollowList(ctx context.Context, req *follow.FollowListReq, callOptions ...callopt.Option) (r *follow.FollowListResp, err error)
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
	return &kFollowServiceClient{
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

type kFollowServiceClient struct {
	*kClient
}

func (p *kFollowServiceClient) Follow(ctx context.Context, req *follow.FollowReq, callOptions ...callopt.Option) (r *follow.FollowResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Follow(ctx, req)
}

func (p *kFollowServiceClient) UnFollow(ctx context.Context, req *follow.FollowReq, callOptions ...callopt.Option) (r *follow.FollowResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UnFollow(ctx, req)
}

func (p *kFollowServiceClient) IsFollow(ctx context.Context, req *follow.FollowReq, callOptions ...callopt.Option) (r *follow.IsFollorResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IsFollow(ctx, req)
}

func (p *kFollowServiceClient) FollowList(ctx context.Context, req *follow.FollowListReq, callOptions ...callopt.Option) (r *follow.FollowListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowList(ctx, req)
}
