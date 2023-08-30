package design

import (
	. "goa.design/goa/v3/dsl"
)

// This service only uses IDs and counts on callers to use the proper IDs
var _ = Service("inventory", func() {
	Description("The inventory service is the service for managing character inventories")

	HTTP(func() {
		Path("/back/inventory")
	})

	Method("show", func() {
		Description("Show the inventory for a character as a list of item IDs")
		Payload(func() {
			Field(1, "id", String, "ID of the character")
			Required("id")
		})
		Result(ArrayOf(String))
		Error("not_found", NotFound, "Character not found")
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
		GRPC(func() {
			Response(CodeOK)
			Response("not_found", CodeNotFound)
		})
	})

	Method("add", func() {
		Description("Add an item ID to a character's inventory")
		Payload(func() {
			Field(1, "id", String, "ID of the character")
			Field(2, "item_id", String, "ID of the item")
			Required("id", "item_id")
		})
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("remove", func() {
		Description("Remove an item ID from a character's inventory")
		Payload(func() {
			Field(1, "id", String, "ID of the character")
			Field(2, "item_id", String, "ID of the item")
			Required("id", "item_id")
		})
		Error("not_found", NotFound, "Character or item not found")
		HTTP(func() {
			DELETE("/")
			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})
})
