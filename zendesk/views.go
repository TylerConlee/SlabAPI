package zendesk

import (
	"context"

	"github.com/tylerconlee/SlabAPI/model"
	"go.uber.org/zap"
)

// GetViews takes the client, c, and requests the details for the
// organization provided by the context to the Zendesk API wrapper. Once it
// retreives that data from Zendesk, it converts the output into a model.
// Organization.
func (c *Client) GetViews(ctx context.Context) ([]*model.View, error) {
	var output []*model.View
	o, _, err := c.client.GetViews(ctx)
	if err != nil {
		log.Error("Error retrieving views", zap.String("Error", err.Error()))
		return nil, err
	}

	for _, v := range o {
		output = append(output, &model.View{
			ID:        int(v.ID),
			Title:     v.Title,
			CreatedAt: v.CreatedAt.String(),
			UpdatedAt: v.UpdatedAt.String(),
		})
	}
	return output, nil
}
