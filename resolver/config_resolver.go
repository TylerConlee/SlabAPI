package resolver

import (
	"context"
	"log"

	"github.com/tylerconlee/SlabAPI/model"
)

// ***** GET config functions ***** //
func (r *queryResolver) Zendeskconfig(ctx context.Context) (*model.ZendeskConfig, error) {

	con, err := db.Query(`SELECT name, apikey, url FROM zendesk`)
	if err != nil {
		log.Fatal(err.Error())
	}
	var zenconfig *model.ZendeskConfig
	for con.Next() {
		var user, apikey, url string
		err = con.Scan(&user, &apikey, &url)
		if err != nil {
			panic(err.Error())
		}
		zenconfig = &model.ZendeskConfig{User: user, Apikey: apikey, URL: url}
	}
	defer db.Close()
	return zenconfig, nil
}
func (r *queryResolver) Slackconfig(ctx context.Context) (*model.SlackConfig, error) {
	panic("not implemented")
}
func (r *queryResolver) Postgresconfig(ctx context.Context) (*model.PostgresConfig, error) {
	panic("not implemented")
}

// ***** UPDATE (PUT) config functions ***** //

func (r *mutationResolver) UpdateZendeskConfig(ctx context.Context, user string, apikey string, url string) (*model.ZendeskConfig, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateSlackConfig(ctx context.Context, apikey string, channel string) (*model.SlackConfig, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdatePostgresConfig(ctx context.Context, host string, port int, user string, password string, dbname string) (*model.PostgresConfig, error) {
	panic("not implemented")
}
