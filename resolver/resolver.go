package resolver

//go:generate go run github.com/99designs/gqlgen
import (
	"github.com/tylerconlee/SlabAPI/graph"
	l "github.com/tylerconlee/SlabAPI/log"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
var (
	log = l.Log
)

// Resolver references a generic resolution to a query made against the API
type Resolver struct{}

// Query contains everything that runs during a general query against the API
func (r *Resolver) Query() graph.QueryResolver {

	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }
