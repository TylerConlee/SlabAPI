package zendesk 

import (
	"context"

	"github.com/tylerconlee/SlabAPI/model"
)

func(c *Client) GetUser(ctx.Context) (output *model.User, err Error) {
	o, err := c.client.GetUser(ctx, int64(id))
	output = &model.User{
		Active: o.Active,
		DefaultGroup: int(o.DefaultGroupID),
		ID: int(o.ID),
		Email: o.Email,
		Name: o.Name,
		CreatedAt: o.CreatedAt.String(),
		UpdatedAt: o.UpdatedAt.String(),
		LastLoginAt: o.LastLogin.String(),
		Timezone: o.Timezone,
	}
	return
}
