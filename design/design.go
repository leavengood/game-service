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

var _ = Service("item", func() {
	Description("The item service is the service for managing items")

	HTTP(func() {
		Path("/back/items")
	})

	Method("list", func() {
		Description("List all items")
		Result(CollectionOf(StoredItem), func() {
			View("default")
		})
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("show", func() {
		Description("Show item by ID")
		Payload(func() {
			Field(1, "id", String, "ID of item to show")
			Field(2, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
			Required("id")
		})
		Result(StoredItem)
		Error("not_found", NotFound, "Item not found")
		HTTP(func() {
			GET("/{id}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
		GRPC(func() {
			Metadata(func() {
				Attribute("view")
			})
			Response(CodeOK)
			Response("not_found", CodeNotFound)
		})
	})

	Method("add", func() {
		Description("Add new item and return its ID")
		Payload(Item)
		Result(String)
		Error("name_taken", NameTaken, "name is taken")
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
			Response("name_taken", StatusConflict)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("update", func() {
		Description("Update an item with the given ID")
		Payload(func() {
			Field(1, "id", String, "ID of the item to be updated")
			Field(2, "item", Item, "Item with updated fields")
			Required("id", "item")
		})
		Error("not_found", NotFound, "Item not found")
		Error("name_taken", NameTaken, "name is taken")
		HTTP(func() {
			PUT("/{id}")
			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
			Response("name_taken", StatusConflict)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("remove", func() {
		Description("Remove an item")
		Payload(func() {
			Field(1, "id", String, "ID of item to remove")
			Required("id")
		})
		Error("not_found", NotFound, "Item not found")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})
})

var _ = Service("character", func() {
	Description("The character service is the service for managing characters")

	HTTP(func() {
		Path("/back/characters")
	})

	Method("list", func() {
		Description("List all characters")
		Result(CollectionOf(StoredCharacter), func() {
			View("default")
		})
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("show", func() {
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
			GET("/{id}")
			Param("view")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
		GRPC(func() {
			Metadata(func() {
				Attribute("view")
			})
			Response(CodeOK)
			Response("not_found", CodeNotFound)
		})
	})

	Method("add", func() {
		Description("Add new character and return its ID")
		Payload(Character)
		Result(String)
		Error("name_taken", NameTaken, "name is taken")
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
			Response("name_taken", StatusConflict)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("update", func() {
		Description("Update a character with the given ID")
		Payload(func() {
			Field(1, "id", String, "ID of the character to be updated")
			Field(2, "character", Character, "character with updated fields")
			Required("id", "character")
		})
		Error("not_found", NotFound, "character not found")
		Error("name_taken", NameTaken, "name is taken")
		HTTP(func() {
			PUT("/{id}")
			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
			Response("name_taken", StatusConflict)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("remove", func() {
		Description("Remove a character")
		Payload(func() {
			Field(1, "id", String, "ID of character to remove")
			Required("id")
		})
		Error("not_found", NotFound, "character not found")
		HTTP(func() {
			DELETE("/{id}")
			Response(StatusNoContent)
			Response("not_found", StatusNotFound)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})
})

var StoredItem = ResultType("application/vnd.game.stored-item", func() {
	Description("A StoredItem describes an item stored with an ID")
	Reference(Item)
	TypeName("StoredItem")

	Attributes(func() {
		Attribute("id", String, "ID is the unique id of the item.", func() {
			Example("123abc")
			Meta("rpc:tag", "7")
		})
		Field(2, "name")
		Field(3, "description")
		Field(4, "damage")
		Field(5, "healing")
		Field(6, "protection")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("description")
		Attribute("damage")
		Attribute("healing")
		Attribute("protection")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("name")
	})

	Required("id", "name", "damage", "healing", "protection")
})

var Item = Type("Item", func() {
	Description("Item describes an item in a player's inventory")
	// TODO: Figure out how to ensure the name is unique
	Attribute("name", String, "Name of item", func() {
		MaxLength(100)
		Example("Sword of Wisdom")
		Meta("rpc:tag", "1")
	})
	Attribute("description", String, "Description of item", func() {
		MaxLength(2000)
		Example("A magnificent sword which grants the bearer +2 wisdom")
		Meta("rpc:tag", "2")
	})
	Attribute("damage", UInt32, "Amount of damage the item can do", func() {
		Minimum(0)
		Maximum(200)
		Meta("rpc:tag", "3")
	})
	Attribute("healing", UInt32, "Amount of healing the item can provide", func() {
		Minimum(0)
		Maximum(200)
		Meta("rpc:tag", "4")
	})
	Attribute("protection", UInt32, "Amount of protection the item can provide", func() {
		Minimum(0)
		Maximum(20)
		Meta("rpc:tag", "5")
	})
	Required("name", "damage", "healing", "protection")
})

var StoredCharacter = ResultType("application/vnd.game.stored-character", func() {
	Description("A StoredCharacter describes an character stored with an ID")
	Reference(Character)
	TypeName("StoredCharacter")

	Attributes(func() {
		Attribute("id", String, "ID is the unique id of the character.", func() {
			Example("123abc")
			Meta("rpc:tag", "6")
		})
		Field(2, "name")
		Field(3, "description")
		Field(4, "health")
		Field(5, "experience")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("description")
		Attribute("health")
		Attribute("experience")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("name")
	})

	Required("id", "name", "health", "experience")
})

var Character = Type("Character", func() {
	Description("Character describes a character in the game")
	// TODO: Figure out how to ensure the name is unique
	Attribute("name", String, "Name of the character", func() {
		MaxLength(100)
		Example("Arvish the Wise")
		Meta("rpc:tag", "1")
	})
	Attribute("description", String, "Description of the character", func() {
		MaxLength(2000)
		Example("A grizzled wizard with a penchant for mayhem and mead")
		Meta("rpc:tag", "2")
	})
	Attribute("health", UInt32, "Amount of health the character has", func() {
		Minimum(0)
		Maximum(2000)
		Meta("rpc:tag", "3")
	})
	Attribute("experience", UInt32, "Amount of experience the character has", func() {
		Minimum(0)
		Maximum(100_000)
		Meta("rpc:tag", "4")
	})
	Required("name", "health")
})

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
