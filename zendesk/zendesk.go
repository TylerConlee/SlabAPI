package zendesk

import (
	"context"
	"log"

	"github.com/nukosuke/go-zendesk/zendesk"
	"github.com/tylerconlee/SlabAPI/model"
)

// Client contains the instance of the Zendesk API wrapper client from
// go-zendesk/zendesk, as well as the configuration for connecting to
// Zendesk.
type Client struct {
	config *model.ZendeskConfigInput
	client *zendesk.Client
}

// Connect takes a configuration for Zendesk and uses it to create a new
// instance of a Zendesk client. c is then used as a receiver for any other
// interaction with Zendesk, as it's preconfigured, and contains an http.Client
func Connect(config *model.ZendeskConfigInput) *Client {

	var err error
	// c is the instance of a Client that gets used for all functions.
	c := &Client{config: config}
	c.client, err = zendesk.NewClient(nil)
	if err != nil {

	}
	c.client.SetSubdomain(config.URL)
	c.client.SetCredential(zendesk.NewAPITokenCredential(config.User, config.Apikey))
	log.Print("Zendesk credentials set. Client successfully created")
	return c
}

// GetTickets uses the preconfigured client, c, and sends a request for all
// tickets to the Zendesk API wrapper, asking it to sort tickets by last
// updated. Once it grabs the array of tickets, it makes sure that any
// pagination is handled, and converts the ticket output into an array of model.Ticket.
func (c *Client) GetTickets(ctx context.Context) (output []*model.Ticket, err error) {
	// Sort by most recent tickest first
	opts := &zendesk.TicketListOptions{SortBy: "updated_at"}

	// Check to make sure Zendesk Client is active
	if c.client != nil {

		// Make the request to the Zendesk API to get all tickets
		tickets, page, err := c.client.GetTickets(ctx, opts)
		// If there are additional pages, retrieve those next.
		if page.HasNext() {

		}
		// Handle any errors from the request to Zendesk
		if err != nil {
			log.Printf("Error detected while getting tickets: %s", err)
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
	} else {
		log.Print(c)
	}
	return
}
