package resolver

import (
	"context"

	"github.com/tylerconlee/SlabAPI/model"
	"github.com/tylerconlee/SlabAPI/zendesk"
)

var c *zendesk.Client

// ***** GET ticket functions ***** //
// GetAllTickets takes the ZendeskConfig object of username, APIkey and URL and
// makes a request to Zendesk to the /tickets.json endpoint. This returns all
// tickets in the Tickets type, found in the schema.
func (r *queryResolver) GetAllTickets(ctx context.Context, config model.ZendeskConfigInput) (*model.Tickets, error) {
	c = zendesk.Connect(&config)
	output, err := c.GetTickets(ctx)
	if err != nil {
		return nil, err
	}
	tickets := &model.Tickets{
		Tickets: output,
		Count:   len(output),
	}

	return tickets, nil
}

// ***** GET organization functions ***** //
// GetOrganization takes the ZendeskConfig object of username, APIkey and URL,
// as well as an organization ID and makes a request to Zendesk to the /
// organization.json endpoint. This returns the information related to that
// organization.
func (r *queryResolver) GetOrganization(ctx context.Context, config model.ZendeskConfigInput, id int) (*model.Organization, error) {
	c = zendesk.Connect(&config)
	org, err := c.GetOrganization(ctx, id)
	if err != nil {
		return nil, err
	}

	return org, nil
}
