// Code generated by goa v3.12.4, DO NOT EDIT.
//
// front service
//
// Command:
// $ goa gen game-service/design

package front

import (
	"context"
	frontviews "game-service/gen/front/views"
)

// The front service is the main HTTP API for the game
type Service interface {
	// List all items
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	ListItems(context.Context) (res StoredItemCollection, view string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "front"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"list-items"}

// A StoredItem describes an item stored with an ID
type StoredItem struct {
	// ID is the unique id of the item.
	ID string
	// Name of item
	Name string
	// Description of item
	Description *string
	// Amount of damage the item can do
	Damage uint32
	// Amount of healing the item can provide
	Healing uint32
	// Amount of protection the item can provide
	Protection uint32
}

// StoredItemCollection is the result type of the front service list-items
// method.
type StoredItemCollection []*StoredItem

// NewStoredItemCollection initializes result type StoredItemCollection from
// viewed result type StoredItemCollection.
func NewStoredItemCollection(vres frontviews.StoredItemCollection) StoredItemCollection {
	var res StoredItemCollection
	switch vres.View {
	case "default", "":
		res = newStoredItemCollection(vres.Projected)
	case "tiny":
		res = newStoredItemCollectionTiny(vres.Projected)
	}
	return res
}

// NewViewedStoredItemCollection initializes viewed result type
// StoredItemCollection from result type StoredItemCollection using the given
// view.
func NewViewedStoredItemCollection(res StoredItemCollection, view string) frontviews.StoredItemCollection {
	var vres frontviews.StoredItemCollection
	switch view {
	case "default", "":
		p := newStoredItemCollectionView(res)
		vres = frontviews.StoredItemCollection{Projected: p, View: "default"}
	case "tiny":
		p := newStoredItemCollectionViewTiny(res)
		vres = frontviews.StoredItemCollection{Projected: p, View: "tiny"}
	}
	return vres
}

// newStoredItemCollection converts projected type StoredItemCollection to
// service type StoredItemCollection.
func newStoredItemCollection(vres frontviews.StoredItemCollectionView) StoredItemCollection {
	res := make(StoredItemCollection, len(vres))
	for i, n := range vres {
		res[i] = newStoredItem(n)
	}
	return res
}

// newStoredItemCollectionTiny converts projected type StoredItemCollection to
// service type StoredItemCollection.
func newStoredItemCollectionTiny(vres frontviews.StoredItemCollectionView) StoredItemCollection {
	res := make(StoredItemCollection, len(vres))
	for i, n := range vres {
		res[i] = newStoredItemTiny(n)
	}
	return res
}

// newStoredItemCollectionView projects result type StoredItemCollection to
// projected type StoredItemCollectionView using the "default" view.
func newStoredItemCollectionView(res StoredItemCollection) frontviews.StoredItemCollectionView {
	vres := make(frontviews.StoredItemCollectionView, len(res))
	for i, n := range res {
		vres[i] = newStoredItemView(n)
	}
	return vres
}

// newStoredItemCollectionViewTiny projects result type StoredItemCollection to
// projected type StoredItemCollectionView using the "tiny" view.
func newStoredItemCollectionViewTiny(res StoredItemCollection) frontviews.StoredItemCollectionView {
	vres := make(frontviews.StoredItemCollectionView, len(res))
	for i, n := range res {
		vres[i] = newStoredItemViewTiny(n)
	}
	return vres
}

// newStoredItem converts projected type StoredItem to service type StoredItem.
func newStoredItem(vres *frontviews.StoredItemView) *StoredItem {
	res := &StoredItem{
		Description: vres.Description,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Damage != nil {
		res.Damage = *vres.Damage
	}
	if vres.Healing != nil {
		res.Healing = *vres.Healing
	}
	if vres.Protection != nil {
		res.Protection = *vres.Protection
	}
	return res
}

// newStoredItemTiny converts projected type StoredItem to service type
// StoredItem.
func newStoredItemTiny(vres *frontviews.StoredItemView) *StoredItem {
	res := &StoredItem{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	return res
}

// newStoredItemView projects result type StoredItem to projected type
// StoredItemView using the "default" view.
func newStoredItemView(res *StoredItem) *frontviews.StoredItemView {
	vres := &frontviews.StoredItemView{
		ID:          &res.ID,
		Name:        &res.Name,
		Description: res.Description,
		Damage:      &res.Damage,
		Healing:     &res.Healing,
		Protection:  &res.Protection,
	}
	return vres
}

// newStoredItemViewTiny projects result type StoredItem to projected type
// StoredItemView using the "tiny" view.
func newStoredItemViewTiny(res *StoredItem) *frontviews.StoredItemView {
	vres := &frontviews.StoredItemView{
		ID:   &res.ID,
		Name: &res.Name,
	}
	return vres
}