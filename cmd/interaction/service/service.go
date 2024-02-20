package service

import "context"

type InteractionService struct {
	ctx context.Context
}

// InteractionService new InteractionService
func NewInteractionService(ctx context.Context) *InteractionService {
	return &InteractionService{ctx: ctx}
}
