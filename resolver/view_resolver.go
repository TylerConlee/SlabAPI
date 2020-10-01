package resolver

import (
	"context"

	"github.com/tylerconlee/SlabAPI/model"
	"github.com/tylerconlee/SlabAPI/zendesk"
)

// *** GET View functions *** //
// GetView retreives the specified view from Zendesk, and uses the
// ZendeskConfig object to authenticate the request.
// Endpoint: /views/{id}.json
func (r *queryResolver) GetView(ctx context.Context, config model.ZendeskConfigInput) (*model.View, error) {
	c = zendesk.Connect(&config)
	output, err := c.GetView(ctx)
	if err != nil {
		return nil, err
	}
	view := &model.View{}

	return view, nil
}
