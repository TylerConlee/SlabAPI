package zendesk

import (
	"context"

	"github.com/tylerconlee/SlabAPI/model"
)

// GetTriggers uses the preconfigured client, c, and sends a request for all
// triggers to the Zendesk API wrapper. Once it grabs the array of triggers, it
// makes sure that any pagination is handled, and converts the ticket output
// into an array of model.Trigger.
func (c *Client) GetTriggers(ctx context.Context) (output []*model.Trigger, err error) {
	return nil, nil
}

// GetTrigger uses the preconfigured client, c, and sends a request for all
// triggers to the Zendesk API wrapper. Once it grabs the specified trigger, it
// converts the output into a type of model.Trigger.
func (c *Client) GetTrigger(ctx context.Context, id int) (output *model.Trigger, err error) {
	return nil, nil
}
