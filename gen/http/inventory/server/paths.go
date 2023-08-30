// Code generated by goa v3.12.4, DO NOT EDIT.
//
// HTTP request path constructors for the inventory service.
//
// Command:
// $ goa gen game-service/design

package server

import (
	"fmt"
)

// ShowInventoryPath returns the URL path to the inventory service show HTTP endpoint.
func ShowInventoryPath(id string) string {
	return fmt.Sprintf("/back/inventory/%v", id)
}

// AddInventoryPath returns the URL path to the inventory service add HTTP endpoint.
func AddInventoryPath() string {
	return "/back/inventory"
}

// RemoveInventoryPath returns the URL path to the inventory service remove HTTP endpoint.
func RemoveInventoryPath() string {
	return "/back/inventory"
}
