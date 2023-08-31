// Code generated by goa v3.12.4, DO NOT EDIT.
//
// HTTP request path constructors for the front service.
//
// Command:
// $ goa gen game-service/design

package server

import (
	"fmt"
)

// ListCharactersFrontPath returns the URL path to the front service list-characters HTTP endpoint.
func ListCharactersFrontPath() string {
	return "/characters"
}

// ShowCharacterFrontPath returns the URL path to the front service show-character HTTP endpoint.
func ShowCharacterFrontPath(id string) string {
	return fmt.Sprintf("/characters/%v", id)
}

// AddCharacterFrontPath returns the URL path to the front service add-character HTTP endpoint.
func AddCharacterFrontPath() string {
	return "/characters"
}

// UpdateCharacterFrontPath returns the URL path to the front service update-character HTTP endpoint.
func UpdateCharacterFrontPath(id string) string {
	return fmt.Sprintf("/characters/%v", id)
}

// RemoveCharacterFrontPath returns the URL path to the front service remove-character HTTP endpoint.
func RemoveCharacterFrontPath(id string) string {
	return fmt.Sprintf("/characters/%v", id)
}

// AddItemFrontPath returns the URL path to the front service add-item HTTP endpoint.
func AddItemFrontPath() string {
	return "/"
}

// RemoveItemFrontPath returns the URL path to the front service remove-item HTTP endpoint.
func RemoveItemFrontPath() string {
	return "/"
}
