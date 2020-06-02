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
