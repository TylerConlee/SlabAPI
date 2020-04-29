package resolver

import (
	"context"

	"github.com/tylerconlee/SlabAPI/model"
)

// ***** GET ticket functions ***** //
// GetAllTickets takes the ZendeskConfig object of username, APIkey and URL and
// makes a request to Zendesk to the /tickets.json endpoint. This returns all
// tickets in the Tickets type, found in the schema.
func (r *queryResolver) GetAllTickets(ctx context.Context, config model.ZendeskConfigInput) (*model.Tickets, error) {
	var tickets *model.Tickets

	return tickets, nil
}
