package resolver

import (
	"context"

	"github.com/tylerconlee/SlabAPI/model"
)

func (r *queryResolver) Zendeskconfig(ctx context.Context) (*model.ZendeskConfig, error) {
	var query = "SELECT * FROM zendesk"
	con, err := db.Query(query)
	var zenconfig *model.ZendeskConfig
	for con.Next() {
		var user, apikey, url string
		err = con.Scan(&user, &apikey, &url)
		if err != nil {
			panic(err.Error())
		}
		zenconfig = &model.ZendeskConfig{User: user, Apikey: apikey, URL: url}
	}
	return zenconfig, err
}
func (r *queryResolver) Slackconfig(ctx context.Context) (*model.SlackConfig, error) {
	panic("not implemented")
}
func (r *queryResolver) Postgresconfig(ctx context.Context) (*model.PostgresConfig, error) {
	panic("not implemented")
}
