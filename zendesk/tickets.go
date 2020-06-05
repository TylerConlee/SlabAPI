package zendesk

import (
	"context"
	"strconv"
	"time"

	"github.com/tylerconlee/SlabAPI/model"
	"github.com/tylerconlee/zendesk-go/zendesk"
	"go.uber.org/zap"
)

// GetTickets uses the preconfigured client, c, and sends a request for all
// tickets to the Zendesk API wrapper, asking it to sort tickets by last
// updated. Once it grabs the array of tickets, it makes sure that any
// pagination is handled, and converts the ticket output into an array of model.
// Ticket.
func (c *Client) GetTickets(ctx context.Context) (output []*model.Ticket, err error) {

	t := time.Now().AddDate(0, 0, -5).Unix()
	// Initialize first page
	opts := zendesk.TicketListOptions{
		StartTime: strconv.Itoa(int(t)),
	}

	var tickets []zendesk.Ticket
	log.Debug("Beginning GetTickets loop")
	// Loop through all pages of API response
	for {

		// Send a request to Zendesk with the specified page number and
		// sort by the most recently updated ticket
		t, _, eos, err := c.client.GetIncrementalTickets(context.Background(), &opts)

		if err != nil {
			log.Fatal("Fatal error", zap.String("Error", err.Error()))
		}
		log.Debug("Retrieved tickets from Zendesk in GetTickets loop", zap.Int("ticket_count", len(t)), zap.Int("total_count", len(tickets)))
		tickets = append(tickets, t...)

		if !eos {
			log.Debug("Reached end of GetTickets loop", zap.Int("total_count", len(tickets)))
			break
		}

	}
	// Take the []zendesk.Ticket returned from the Zendesk wrapper
	// and turn it into the []*model.Ticket used in Slab
	for _, ticket := range tickets {
		save := &model.Ticket{
			URL:            ticket.URL,
			ID:             int(ticket.ID),
			CreatedAt:      ticket.CreatedAt.String(),
			UpdatedAt:      ticket.UpdatedAt.String(),
			Subject:        ticket.Subject,
			Description:    ticket.Description,
			Priority:       ticket.Priority,
			Status:         ticket.Status,
			RequesterID:    int(ticket.RequesterID),
			OrganizationID: int(ticket.OrganizationID),
			GroupID:        int(ticket.GroupID),
			Tags:           ticket.Tags,
		}
		output = append(output, save)
	}
	return
}
