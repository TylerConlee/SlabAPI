package zendesk 

import (
	"context"

	"github.com/tylerconlee/SlabAPI/model"
)

// GetUser takes the client, c, and requests the details for the
// user provided by the context to the Zendesk API wrapper. Once it
// retreives that data from Zendesk, it converts the output into a model.
// user.
func (c *Client) GetUser(ctx.Context) (output *model.User, err Error) {
	o, err := c.client.GetUser(ctx, int64(id))
	output = &model.User{
		Active: o.Active,
		DefaultGroup: int(o.DefaultGroupID),
		ID: int(o.ID),
		Email: o.Email,
		Name: o.Name,
		CreatedAt: o.CreatedAt.String(),
		UpdatedAt: o.UpdatedAt.String(),
		LastLogin: o.LastLoginAt.String(),
		Timezone: o.Timezone,
	}
	return
}
