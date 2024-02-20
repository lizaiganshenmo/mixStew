// Code generated by Kitex v0.7.2. DO NOT EDIT.

package articleservice

import (
	server "github.com/cloudwego/kitex/server"
	article "github.com/lizaiganshenmo/mixStew/kitex_gen/article"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler article.ArticleService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
