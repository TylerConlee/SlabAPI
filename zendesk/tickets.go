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

	time := time.Now().AddDate(0, 0, -5).Unix()
	// Initialize first page
	opts := zendesk.TicketListOptions{
		StartTime: strconv.Itoa(int(time)),
		Sideload:  "slas",
	}

	var tickets []zendesk.Ticket
	log.Debug("Beginning GetTickets loop")

	// Send a request to Zendesk with the specified page number and
	// sort by the most recently updated ticket
	t, cursor, _, err := c.client.GetIncrementalTickets(context.Background(), &opts)

	if err != nil {
		log.Fatal("Fatal error", zap.String("Error", err.Error()))
	}
	log.Debug("Retrieved tickets from Zendesk in GetTickets loop", zap.Int("ticket_count", len(t)), zap.Int("total_count", len(tickets)))
	tickets = append(tickets, t...)
	opts.StartTime = ""
	opts.Cursor = cursor

	// Take the []zendesk.Ticket returned from the Zendesk wrapper
	// and turn it into the []*model.Ticket used in Slab
	for _, ticket := range tickets {
		var sla string
		if len(ticket.Slas.PolicyMetrics) >= 1 {
			p := ticket.Slas.PolicyMetrics[0].(map[string]interface{})
			if p["breach_at"] != nil {
				sla = p["breach_at"].(string)
			}
		}
		save := &model.Ticket{
			URL:            ticket.URL,
			ID:             int(ticket.ID),
			Createdat:      ticket.CreatedAt.String(),
			Updatedat:      ticket.UpdatedAt.String(),
			Subject:        ticket.Subject,
			Description:    ticket.Description,
			Priority:       ticket.Priority,
			Status:         ticket.Status,
			Requesterid:    int(ticket.RequesterID),
			Organizationid: int(ticket.OrganizationID),
			Groupid:        int(ticket.GroupID),
			Tags:           ticket.Tags,
			SLA:            sla,
		}
		output = append(output, save)
	}
	return
}
