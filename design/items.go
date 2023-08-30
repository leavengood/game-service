package design

import (
	. "goa.design/goa/v3/dsl"
)

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
			Response("name_taken", CodeAlreadyExists)
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
			Response("not_found", CodeNotFound)
			Response("name_taken", CodeAlreadyExists)
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
			Response("not_found", CodeNotFound)
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
