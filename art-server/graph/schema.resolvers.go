package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"fmt"

	"github.com/humgal/art-server/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (string, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// PlaceBid is the resolver for the placeBid field.
func (r *queryResolver) PlaceBid(ctx context.Context, bid *model.Bid) (*string, error) {
	panic(fmt.Errorf("not implemented: PlaceBid - placeBid"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
