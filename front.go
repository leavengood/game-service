package game

import (
	"context"
	"errors"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	character "game-service/gen/character"
	front "game-service/gen/front"
	characterc "game-service/gen/grpc/character/client"
)

// front service example implementation.
// The example methods log the requests and return zero values.
type frontsrvc struct {
	logger     *log.Logger
	grpcClient *grpc.ClientConn
}

// NewFront returns the front service implementation.
func NewFront(logger *log.Logger, endpoint string) front.Service {
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// We could return an error but the main that calls this just would panic anyhow
	if err != nil {
		panic(err)
	}
	return &frontsrvc{
		logger:     logger,
		grpcClient: conn,
	}
}

// List all characters
func (s *frontsrvc) ListCharacters(ctx context.Context) (front.StoredCharacterCollection, error) {
	s.logger.Print("front.list-characters")

	// TODO: cache these clients in the struct
	c := characterc.NewClient(s.grpcClient)
	endpoint := c.List()

	anyChars, err := endpoint(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	chars := anyChars.(character.StoredCharacterCollection)
	result := make(front.StoredCharacterCollection, 0, len(chars))
	for _, c := range chars {
		result = append(result, toFrontChar(c))
	}

	return result, nil
}

func toFrontChar(c *character.StoredCharacter) *front.StoredCharacter {
	return &front.StoredCharacter{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Health:      c.Health,
		Experience:  c.Experience,
	}
}

// Show character by ID
func (s *frontsrvc) ShowCharacter(ctx context.Context, p *front.ShowCharacterPayload) (*front.StoredCharacter, string, error) {
	s.logger.Print("front.show-character")

	var view string
	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}

	c := characterc.NewClient(s.grpcClient)
	endpoint := c.Show()
	payload := &character.ShowPayload{
		ID:   p.ID,
		View: &view,
	}

	anyChar, err := endpoint(context.Background(), payload)
	if err != nil {
		// The special errors like NotFound have to be converted to the
		// type in each package otherwise the server has a panic
		var nf *character.NotFound
		if errors.As(err, &nf) {
			return nil, "", &front.NotFound{
				ID:      nf.ID,
				Message: nf.Message,
			}
		}
		return nil, "", err
	}

	// This any based interface leaves a lot to be desired...
	char := anyChar.(*character.StoredCharacter)
	result := toFrontChar(char)

	return result, view, nil
}

// Add new character and return its ID
func (s *frontsrvc) AddCharacter(ctx context.Context, p *front.Character) (res string, err error) {
	s.logger.Print("front.add-character")
	return
}

// Update a character with the given ID
func (s *frontsrvc) UpdateCharacter(ctx context.Context, p *front.UpdateCharacterPayload) (err error) {
	s.logger.Print("front.update-character")
	return
}

// Remove a character
func (s *frontsrvc) RemoveCharacter(ctx context.Context, p *front.RemoveCharacterPayload) (err error) {
	s.logger.Print("front.remove-character")
	return
}

// Add an item to a character's inventory and return its ID
func (s *frontsrvc) AddItem(ctx context.Context, p *front.AddItemPayload) (res string, err error) {
	s.logger.Print("front.add-item")
	return
}

// Remove an item ID from a character's inventory
func (s *frontsrvc) RemoveItem(ctx context.Context, p *front.RemoveItemPayload) (err error) {
	s.logger.Print("front.remove-item")
	return
}
