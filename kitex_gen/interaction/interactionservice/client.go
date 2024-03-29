// Code generated by Kitex v0.7.2. DO NOT EDIT.

package interactionservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	interaction "github.com/lizaiganshenmo/mixStew/kitex_gen/interaction"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateComment(ctx context.Context, req *interaction.CreateCommentReq, callOptions ...callopt.Option) (r *interaction.CreateCommentResp, err error)
	DeleteComment(ctx context.Context, req *interaction.DeleteCommentReq, callOptions ...callopt.Option) (r *interaction.DeleteCommentResp, err error)
	GetComment(ctx context.Context, req *interaction.GetCommentReq, callOptions ...callopt.Option) (r *interaction.GetCommentResp, err error)
	CreateCommentReply(ctx context.Context, req *interaction.CreateCommentReplyReq, callOptions ...callopt.Option) (r *interaction.CreateCommentReplyResp, err error)
	DeleteCommentReply(ctx context.Context, req *interaction.DeleteCommentReplyReq, callOptions ...callopt.Option) (r *interaction.DeleteCommentReplyResp, err error)
	GetCommentReply(ctx context.Context, req *interaction.GetCommentReplyReq, callOptions ...callopt.Option) (r *interaction.GetCommentReplyResp, err error)
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
	return &kInteractionServiceClient{
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

type kInteractionServiceClient struct {
	*kClient
}

func (p *kInteractionServiceClient) CreateComment(ctx context.Context, req *interaction.CreateCommentReq, callOptions ...callopt.Option) (r *interaction.CreateCommentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateComment(ctx, req)
}

func (p *kInteractionServiceClient) DeleteComment(ctx context.Context, req *interaction.DeleteCommentReq, callOptions ...callopt.Option) (r *interaction.DeleteCommentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteComment(ctx, req)
}

func (p *kInteractionServiceClient) GetComment(ctx context.Context, req *interaction.GetCommentReq, callOptions ...callopt.Option) (r *interaction.GetCommentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetComment(ctx, req)
}

func (p *kInteractionServiceClient) CreateCommentReply(ctx context.Context, req *interaction.CreateCommentReplyReq, callOptions ...callopt.Option) (r *interaction.CreateCommentReplyResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateCommentReply(ctx, req)
}

func (p *kInteractionServiceClient) DeleteCommentReply(ctx context.Context, req *interaction.DeleteCommentReplyReq, callOptions ...callopt.Option) (r *interaction.DeleteCommentReplyResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteCommentReply(ctx, req)
}

func (p *kInteractionServiceClient) GetCommentReply(ctx context.Context, req *interaction.GetCommentReplyReq, callOptions ...callopt.Option) (r *interaction.GetCommentReplyResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCommentReply(ctx, req)
}
