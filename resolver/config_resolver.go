package resolver

import (
	"context"
	"log"

	"github.com/tylerconlee/SlabAPI/model"
)

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
	return zenconfig, nil
}
func (r *queryResolver) Slackconfig(ctx context.Context) (*model.SlackConfig, error) {
	panic("not implemented")
}
func (r *queryResolver) Postgresconfig(ctx context.Context) (*model.PostgresConfig, error) {
	panic("not implemented")
}
