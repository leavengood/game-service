package game

import (
	"context"
	"log"

	"github.com/google/uuid"

	character "game-service/gen/character"
)

// character service example implementation.
// The example methods log the requests and return zero values.
type charactersrvc struct {
	logger     *log.Logger
	characters map[string]*character.StoredCharacter
}

// NewCharacter returns the character service implementation.
func NewCharacter(logger *log.Logger) character.Service {
	startingCharacters := []*character.StoredCharacter{
		{
			ID:          uuid.NewString(),
			Name:        "Gordax the Orc",
			Description: stringPtr("A grumpy old orc with a missing eye and golden tusks"),
			Health:      1534,
			Experience:  43981,
		},
		{
			ID:          uuid.NewString(),
			Name:        "Silva De Mort",
			Description: stringPtr("A human thief adept at camouflage and stealth"),
			Health:      671,
			Experience:  17398,
		},
	}

	return &charactersrvc{
		logger: logger,
		characters: map[string]*character.StoredCharacter{
			startingCharacters[0].ID: startingCharacters[0],
			startingCharacters[1].ID: startingCharacters[1],
		},
	}
}

func stringPtr(s string) *string {
	return &s
}

// List all characters
func (s *charactersrvc) List(ctx context.Context) (res character.StoredCharacterCollection, err error) {
	s.logger.Print("character.list")

	for _, c := range s.characters {
		res = append(res, c)
	}

	return
}

// Show character by ID
func (s *charactersrvc) Show(ctx context.Context, p *character.ShowPayload) (res *character.StoredCharacter, view string, err error) {
	s.logger.Print("character.show")

	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}

	res = s.characters[p.ID]

	if res == nil {
		return res, view, characterNotFound(p.ID)
	}

	return
}

// Add new character and return its ID
func (s *charactersrvc) Add(ctx context.Context, p *character.Character) (string, error) {
	s.logger.Print("character.add")

	if err := s.checkNameUniqueness(p.Name); err != nil {
		return "", err
	}

	// TODO: Validation of fields
	id := uuid.NewString()
	xp := uint32(0)
	if p.Experience != nil {
		xp = *p.Experience
	}
	stored := &character.StoredCharacter{
		ID:          id,
		Name:        p.Name,
		Description: p.Description,
		Health:      p.Health,
		Experience:  xp,
	}
	s.characters[stored.ID] = stored

	return id, nil
}

// Update a character with the given ID
func (s *charactersrvc) Update(ctx context.Context, p *character.UpdatePayload) error {
	s.logger.Print("character.update")

	existing := s.characters[p.ID]
	if existing == nil {
		return characterNotFound(p.ID)
	}

	// Perform the update
	if p.Character.Name != "" {
		if err := s.checkNameUniqueness(p.Character.Name); err != nil {
			return err
		}

		existing.Name = p.Character.Name
	}
	if p.Character.Description != nil && *p.Character.Description != "" {
		existing.Description = p.Character.Description
	}
	if p.Character.Health != 0 {
		existing.Health = p.Character.Health
	}
	if p.Character.Experience != nil && *p.Character.Experience != 0 {
		existing.Experience = *p.Character.Experience
	}

	return nil
}

// Remove a character
func (s *charactersrvc) Remove(ctx context.Context, p *character.RemovePayload) error {
	s.logger.Print("character.remove")

	if _, found := s.characters[p.ID]; !found {
		return characterNotFound(p.ID)
	}

	delete(s.characters, p.ID)

	return nil
}

func (s *charactersrvc) checkNameUniqueness(name string) error {
	// This is not very efficient but in real code we would have a database with indexes
	for _, c := range s.characters {
		if c.Name == name {
			return &character.NameTaken{
				Message: "a character with that name already exists",
				Name:    name,
			}
		}
	}
	return nil
}

func characterNotFound(id string) error {
	return &character.NotFound{
		Message: "could not find a character with this ID",
		ID:      id,
	}
}
