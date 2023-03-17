package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"fmt"

	"github.com/humgal/art-server/graph/model"
)

// PlaceBid is the resolver for the placeBid field.
func (r *mutationResolver) PlaceBid(ctx context.Context, bid *model.Bid) (*string, error) {
	panic(fmt.Errorf("not implemented: PlaceBid - placeBid"))
}

// UploadArt is the resolver for the uploadArt field.
func (r *mutationResolver) UploadArt(ctx context.Context, items []*model.UploadItem) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: UploadArt - uploadArt"))
}

// CreateCollection is the resolver for the createCollection field.
func (r *mutationResolver) CreateCollection(ctx context.Context, param model.CollectionParm) (*model.Collection, error) {
	panic(fmt.Errorf("not implemented: CreateCollection - createCollection"))
}

// SearchItems is the resolver for the searchItems field.
func (r *queryResolver) SearchItems(ctx context.Context, param model.SearchParm) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: SearchItems - searchItems"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// ItemDetail is the resolver for the itemDetail field.
func (r *queryResolver) ItemDetail(ctx context.Context, id string) (*model.Item, error) {
	panic(fmt.Errorf("not implemented: ItemDetail - itemDetail"))
}

// Collection is the resolver for the collection field.
func (r *queryResolver) Collection(ctx context.Context) (*model.Collection, error) {
	panic(fmt.Errorf("not implemented: Collection - collection"))
}

// DiscoverCreator is the resolver for the discoverCreator field.
func (r *queryResolver) DiscoverCreator(ctx context.Context, typeArg *int) ([]*model.Creator, error) {
	panic(fmt.Errorf("not implemented: DiscoverCreator - discoverCreator"))
}

// CreateorProfile is the resolver for the createorProfile field.
func (r *queryResolver) CreateorProfile(ctx context.Context, name *string) (model.CreateorProfile, error) {
	panic(fmt.Errorf("not implemented: CreateorProfile - createorProfile"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Items(ctx context.Context, param model.SearchParm) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: Items - items"))
}
func (r *queryResolver) PlaceBid(ctx context.Context, bid *model.Bid) (*string, error) {
	panic(fmt.Errorf("not implemented: PlaceBid - placeBid"))
}
