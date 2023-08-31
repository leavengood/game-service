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

	Method("list-characters", func() {
		Description("List all characters")
		Result(CollectionOf(StoredCharacter), func() {
			// TODO: Change to default?
			View("tiny")
		})
		HTTP(func() {
			GET("/characters")
			Response(StatusOK)
		})
	})

	Method("show-character", func() {
		Description("Show character by ID")
		Payload(func() {
			Field(1, "id", String, "ID of character to show")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})
		Result(StoredCharacter)
		Error("not_found", NotFound, "character not found")
		HTTP(func() {
			GET("characters/{id}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})

	Method("add-character", func() {
		Description("Add new character and return its ID")
		Payload(Character)
		Result(String)
		Error("name_taken", NameTaken, "name is taken")
		HTTP(func() {
			POST("/characters")
			Response(StatusCreated)
			Response("name_taken", StatusConflict)
		})
	})

	Method("update-character", func() {
		Description("Update a character with the given ID")
		Payload(func() {
			Field(1, "id", String, "ID of the character to be updated")
			Field(2, "character", Character, "character with updated fields")
			Required("id", "character")
		})
		Error("not_found", NotFound, "character not found")
		Error("name_taken", NameTaken, "name is taken")
		HTTP(func() {
			PUT("/characters/{id}")
			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
			Response("name_taken", StatusConflict)
		})
	})

	Method("remove-character", func() {
		Description("Remove a character")
		Payload(func() {
			Field(1, "id", String, "ID of character to remove")
			Required("id")
		})
		Error("not_found", NotFound, "character not found")
		HTTP(func() {
			DELETE("/characters/{id}")
			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
		})
	})

	Method("add-item", func() {
		Description("Add an item to a character's inventory and return its ID")
		Payload(func() {
			Field(1, "id", String, "ID of the character to be updated")
			Field(2, "item", Item, "Item to add")
			Required("id", "item")
		})
		Result(String)
		Error("not_found", NotFound, "character not found")
		Error("name_taken", NameTaken, "name is taken")
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
			Response("name_taken", StatusConflict)
			Response("not_found", StatusNotFound)
		})
	})

	Method("remove-item", func() {
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
	})
})

// // Inventory is a nicer type used in the front API to map the plain IDs in the
// // inventory backing API to nicer objects
// var Inventory = Type("Inventory", func() {
// 	Field(1, "character", StoredCharacter, "Character that owns this inventory")
// 	Field(2, "items", CollectionOf(StoredItem), "Items is the list of item's in this inventory")
// })

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
