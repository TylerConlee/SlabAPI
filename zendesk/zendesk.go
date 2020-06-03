package zendesk

import (
	"context"
	"fmt"
	"log"
	"os"

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
// pagination is handled, and converts the ticket output into an array of model.
// Ticket.
func (c *Client) GetTickets(ctx context.Context) (output []*model.Ticket, err error) {
	// Initialize first page
	pageNum := 1
	var tickets []zendesk.Ticket
	// Loop through all pages of API response
	for {
		// Send a request to Zendesk with the specified page number and
		// sort by the most recently updated ticket
		t, page, err := c.client.GetTickets(context.Background(), &zendesk.TicketListOptions{
			PageOptions: zendesk.PageOptions{
				Page: pageNum,
			},
			SortBy: "updated_at",
		})

		if err != nil {
			fmt.Errorf("[E] %v", err)
			os.Exit(1)
		}

		tickets = append(tickets, t...)

		if !page.HasNext() {
			break
		}

		pageNum = pageNum + 1
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

// GetOrganization takes the client, c, and requests the details for the
// organization provided by the context to the Zendesk API wrapper. Once it
// retreives that data from Zendesk, it converts the output into a model.
// Organization.
func (c *Client) GetOrganization(ctx context.Context, id int) (output *model.Organization, err error) {
	o, err := c.client.GetOrganization(ctx, int64(id))
	output = &model.Organization{
		URL:         o.URL,
		ID:          int(o.ID),
		Name:        o.Name,
		CreatedAt:   o.CreatedAt.String(),
		UpdatedAt:   o.UpdatedAt.String(),
		DomainNames: o.DomainNames,
		Tags:        o.Tags,
	}
	return
}
