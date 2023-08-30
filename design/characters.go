package design

import (
	. "goa.design/goa/v3/dsl"
)

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
			Response("name_taken", CodeAlreadyExists)
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
			Response("not_found", CodeNotFound)
			Response("name_taken", CodeAlreadyExists)
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
			Response("not_found", CodeNotFound)
		})
	})
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
