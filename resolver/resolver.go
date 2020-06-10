package resolver

//go:generate go run github.com/99designs/gqlgen
import (
	"github.com/tylerconlee/SlabAPI/datastore"
	"github.com/tylerconlee/SlabAPI/graph"
	l "github.com/tylerconlee/slab/log"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
var (
	db  *datastore.Db
	log = l.Log
)

// Resolver references a generic resolution to a query made against the API
type Resolver struct{}

// Mutation contains everything that runs during a mutation query against the API
func (r *Resolver) Mutation() graph.MutationResolver {
	db, _ = datastore.New(
		datastore.ConnString(),
	)
	return &mutationResolver{r}
}

// Query contains everything that runs during a general query against the API
func (r *Resolver) Query() graph.QueryResolver {
	db, _ = datastore.New(
		datastore.ConnString(),
	)

	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
