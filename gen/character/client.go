// Code generated by goa v3.12.4, DO NOT EDIT.
//
// character client
//
// Command:
// $ goa gen game-service/design

package character

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "character" service client.
type Client struct {
	ListEndpoint   goa.Endpoint
	ShowEndpoint   goa.Endpoint
	AddEndpoint    goa.Endpoint
	UpdateEndpoint goa.Endpoint
	RemoveEndpoint goa.Endpoint
}

// NewClient initializes a "character" service client given the endpoints.
func NewClient(list, show, add, update, remove goa.Endpoint) *Client {
	return &Client{
		ListEndpoint:   list,
		ShowEndpoint:   show,
		AddEndpoint:    add,
		UpdateEndpoint: update,
		RemoveEndpoint: remove,
	}
}

// List calls the "list" endpoint of the "character" service.
func (c *Client) List(ctx context.Context) (res StoredCharacterCollection, err error) {
	var ires any
	ires, err = c.ListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(StoredCharacterCollection), nil
}

// Show calls the "show" endpoint of the "character" service.
// Show may return the following errors:
//   - "not_found" (type *NotFound): character not found
//   - error: internal error
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *StoredCharacter, err error) {
	var ires any
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StoredCharacter), nil
}

// Add calls the "add" endpoint of the "character" service.
// Add may return the following errors:
//   - "name_taken" (type *NameTaken): name is taken
//   - error: internal error
func (c *Client) Add(ctx context.Context, p *Character) (res string, err error) {
	var ires any
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Update calls the "update" endpoint of the "character" service.
// Update may return the following errors:
//   - "not_found" (type *NotFound): character not found
//   - "name_taken" (type *NameTaken): name is taken
//   - error: internal error
func (c *Client) Update(ctx context.Context, p *UpdatePayload) (err error) {
	_, err = c.UpdateEndpoint(ctx, p)
	return
}

// Remove calls the "remove" endpoint of the "character" service.
// Remove may return the following errors:
//   - "not_found" (type *NotFound): character not found
//   - error: internal error
func (c *Client) Remove(ctx context.Context, p *RemovePayload) (err error) {
	_, err = c.RemoveEndpoint(ctx, p)
	return
}