package resolver

import (
	"github.com/tylerconlee/SlabAPI/datastore"
	"github.com/tylerconlee/SlabAPI/graph"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
var db *datastore.Db

type Resolver struct{}

func (r *Resolver) Mutation() graph.MutationResolver {
	db, _ = datastore.New(
		datastore.ConnString(),
	)
	return &mutationResolver{r}
}
func (r *Resolver) Query() graph.QueryResolver {
	db, _ = datastore.New(
		datastore.ConnString(),
	)
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
