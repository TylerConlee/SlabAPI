package zendesk

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/tylerconlee/SlabAPI/model"
	"github.com/tylerconlee/zendesk-go/zendesk"
)

// Client contains the instance of the Zendesk API wrapper client from
// go-zendesk/zendesk, as well as the configuration for connecting to
// Zendesk.
type Client struct {
	config *model.ZendeskConfigInput
	client *zendesk.Client
}

// Connect takes a configuration for Zendesk and uses it to create a new
// instance of a Zendesk client. c is then used as a receiver for any other
// interaction with Zendesk, as it's preconfigured, and contains an http.Client
func Connect(config *model.ZendeskConfigInput) *Client {

	var err error
	// c is the instance of a Client that gets used for all functions.
	c := &Client{config: config}
	client := http.Client{
		Timeout: time.Second * 480,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}
	c.client, err = zendesk.NewClient(&client)
	if err != nil {

	}
	c.client.SetSubdomain(config.URL)
	c.client.SetCredential(zendesk.NewAPITokenCredential(config.User, config.Apikey))
	log.Print("Zendesk credentials set. Client successfully created")
	return c
}

// GetOrganization takes the client, c, and requests the details for the
// organization provided by the context to the Zendesk API wrapper. Once it
// retreives that data from Zendesk, it converts the output into a model.
// Organization.
func (c *Client) GetOrganization(ctx context.Context, id int) (output *model.Organization, err error) {
	o, err := c.client.GetOrganization(ctx, int64(id))
	output = &model.Organization{
		URL:         o.URL,
		ID:          int(o.ID),
		Name:        o.Name,
		CreatedAt:   o.CreatedAt.String(),
		UpdatedAt:   o.UpdatedAt.String(),
		DomainNames: o.DomainNames,
		Tags:        o.Tags,
	}
	return
}
