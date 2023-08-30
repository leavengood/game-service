// Code generated by goa v3.12.4, DO NOT EDIT.
//
// front client
//
// Command:
// $ goa gen game-service/design

package front

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "front" service client.
type Client struct {
	ListItemsEndpoint goa.Endpoint
}

// NewClient initializes a "front" service client given the endpoints.
func NewClient(listItems goa.Endpoint) *Client {
	return &Client{
		ListItemsEndpoint: listItems,
	}
}

// ListItems calls the "list-items" endpoint of the "front" service.
func (c *Client) ListItems(ctx context.Context) (res StoredItemCollection, err error) {
	var ires any
	ires, err = c.ListItemsEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(StoredItemCollection), nil
}
