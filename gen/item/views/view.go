// Code generated by goa v3.12.4, DO NOT EDIT.
//
// item views
//
// Command:
// $ goa gen game-service/design

package views

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// StoredItemCollection is the viewed result type that is projected based on a
// view.
type StoredItemCollection struct {
	// Type to project
	Projected StoredItemCollectionView
	// View to render
	View string
}

// StoredItem is the viewed result type that is projected based on a view.
type StoredItem struct {
	// Type to project
	Projected *StoredItemView
	// View to render
	View string
}

// StoredItemCollectionView is a type that runs validations on a projected type.
type StoredItemCollectionView []*StoredItemView

// StoredItemView is a type that runs validations on a projected type.
type StoredItemView struct {
	// ID is the unique id of the item.
	ID *string
	// Name of item
	Name *string
	// Description of item
	Description *string
	// Amount of damage the item can do
	Damage *uint32
	// Amount of healing the item can provide
	Healing *uint32
	// Amount of protection the item can provide
	Protection *uint32
}

var (
	// StoredItemCollectionMap is a map indexing the attribute names of
	// StoredItemCollection by view name.
	StoredItemCollectionMap = map[string][]string{
		"default": {
			"id",
			"name",
			"description",
			"damage",
			"healing",
			"protection",
		},
		"tiny": {
			"id",
			"name",
		},
	}
	// StoredItemMap is a map indexing the attribute names of StoredItem by view
	// name.
	StoredItemMap = map[string][]string{
		"default": {
			"id",
			"name",
			"description",
			"damage",
			"healing",
			"protection",
		},
		"tiny": {
			"id",
			"name",
		},
	}
)

// ValidateStoredItemCollection runs the validations defined on the viewed
// result type StoredItemCollection.
func ValidateStoredItemCollection(result StoredItemCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStoredItemCollectionView(result.Projected)
	case "tiny":
		err = ValidateStoredItemCollectionViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default", "tiny"})
	}
	return
}

// ValidateStoredItem runs the validations defined on the viewed result type
// StoredItem.
func ValidateStoredItem(result *StoredItem) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStoredItemView(result.Projected)
	case "tiny":
		err = ValidateStoredItemViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default", "tiny"})
	}
	return
}

// ValidateStoredItemCollectionView runs the validations defined on
// StoredItemCollectionView using the "default" view.
func ValidateStoredItemCollectionView(result StoredItemCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateStoredItemView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStoredItemCollectionViewTiny runs the validations defined on
// StoredItemCollectionView using the "tiny" view.
func ValidateStoredItemCollectionViewTiny(result StoredItemCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateStoredItemViewTiny(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStoredItemView runs the validations defined on StoredItemView using
// the "default" view.
func ValidateStoredItemView(result *StoredItemView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Damage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("damage", "result"))
	}
	if result.Healing == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("healing", "result"))
	}
	if result.Protection == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("protection", "result"))
	}
	if result.Name != nil {
		if utf8.RuneCountInString(*result.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.name", *result.Name, utf8.RuneCountInString(*result.Name), 100, false))
		}
	}
	if result.Description != nil {
		if utf8.RuneCountInString(*result.Description) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.description", *result.Description, utf8.RuneCountInString(*result.Description), 2000, false))
		}
	}
	if result.Damage != nil {
		if *result.Damage < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.damage", *result.Damage, 0, true))
		}
	}
	if result.Damage != nil {
		if *result.Damage > 200 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.damage", *result.Damage, 200, false))
		}
	}
	if result.Healing != nil {
		if *result.Healing < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.healing", *result.Healing, 0, true))
		}
	}
	if result.Healing != nil {
		if *result.Healing > 200 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.healing", *result.Healing, 200, false))
		}
	}
	if result.Protection != nil {
		if *result.Protection < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.protection", *result.Protection, 0, true))
		}
	}
	if result.Protection != nil {
		if *result.Protection > 20 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.protection", *result.Protection, 20, false))
		}
	}
	return
}

// ValidateStoredItemViewTiny runs the validations defined on StoredItemView
// using the "tiny" view.
func ValidateStoredItemViewTiny(result *StoredItemView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Name != nil {
		if utf8.RuneCountInString(*result.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.name", *result.Name, utf8.RuneCountInString(*result.Name), 100, false))
		}
	}
	return
}
