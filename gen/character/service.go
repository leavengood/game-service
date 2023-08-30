// Code generated by goa v3.12.4, DO NOT EDIT.
//
// character service
//
// Command:
// $ goa gen game-service/design

package character

import (
	"context"
	characterviews "game-service/gen/character/views"
)

// The character service is the service for managing characters
type Service interface {
	// List all characters
	List(context.Context) (res StoredCharacterCollection, err error)
	// Show character by ID
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	Show(context.Context, *ShowPayload) (res *StoredCharacter, view string, err error)
	// Add new character and return its ID
	Add(context.Context, *Character) (res string, err error)
	// Update a character with the given ID
	Update(context.Context, *UpdatePayload) (err error)
	// Remove a character
	Remove(context.Context, *RemovePayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "character"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"list", "show", "add", "update", "remove"}

// Character is the payload type of the character service add method.
type Character struct {
	// Name of the character
	Name string
	// Description of the character
	Description *string
	// Amount of health the character has
	Health uint32
	// Amount of experience the character has
	Experience *uint32
}

// NameTaken is the type returned when a name for an API object already exists
type NameTaken struct {
	// name taken
	Message string
	// name that is not unique
	Name string
}

// NotFound is the type returned when attempting to show, update or delete an
// API object that does not exist
type NotFound struct {
	// not found
	Message string
	// ID of missing item
	ID string
}

// RemovePayload is the payload type of the character service remove method.
type RemovePayload struct {
	// ID of character to remove
	ID string
}

// ShowPayload is the payload type of the character service show method.
type ShowPayload struct {
	// ID of character to show
	ID string
	// View to render
	View *string
}

// StoredCharacter is the result type of the character service show method.
type StoredCharacter struct {
	// ID is the unique id of the character.
	ID string
	// Name of the character
	Name string
	// Description of the character
	Description *string
	// Amount of health the character has
	Health uint32
	// Amount of experience the character has
	Experience uint32
}

// StoredCharacterCollection is the result type of the character service list
// method.
type StoredCharacterCollection []*StoredCharacter

// UpdatePayload is the payload type of the character service update method.
type UpdatePayload struct {
	// ID of the character to be updated
	ID string
	// character with updated fields
	Character *Character
}

// Error returns an error description.
func (e *NameTaken) Error() string {
	return "NameTaken is the type returned when a name for an API object already exists"
}

// ErrorName returns "NameTaken".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *NameTaken) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "NameTaken".
func (e *NameTaken) GoaErrorName() string {
	return "name_taken"
}

// Error returns an error description.
func (e *NotFound) Error() string {
	return "NotFound is the type returned when attempting to show, update or delete an API object that does not exist"
}

// ErrorName returns "NotFound".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *NotFound) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "NotFound".
func (e *NotFound) GoaErrorName() string {
	return "not_found"
}

// NewStoredCharacterCollection initializes result type
// StoredCharacterCollection from viewed result type StoredCharacterCollection.
func NewStoredCharacterCollection(vres characterviews.StoredCharacterCollection) StoredCharacterCollection {
	var res StoredCharacterCollection
	switch vres.View {
	case "default", "":
		res = newStoredCharacterCollection(vres.Projected)
	case "tiny":
		res = newStoredCharacterCollectionTiny(vres.Projected)
	}
	return res
}

// NewViewedStoredCharacterCollection initializes viewed result type
// StoredCharacterCollection from result type StoredCharacterCollection using
// the given view.
func NewViewedStoredCharacterCollection(res StoredCharacterCollection, view string) characterviews.StoredCharacterCollection {
	var vres characterviews.StoredCharacterCollection
	switch view {
	case "default", "":
		p := newStoredCharacterCollectionView(res)
		vres = characterviews.StoredCharacterCollection{Projected: p, View: "default"}
	case "tiny":
		p := newStoredCharacterCollectionViewTiny(res)
		vres = characterviews.StoredCharacterCollection{Projected: p, View: "tiny"}
	}
	return vres
}

// NewStoredCharacter initializes result type StoredCharacter from viewed
// result type StoredCharacter.
func NewStoredCharacter(vres *characterviews.StoredCharacter) *StoredCharacter {
	var res *StoredCharacter
	switch vres.View {
	case "default", "":
		res = newStoredCharacter(vres.Projected)
	case "tiny":
		res = newStoredCharacterTiny(vres.Projected)
	}
	return res
}

// NewViewedStoredCharacter initializes viewed result type StoredCharacter from
// result type StoredCharacter using the given view.
func NewViewedStoredCharacter(res *StoredCharacter, view string) *characterviews.StoredCharacter {
	var vres *characterviews.StoredCharacter
	switch view {
	case "default", "":
		p := newStoredCharacterView(res)
		vres = &characterviews.StoredCharacter{Projected: p, View: "default"}
	case "tiny":
		p := newStoredCharacterViewTiny(res)
		vres = &characterviews.StoredCharacter{Projected: p, View: "tiny"}
	}
	return vres
}

// newStoredCharacterCollection converts projected type
// StoredCharacterCollection to service type StoredCharacterCollection.
func newStoredCharacterCollection(vres characterviews.StoredCharacterCollectionView) StoredCharacterCollection {
	res := make(StoredCharacterCollection, len(vres))
	for i, n := range vres {
		res[i] = newStoredCharacter(n)
	}
	return res
}

// newStoredCharacterCollectionTiny converts projected type
// StoredCharacterCollection to service type StoredCharacterCollection.
func newStoredCharacterCollectionTiny(vres characterviews.StoredCharacterCollectionView) StoredCharacterCollection {
	res := make(StoredCharacterCollection, len(vres))
	for i, n := range vres {
		res[i] = newStoredCharacterTiny(n)
	}
	return res
}

// newStoredCharacterCollectionView projects result type
// StoredCharacterCollection to projected type StoredCharacterCollectionView
// using the "default" view.
func newStoredCharacterCollectionView(res StoredCharacterCollection) characterviews.StoredCharacterCollectionView {
	vres := make(characterviews.StoredCharacterCollectionView, len(res))
	for i, n := range res {
		vres[i] = newStoredCharacterView(n)
	}
	return vres
}

// newStoredCharacterCollectionViewTiny projects result type
// StoredCharacterCollection to projected type StoredCharacterCollectionView
// using the "tiny" view.
func newStoredCharacterCollectionViewTiny(res StoredCharacterCollection) characterviews.StoredCharacterCollectionView {
	vres := make(characterviews.StoredCharacterCollectionView, len(res))
	for i, n := range res {
		vres[i] = newStoredCharacterViewTiny(n)
	}
	return vres
}

// newStoredCharacter converts projected type StoredCharacter to service type
// StoredCharacter.
func newStoredCharacter(vres *characterviews.StoredCharacterView) *StoredCharacter {
	res := &StoredCharacter{
		Description: vres.Description,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Health != nil {
		res.Health = *vres.Health
	}
	if vres.Experience != nil {
		res.Experience = *vres.Experience
	}
	return res
}

// newStoredCharacterTiny converts projected type StoredCharacter to service
// type StoredCharacter.
func newStoredCharacterTiny(vres *characterviews.StoredCharacterView) *StoredCharacter {
	res := &StoredCharacter{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	return res
}

// newStoredCharacterView projects result type StoredCharacter to projected
// type StoredCharacterView using the "default" view.
func newStoredCharacterView(res *StoredCharacter) *characterviews.StoredCharacterView {
	vres := &characterviews.StoredCharacterView{
		ID:          &res.ID,
		Name:        &res.Name,
		Description: res.Description,
		Health:      &res.Health,
		Experience:  &res.Experience,
	}
	return vres
}

// newStoredCharacterViewTiny projects result type StoredCharacter to projected
// type StoredCharacterView using the "tiny" view.
func newStoredCharacterViewTiny(res *StoredCharacter) *characterviews.StoredCharacterView {
	vres := &characterviews.StoredCharacterView{
		ID:   &res.ID,
		Name: &res.Name,
	}
	return vres
}