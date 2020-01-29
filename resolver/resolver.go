package resolver

import (
	"context"

	"github.com/tylerconlee/SlabAPI/datastore"
	"github.com/tylerconlee/SlabAPI/graph"
	"github.com/tylerconlee/SlabAPI/model"
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

func (r *mutationResolver) UpdateZendeskConfig(ctx context.Context, user string, apikey string, url string) (*model.ZendeskConfig, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateSlackConfig(ctx context.Context, apikey string, channel string) (*model.SlackConfig, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdatePostgresConfig(ctx context.Context, host string, port int, user string, password string, dbname string) (*model.PostgresConfig, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }
