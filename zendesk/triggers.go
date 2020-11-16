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
	o, err := c.client.GetTrigger(ctx, int64(id))
	var any []*model.TriggerCondition
	var all []*model.TriggerCondition

	for _, a := range o.Conditions.Any {
		c := &model.TriggerCondition{
			Field:    a.Field,
			Operator: a.Operator,
			Value:    a.Value,
		}
		any = append(any, c)
	}
	for _, a := range o.Conditions.All {
		c := &model.TriggerCondition{
			Field:    a.Field,
			Operator: a.Operator,
			Value:    a.Value,
		}
		all = append(all, c)
	}

	conditions := &model.TriggerConditions{
		Any: any,
		All: all,
	}

	output = &model.Trigger{
		ID:          int(o.ID),
		Title:       o.Title,
		CreatedAt:   o.CreatedAt.String(),
		UpdatedAt:   o.UpdatedAt.String(),
		Position:    int(o.Position),
		Active:      o.Active,
		Description: o.Description,
		Conditions:  conditions,
	}
	return
}
