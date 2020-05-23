// Code generated by goa v3.1.2, DO NOT EDIT.
//
// getbalance client
//
// Command:
// $ goa gen github.com/wondenge/momo-go/remittance/design

package getbalance

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Client is the "getbalance" service client.
type Client struct {
	ShowEndpoint endpoint.Endpoint
}

// NewClient initializes a "getbalance" service client given the endpoints.
func NewClient(show endpoint.Endpoint) *Client {
	return &Client{
		ShowEndpoint: show,
	}
}

// Show calls the "show" endpoint of the "getbalance" service.
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *Errorreason, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Errorreason), nil
}