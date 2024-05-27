package graph

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
)

type Resolver struct{}

// CreateArticle is the resolver for the CreateArticle field.
func (r *mutationResolver) CreateArticle(ctx context.Context, input NewArticle) (*Article, error) {
	panic("not implemented")
}

// CreateComment is the resolver for the CreateComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input NewComment) (*Comment, error) {
	panic("not implemented")
}

// TurnOfComments is the resolver for the TurnOfComments field.
func (r *mutationResolver) TurnOfComments(ctx context.Context, input SwitchCommentsAbility) (*Article, error) {
	panic("not implemented")
}

// GetArticles is the resolver for the GetArticles field.
func (r *queryResolver) GetArticles(ctx context.Context, limit *int, offset *int) ([]*Article, error) {
	panic("not implemented")
}

// GetArticle is the resolver for the GetArticle field.
func (r *queryResolver) GetArticle(ctx context.Context) (*Article, error) {
	panic("not implemented")
}

// GetMoreComments is the resolver for the GetMoreComments field.
func (r *queryResolver) GetMoreComments(ctx context.Context, limit *int, offset *int) ([]*Comment, error) {
	panic("not implemented")
}

// GetChildComments is the resolver for the GetChildComments field.
func (r *queryResolver) GetChildComments(ctx context.Context, limit *int, offset *int) ([]*Comment, error) {
	panic("not implemented")
}

// NewComments is the resolver for the newComments field.
func (r *subscriptionResolver) NewComments(ctx context.Context, input NewSubscription) (<-chan *Comment, error) {
	panic("not implemented")
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
