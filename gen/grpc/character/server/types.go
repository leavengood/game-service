// Code generated by goa v3.12.4, DO NOT EDIT.
//
// character gRPC server types
//
// Command:
// $ goa gen game-service/design

package server

import (
	character "game-service/gen/character"
	characterviews "game-service/gen/character/views"
	characterpb "game-service/gen/grpc/character/pb"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// NewProtoStoredCharacterCollection builds the gRPC response type from the
// result of the "list" endpoint of the "character" service.
func NewProtoStoredCharacterCollection(result characterviews.StoredCharacterCollectionView) *characterpb.StoredCharacterCollection {
	message := &characterpb.StoredCharacterCollection{}
	message.Field = make([]*characterpb.StoredCharacter, len(result))
	for i, val := range result {
		message.Field[i] = &characterpb.StoredCharacter{
			Id:          *val.ID,
			Name:        *val.Name,
			Description: val.Description,
			Health:      *val.Health,
			Experience:  *val.Experience,
		}
	}
	return message
}

// NewShowPayload builds the payload of the "show" endpoint of the "character"
// service from the gRPC request type.
func NewShowPayload(message *characterpb.ShowRequest, view *string) *character.ShowPayload {
	v := &character.ShowPayload{
		ID: message.Id,
	}
	v.View = view
	return v
}

// NewProtoShowResponse builds the gRPC response type from the result of the
// "show" endpoint of the "character" service.
func NewProtoShowResponse(result *characterviews.StoredCharacterView) *characterpb.ShowResponse {
	message := &characterpb.ShowResponse{
		Id:          *result.ID,
		Name:        *result.Name,
		Description: result.Description,
		Health:      *result.Health,
		Experience:  *result.Experience,
	}
	return message
}

// NewShowNotFoundError builds the gRPC error response type from the error of
// the "show" endpoint of the "character" service.
func NewShowNotFoundError(er *character.NotFound) *characterpb.ShowNotFoundError {
	message := &characterpb.ShowNotFoundError{
		Message_: er.Message,
		Id:       er.ID,
	}
	return message
}

// NewAddPayload builds the payload of the "add" endpoint of the "character"
// service from the gRPC request type.
func NewAddPayload(message *characterpb.AddRequest) *character.Character {
	v := &character.Character{
		Name:        message.Name,
		Description: message.Description,
		Health:      message.Health,
		Experience:  message.Experience,
	}
	return v
}

// NewProtoAddResponse builds the gRPC response type from the result of the
// "add" endpoint of the "character" service.
func NewProtoAddResponse(result string) *characterpb.AddResponse {
	message := &characterpb.AddResponse{}
	message.Field = result
	return message
}

// NewUpdatePayload builds the payload of the "update" endpoint of the
// "character" service from the gRPC request type.
func NewUpdatePayload(message *characterpb.UpdateRequest) *character.UpdatePayload {
	v := &character.UpdatePayload{
		ID: message.Id,
	}
	if message.Character != nil {
		v.Character = protobufCharacterpbCharacter2ToCharacterCharacter(message.Character)
	}
	return v
}

// NewProtoUpdateResponse builds the gRPC response type from the result of the
// "update" endpoint of the "character" service.
func NewProtoUpdateResponse() *characterpb.UpdateResponse {
	message := &characterpb.UpdateResponse{}
	return message
}

// NewRemovePayload builds the payload of the "remove" endpoint of the
// "character" service from the gRPC request type.
func NewRemovePayload(message *characterpb.RemoveRequest) *character.RemovePayload {
	v := &character.RemovePayload{
		ID: message.Id,
	}
	return v
}

// NewProtoRemoveResponse builds the gRPC response type from the result of the
// "remove" endpoint of the "character" service.
func NewProtoRemoveResponse() *characterpb.RemoveResponse {
	message := &characterpb.RemoveResponse{}
	return message
}

// ValidateAddRequest runs the validations defined on AddRequest.
func ValidateAddRequest(message *characterpb.AddRequest) (err error) {
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	if message.Description != nil {
		if utf8.RuneCountInString(*message.Description) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("message.description", *message.Description, utf8.RuneCountInString(*message.Description), 2000, false))
		}
	}
	if message.Health < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.health", message.Health, 0, true))
	}
	if message.Health > 2000 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.health", message.Health, 2000, false))
	}
	if message.Experience != nil {
		if *message.Experience < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.experience", *message.Experience, 0, true))
		}
	}
	if message.Experience != nil {
		if *message.Experience > 100000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.experience", *message.Experience, 100000, false))
		}
	}
	return
}

// ValidateUpdateRequest runs the validations defined on UpdateRequest.
func ValidateUpdateRequest(message *characterpb.UpdateRequest) (err error) {
	if message.Character == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("character", "message"))
	}
	if message.Character != nil {
		if err2 := ValidateCharacter2(message.Character); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateCharacter2 runs the validations defined on Character2.
func ValidateCharacter2(character *characterpb.Character2) (err error) {
	if utf8.RuneCountInString(character.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("character.name", character.Name, utf8.RuneCountInString(character.Name), 100, false))
	}
	if character.Description != nil {
		if utf8.RuneCountInString(*character.Description) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("character.description", *character.Description, utf8.RuneCountInString(*character.Description), 2000, false))
		}
	}
	if character.Health < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("character.health", character.Health, 0, true))
	}
	if character.Health > 2000 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("character.health", character.Health, 2000, false))
	}
	if character.Experience != nil {
		if *character.Experience < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("character.experience", *character.Experience, 0, true))
		}
	}
	if character.Experience != nil {
		if *character.Experience > 100000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("character.experience", *character.Experience, 100000, false))
		}
	}
	return
}

// protobufCharacterpbCharacter2ToCharacterCharacter builds a value of type
// *character.Character from a value of type *characterpb.Character2.
func protobufCharacterpbCharacter2ToCharacterCharacter(v *characterpb.Character2) *character.Character {
	res := &character.Character{
		Name:        v.Name,
		Description: v.Description,
		Health:      v.Health,
		Experience:  v.Experience,
	}

	return res
}

// svcCharacterCharacterToCharacterpbCharacter2 builds a value of type
// *characterpb.Character2 from a value of type *character.Character.
func svcCharacterCharacterToCharacterpbCharacter2(v *character.Character) *characterpb.Character2 {
	res := &characterpb.Character2{
		Name:        v.Name,
		Description: v.Description,
		Health:      v.Health,
		Experience:  v.Experience,
	}

	return res
}
