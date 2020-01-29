package resolver

import (
	"context"
	"log"

	"github.com/tylerconlee/SlabAPI/datastore"
	"github.com/tylerconlee/SlabAPI/graph"
	"github.com/tylerconlee/SlabAPI/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

var db *datastore.Db

func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() graph.QueryResolver {
	db, err := datastore.New(
		datastore.ConnString(),
	)
	if err != nil {
		log.Fatal("Error connecting to Postgres database: %s", err)
	}
	defer db.Close()
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
