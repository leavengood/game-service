// Code generated by goa v3.12.4, DO NOT EDIT.
//
// character HTTP server types
//
// Command:
// $ goa gen game-service/design

package server

import (
	character "game-service/gen/character"
	characterviews "game-service/gen/character/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "character" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// Name of the character
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description of the character
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Amount of health the character has
	Health *uint32 `form:"health,omitempty" json:"health,omitempty" xml:"health,omitempty"`
	// Amount of experience the character has
	Experience *uint32 `form:"experience,omitempty" json:"experience,omitempty" xml:"experience,omitempty"`
}

// UpdateRequestBody is the type of the "character" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	// character with updated fields
	Character *CharacterRequestBody `form:"character,omitempty" json:"character,omitempty" xml:"character,omitempty"`
}

// StoredCharacterResponseCollection is the type of the "character" service
// "list" endpoint HTTP response body.
type StoredCharacterResponseCollection []*StoredCharacterResponse

// ShowResponseBody is the type of the "character" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID is the unique id of the character.
	ID string `form:"id" json:"id" xml:"id"`
	// Name of the character
	Name string `form:"name" json:"name" xml:"name"`
	// Description of the character
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Amount of health the character has
	Health uint32 `form:"health" json:"health" xml:"health"`
	// Amount of experience the character has
	Experience uint32 `form:"experience" json:"experience" xml:"experience"`
}

// ShowResponseBodyTiny is the type of the "character" service "show" endpoint
// HTTP response body.
type ShowResponseBodyTiny struct {
	// ID is the unique id of the character.
	ID string `form:"id" json:"id" xml:"id"`
	// Name of the character
	Name string `form:"name" json:"name" xml:"name"`
}

// ShowNotFoundResponseBody is the type of the "character" service "show"
// endpoint HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// not found
	Message string `form:"message" json:"message" xml:"message"`
	// ID of missing item
	ID string `form:"id" json:"id" xml:"id"`
}

// AddNameTakenResponseBody is the type of the "character" service "add"
// endpoint HTTP response body for the "name_taken" error.
type AddNameTakenResponseBody struct {
	// name taken
	Message string `form:"message" json:"message" xml:"message"`
	// name that is not unique
	Name string `form:"name" json:"name" xml:"name"`
}

// UpdateNameTakenResponseBody is the type of the "character" service "update"
// endpoint HTTP response body for the "name_taken" error.
type UpdateNameTakenResponseBody struct {
	// name taken
	Message string `form:"message" json:"message" xml:"message"`
	// name that is not unique
	Name string `form:"name" json:"name" xml:"name"`
}

// UpdateNotFoundResponseBody is the type of the "character" service "update"
// endpoint HTTP response body for the "not_found" error.
type UpdateNotFoundResponseBody struct {
	// not found
	Message string `form:"message" json:"message" xml:"message"`
	// ID of missing item
	ID string `form:"id" json:"id" xml:"id"`
}

// RemoveNotFoundResponseBody is the type of the "character" service "remove"
// endpoint HTTP response body for the "not_found" error.
type RemoveNotFoundResponseBody struct {
	// not found
	Message string `form:"message" json:"message" xml:"message"`
	// ID of missing item
	ID string `form:"id" json:"id" xml:"id"`
}

// StoredCharacterResponse is used to define fields on response body types.
type StoredCharacterResponse struct {
	// ID is the unique id of the character.
	ID string `form:"id" json:"id" xml:"id"`
	// Name of the character
	Name string `form:"name" json:"name" xml:"name"`
	// Description of the character
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Amount of health the character has
	Health uint32 `form:"health" json:"health" xml:"health"`
	// Amount of experience the character has
	Experience uint32 `form:"experience" json:"experience" xml:"experience"`
}

// CharacterRequestBody is used to define fields on request body types.
type CharacterRequestBody struct {
	// Name of the character
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description of the character
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Amount of health the character has
	Health *uint32 `form:"health,omitempty" json:"health,omitempty" xml:"health,omitempty"`
	// Amount of experience the character has
	Experience *uint32 `form:"experience,omitempty" json:"experience,omitempty" xml:"experience,omitempty"`
}

// NewStoredCharacterResponseCollection builds the HTTP response body from the
// result of the "list" endpoint of the "character" service.
func NewStoredCharacterResponseCollection(res characterviews.StoredCharacterCollectionView) StoredCharacterResponseCollection {
	body := make([]*StoredCharacterResponse, len(res))
	for i, val := range res {
		body[i] = marshalCharacterviewsStoredCharacterViewToStoredCharacterResponse(val)
	}
	return body
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "character" service.
func NewShowResponseBody(res *characterviews.StoredCharacterView) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:          *res.ID,
		Name:        *res.Name,
		Description: res.Description,
		Health:      *res.Health,
		Experience:  *res.Experience,
	}
	return body
}

// NewShowResponseBodyTiny builds the HTTP response body from the result of the
// "show" endpoint of the "character" service.
func NewShowResponseBodyTiny(res *characterviews.StoredCharacterView) *ShowResponseBodyTiny {
	body := &ShowResponseBodyTiny{
		ID:   *res.ID,
		Name: *res.Name,
	}
	return body
}

// NewShowNotFoundResponseBody builds the HTTP response body from the result of
// the "show" endpoint of the "character" service.
func NewShowNotFoundResponseBody(res *character.NotFound) *ShowNotFoundResponseBody {
	body := &ShowNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewAddNameTakenResponseBody builds the HTTP response body from the result of
// the "add" endpoint of the "character" service.
func NewAddNameTakenResponseBody(res *character.NameTaken) *AddNameTakenResponseBody {
	body := &AddNameTakenResponseBody{
		Message: res.Message,
		Name:    res.Name,
	}
	return body
}

// NewUpdateNameTakenResponseBody builds the HTTP response body from the result
// of the "update" endpoint of the "character" service.
func NewUpdateNameTakenResponseBody(res *character.NameTaken) *UpdateNameTakenResponseBody {
	body := &UpdateNameTakenResponseBody{
		Message: res.Message,
		Name:    res.Name,
	}
	return body
}

// NewUpdateNotFoundResponseBody builds the HTTP response body from the result
// of the "update" endpoint of the "character" service.
func NewUpdateNotFoundResponseBody(res *character.NotFound) *UpdateNotFoundResponseBody {
	body := &UpdateNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewRemoveNotFoundResponseBody builds the HTTP response body from the result
// of the "remove" endpoint of the "character" service.
func NewRemoveNotFoundResponseBody(res *character.NotFound) *RemoveNotFoundResponseBody {
	body := &RemoveNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewShowPayload builds a character service show endpoint payload.
func NewShowPayload(id string, view *string) *character.ShowPayload {
	v := &character.ShowPayload{}
	v.ID = id
	v.View = view

	return v
}

// NewAddCharacter builds a character service add endpoint payload.
func NewAddCharacter(body *AddRequestBody) *character.Character {
	v := &character.Character{
		Name:        *body.Name,
		Description: body.Description,
		Health:      *body.Health,
		Experience:  body.Experience,
	}

	return v
}

// NewUpdatePayload builds a character service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody, id string) *character.UpdatePayload {
	v := &character.UpdatePayload{}
	v.Character = unmarshalCharacterRequestBodyToCharacterCharacter(body.Character)
	v.ID = id

	return v
}

// NewRemovePayload builds a character service remove endpoint payload.
func NewRemovePayload(id string) *character.RemovePayload {
	v := &character.RemovePayload{}
	v.ID = id

	return v
}

// ValidateAddRequestBody runs the validations defined on AddRequestBody
func ValidateAddRequestBody(body *AddRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Health == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("health", "body"))
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
	if body.Health != nil {
		if *body.Health < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.health", *body.Health, 0, true))
		}
	}
	if body.Health != nil {
		if *body.Health > 2000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.health", *body.Health, 2000, false))
		}
	}
	if body.Experience != nil {
		if *body.Experience < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.experience", *body.Experience, 0, true))
		}
	}
	if body.Experience != nil {
		if *body.Experience > 100000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.experience", *body.Experience, 100000, false))
		}
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Character == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("character", "body"))
	}
	if body.Character != nil {
		if err2 := ValidateCharacterRequestBody(body.Character); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateCharacterRequestBody runs the validations defined on
// CharacterRequestBody
func ValidateCharacterRequestBody(body *CharacterRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Health == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("health", "body"))
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
	if body.Health != nil {
		if *body.Health < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.health", *body.Health, 0, true))
		}
	}
	if body.Health != nil {
		if *body.Health > 2000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.health", *body.Health, 2000, false))
		}
	}
	if body.Experience != nil {
		if *body.Experience < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.experience", *body.Experience, 0, true))
		}
	}
	if body.Experience != nil {
		if *body.Experience > 100000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.experience", *body.Experience, 100000, false))
		}
	}
	return
}