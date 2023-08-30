package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("game", func() {
	Title("Game Service")
	Description("Service for a multiplayer game")
	Server("game", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

var _ = Service("front", func() {
	Description("The front service is the main HTTP API for the game")

	Method("list-items", func() {
		Description("List all items")
		Result(CollectionOf(StoredItem))
		HTTP(func() {
			GET("/items")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})
})

// Inventory is a nicer type used in the front API to map the plain IDs in the
// inventory backing API to nicer objects
var Inventory = Type("Inventory", func() {
	Field(1, "character", StoredCharacter, "Character that owns this inventory")
	Field(2, "items", CollectionOf(StoredItem), "Items is the list of item's in this inventory")
})

var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when attempting to show, update or delete an API object that does not exist")
	Field(1, "message", String, "not found")
	Field(2, "id", String, "ID of missing item")
	Required("message", "id")
})

var NameTaken = Type("NameTaken", func() {
	Description("NameTaken is the type returned when a name for an API object already exists")
	Field(1, "message", String, "name taken")
	Field(2, "name", String, "name that is not unique")
	Required("message", "name")
})
