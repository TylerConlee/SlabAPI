package resolver

import (
	"context"

	"github.com/tylerconlee/SlabAPI/model"
)

// ***** GET config functions ***** //
// Zendeskconfig retrieves the Zendesk configuration from the postgres database
// and returns it as a Zendeskconfig model.
func (r *queryResolver) Zendeskconfig(ctx context.Context) (*model.ZendeskConfig, error) {
	con, err := db.Query(`SELECT zen_user, zen_apikey, zen_url FROM config`)
	if err != nil {
		log.Fatal(map[string]interface{}{
			"error": err.Error(),
		})
	}
	var zenconfig *model.ZendeskConfig
	for con.Next() {
		var user, apikey, url string
		err = con.Scan(&user, &apikey, &url)
		if err != nil {
			log.Fatal(map[string]interface{}{
				"error": err.Error(),
			})
		}
		zenconfig = &model.ZendeskConfig{User: user, Apikey: apikey, URL: url}
	}
	defer db.Close()
	return zenconfig, nil
}

// ***** UPDATE (PUT) config functions ***** //

func (r *mutationResolver) UpdateZendeskConfig(ctx context.Context, user string, apikey string, url string) (*model.ZendeskConfig, error) {
	var result model.ZendeskConfig

	conf, err := db.Prepare("UPDATE config SET zen_user = $1, zen_apikey = $2, zen_url = $3")
	if err != nil {
		log.Fatal(map[string]interface{}{
			"error": err.Error(),
		})
	}
	conf.Exec(user, apikey, url)
	result = model.ZendeskConfig{User: user, Apikey: apikey, URL: url}
	defer db.Close()
	return &result, err

}
