// Code generated by goa v3.12.4, DO NOT EDIT.
//
// character views
//
// Command:
// $ goa gen game-service/design

package views

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// StoredCharacterCollection is the viewed result type that is projected based
// on a view.
type StoredCharacterCollection struct {
	// Type to project
	Projected StoredCharacterCollectionView
	// View to render
	View string
}

// StoredCharacter is the viewed result type that is projected based on a view.
type StoredCharacter struct {
	// Type to project
	Projected *StoredCharacterView
	// View to render
	View string
}

// StoredCharacterCollectionView is a type that runs validations on a projected
// type.
type StoredCharacterCollectionView []*StoredCharacterView

// StoredCharacterView is a type that runs validations on a projected type.
type StoredCharacterView struct {
	// ID is the unique id of the character.
	ID *string
	// Name of the character
	Name *string
	// Description of the character
	Description *string
	// Amount of health the character has
	Health *uint32
	// Amount of experience the character has
	Experience *uint32
}

var (
	// StoredCharacterCollectionMap is a map indexing the attribute names of
	// StoredCharacterCollection by view name.
	StoredCharacterCollectionMap = map[string][]string{
		"default": {
			"id",
			"name",
			"description",
			"health",
			"experience",
		},
		"tiny": {
			"id",
			"name",
		},
	}
	// StoredCharacterMap is a map indexing the attribute names of StoredCharacter
	// by view name.
	StoredCharacterMap = map[string][]string{
		"default": {
			"id",
			"name",
			"description",
			"health",
			"experience",
		},
		"tiny": {
			"id",
			"name",
		},
	}
)

// ValidateStoredCharacterCollection runs the validations defined on the viewed
// result type StoredCharacterCollection.
func ValidateStoredCharacterCollection(result StoredCharacterCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStoredCharacterCollectionView(result.Projected)
	case "tiny":
		err = ValidateStoredCharacterCollectionViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default", "tiny"})
	}
	return
}

// ValidateStoredCharacter runs the validations defined on the viewed result
// type StoredCharacter.
func ValidateStoredCharacter(result *StoredCharacter) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStoredCharacterView(result.Projected)
	case "tiny":
		err = ValidateStoredCharacterViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default", "tiny"})
	}
	return
}

// ValidateStoredCharacterCollectionView runs the validations defined on
// StoredCharacterCollectionView using the "default" view.
func ValidateStoredCharacterCollectionView(result StoredCharacterCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateStoredCharacterView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStoredCharacterCollectionViewTiny runs the validations defined on
// StoredCharacterCollectionView using the "tiny" view.
func ValidateStoredCharacterCollectionViewTiny(result StoredCharacterCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateStoredCharacterViewTiny(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStoredCharacterView runs the validations defined on
// StoredCharacterView using the "default" view.
func ValidateStoredCharacterView(result *StoredCharacterView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Health == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("health", "result"))
	}
	if result.Experience == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("experience", "result"))
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
	if result.Health != nil {
		if *result.Health < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.health", *result.Health, 0, true))
		}
	}
	if result.Health != nil {
		if *result.Health > 2000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.health", *result.Health, 2000, false))
		}
	}
	if result.Experience != nil {
		if *result.Experience < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.experience", *result.Experience, 0, true))
		}
	}
	if result.Experience != nil {
		if *result.Experience > 100000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.experience", *result.Experience, 100000, false))
		}
	}
	return
}

// ValidateStoredCharacterViewTiny runs the validations defined on
// StoredCharacterView using the "tiny" view.
func ValidateStoredCharacterViewTiny(result *StoredCharacterView) (err error) {
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
