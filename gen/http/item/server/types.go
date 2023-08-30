// Code generated by goa v3.12.4, DO NOT EDIT.
//
// item HTTP server types
//
// Command:
// $ goa gen game-service/design

package server

import (
	item "game-service/gen/item"
	itemviews "game-service/gen/item/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "item" service "add" endpoint HTTP request
// body.
type AddRequestBody struct {
	// Name of item
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description of item
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Amount of damage the item can do
	Damage *uint32 `form:"damage,omitempty" json:"damage,omitempty" xml:"damage,omitempty"`
	// Amount of healing the item can provide
	Healing *uint32 `form:"healing,omitempty" json:"healing,omitempty" xml:"healing,omitempty"`
	// Amount of protection the item can provide
	Protection *uint32 `form:"protection,omitempty" json:"protection,omitempty" xml:"protection,omitempty"`
}

// UpdateRequestBody is the type of the "item" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// Item with updated fields
	Item *ItemRequestBody `form:"item,omitempty" json:"item,omitempty" xml:"item,omitempty"`
}

// StoredItemResponseCollection is the type of the "item" service "list"
// endpoint HTTP response body.
type StoredItemResponseCollection []*StoredItemResponse

// ShowResponseBody is the type of the "item" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID is the unique id of the item.
	ID string `form:"id" json:"id" xml:"id"`
	// Name of item
	Name string `form:"name" json:"name" xml:"name"`
	// Description of item
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Amount of damage the item can do
	Damage uint32 `form:"damage" json:"damage" xml:"damage"`
	// Amount of healing the item can provide
	Healing uint32 `form:"healing" json:"healing" xml:"healing"`
	// Amount of protection the item can provide
	Protection uint32 `form:"protection" json:"protection" xml:"protection"`
}

// ShowResponseBodyTiny is the type of the "item" service "show" endpoint HTTP
// response body.
type ShowResponseBodyTiny struct {
	// ID is the unique id of the item.
	ID string `form:"id" json:"id" xml:"id"`
	// Name of item
	Name string `form:"name" json:"name" xml:"name"`
}

// ShowNotFoundResponseBody is the type of the "item" service "show" endpoint
// HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// not found
	Message string `form:"message" json:"message" xml:"message"`
	// ID of missing item
	ID string `form:"id" json:"id" xml:"id"`
}

// UpdateNotFoundResponseBody is the type of the "item" service "update"
// endpoint HTTP response body for the "not_found" error.
type UpdateNotFoundResponseBody struct {
	// not found
	Message string `form:"message" json:"message" xml:"message"`
	// ID of missing item
	ID string `form:"id" json:"id" xml:"id"`
}

// RemoveNotFoundResponseBody is the type of the "item" service "remove"
// endpoint HTTP response body for the "not_found" error.
type RemoveNotFoundResponseBody struct {
	// not found
	Message string `form:"message" json:"message" xml:"message"`
	// ID of missing item
	ID string `form:"id" json:"id" xml:"id"`
}

// StoredItemResponse is used to define fields on response body types.
type StoredItemResponse struct {
	// ID is the unique id of the item.
	ID string `form:"id" json:"id" xml:"id"`
	// Name of item
	Name string `form:"name" json:"name" xml:"name"`
	// Description of item
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Amount of damage the item can do
	Damage uint32 `form:"damage" json:"damage" xml:"damage"`
	// Amount of healing the item can provide
	Healing uint32 `form:"healing" json:"healing" xml:"healing"`
	// Amount of protection the item can provide
	Protection uint32 `form:"protection" json:"protection" xml:"protection"`
}

// ItemRequestBody is used to define fields on request body types.
type ItemRequestBody struct {
	// Name of item
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description of item
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Amount of damage the item can do
	Damage *uint32 `form:"damage,omitempty" json:"damage,omitempty" xml:"damage,omitempty"`
	// Amount of healing the item can provide
	Healing *uint32 `form:"healing,omitempty" json:"healing,omitempty" xml:"healing,omitempty"`
	// Amount of protection the item can provide
	Protection *uint32 `form:"protection,omitempty" json:"protection,omitempty" xml:"protection,omitempty"`
}

// NewStoredItemResponseCollection builds the HTTP response body from the
// result of the "list" endpoint of the "item" service.
func NewStoredItemResponseCollection(res itemviews.StoredItemCollectionView) StoredItemResponseCollection {
	body := make([]*StoredItemResponse, len(res))
	for i, val := range res {
		body[i] = marshalItemviewsStoredItemViewToStoredItemResponse(val)
	}
	return body
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "item" service.
func NewShowResponseBody(res *itemviews.StoredItemView) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:          *res.ID,
		Name:        *res.Name,
		Description: res.Description,
		Damage:      *res.Damage,
		Healing:     *res.Healing,
		Protection:  *res.Protection,
	}
	return body
}

// NewShowResponseBodyTiny builds the HTTP response body from the result of the
// "show" endpoint of the "item" service.
func NewShowResponseBodyTiny(res *itemviews.StoredItemView) *ShowResponseBodyTiny {
	body := &ShowResponseBodyTiny{
		ID:   *res.ID,
		Name: *res.Name,
	}
	return body
}

// NewShowNotFoundResponseBody builds the HTTP response body from the result of
// the "show" endpoint of the "item" service.
func NewShowNotFoundResponseBody(res *item.NotFound) *ShowNotFoundResponseBody {
	body := &ShowNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewUpdateNotFoundResponseBody builds the HTTP response body from the result
// of the "update" endpoint of the "item" service.
func NewUpdateNotFoundResponseBody(res *item.NotFound) *UpdateNotFoundResponseBody {
	body := &UpdateNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewRemoveNotFoundResponseBody builds the HTTP response body from the result
// of the "remove" endpoint of the "item" service.
func NewRemoveNotFoundResponseBody(res *item.NotFound) *RemoveNotFoundResponseBody {
	body := &RemoveNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewShowPayload builds a item service show endpoint payload.
func NewShowPayload(id string, view *string) *item.ShowPayload {
	v := &item.ShowPayload{}
	v.ID = id
	v.View = view

	return v
}

// NewAddItem builds a item service add endpoint payload.
func NewAddItem(body *AddRequestBody) *item.Item {
	v := &item.Item{
		Name:        *body.Name,
		Description: body.Description,
		Damage:      *body.Damage,
		Healing:     *body.Healing,
		Protection:  *body.Protection,
	}

	return v
}

// NewUpdatePayload builds a item service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, id string) *item.UpdatePayload {
	v := &item.UpdatePayload{}
	v.Item = unmarshalItemRequestBodyToItemItem(body.Item)
	v.ID = id

	return v
}

// NewRemovePayload builds a item service remove endpoint payload.
func NewRemovePayload(id string) *item.RemovePayload {
	v := &item.RemovePayload{}
	v.ID = id

	return v
}

// ValidateAddRequestBody runs the validations defined on AddRequestBody
func ValidateAddRequestBody(body *AddRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Damage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("damage", "body"))
	}
	if body.Healing == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("healing", "body"))
	}
	if body.Protection == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("protection", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 2000, false))
		}
	}
	if body.Damage != nil {
		if *body.Damage < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.damage", *body.Damage, 0, true))
		}
	}
	if body.Damage != nil {
		if *body.Damage > 200 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.damage", *body.Damage, 200, false))
		}
	}
	if body.Healing != nil {
		if *body.Healing < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.healing", *body.Healing, 0, true))
		}
	}
	if body.Healing != nil {
		if *body.Healing > 200 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.healing", *body.Healing, 200, false))
		}
	}
	if body.Protection != nil {
		if *body.Protection < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.protection", *body.Protection, 0, true))
		}
	}
	if body.Protection != nil {
		if *body.Protection > 20 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.protection", *body.Protection, 20, false))
		}
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Item == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("item", "body"))
	}
	if body.Item != nil {
		if err2 := ValidateItemRequestBody(body.Item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateItemRequestBody runs the validations defined on ItemRequestBody
func ValidateItemRequestBody(body *ItemRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Damage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("damage", "body"))
	}
	if body.Healing == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("healing", "body"))
	}
	if body.Protection == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("protection", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 2000, false))
		}
	}
	if body.Damage != nil {
		if *body.Damage < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.damage", *body.Damage, 0, true))
		}
	}
	if body.Damage != nil {
		if *body.Damage > 200 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.damage", *body.Damage, 200, false))
		}
	}
	if body.Healing != nil {
		if *body.Healing < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.healing", *body.Healing, 0, true))
		}
	}
	if body.Healing != nil {
		if *body.Healing > 200 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.healing", *body.Healing, 200, false))
		}
	}
	if body.Protection != nil {
		if *body.Protection < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.protection", *body.Protection, 0, true))
		}
	}
	if body.Protection != nil {
		if *body.Protection > 20 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.protection", *body.Protection, 20, false))
		}
	}
	return
}
