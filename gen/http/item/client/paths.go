// Code generated by goa v3.12.4, DO NOT EDIT.
//
// HTTP request path constructors for the item service.
//
// Command:
// $ goa gen game-service/design

package client

import (
	"fmt"
)

// ListItemPath returns the URL path to the item service list HTTP endpoint.
func ListItemPath() string {
	return "/back/items"
}

// ShowItemPath returns the URL path to the item service show HTTP endpoint.
func ShowItemPath(id string) string {
	return fmt.Sprintf("/back/items/%v", id)
}

// AddItemPath returns the URL path to the item service add HTTP endpoint.
func AddItemPath() string {
	return "/back/items"
}

// UpdateItemPath returns the URL path to the item service update HTTP endpoint.
func UpdateItemPath(id string) string {
	return fmt.Sprintf("/back/items/%v", id)
}

// RemoveItemPath returns the URL path to the item service remove HTTP endpoint.
func RemoveItemPath(id string) string {
	return fmt.Sprintf("/back/items/%v", id)
}