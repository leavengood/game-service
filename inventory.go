package game

import (
	"context"
	inventory "game-service/gen/inventory"
	"log"
)

type set map[string]struct{}

// inventorysrvc manages a set of mappings between characters and their items
type inventorysrvc struct {
	logger *log.Logger
	// Map of character ID to a set of item IDs
	inventory map[string]set
}

// NewInventory returns the inventory service implementation.
func NewInventory(logger *log.Logger) inventory.Service {
	return &inventorysrvc{
		logger:    logger,
		inventory: map[string]set{},
	}
}

// Show the inventory for a character as a list of item IDs
func (s *inventorysrvc) Show(ctx context.Context, p *inventory.ShowPayload) (res []string, err error) {
	s.logger.Print("inventory.show")

	itemSet := s.inventory[p.ID]
	if itemSet == nil {
		return nil, &inventory.NotFound{
			ID:      p.ID,
			Message: "no character inventory found for that ID",
		}
	}

	for id := range itemSet {
		res = append(res, id)
	}

	return
}

// Add an item ID to a character's inventory
func (s *inventorysrvc) Add(ctx context.Context, p *inventory.AddPayload) error {
	s.logger.Print("inventory.add")

	if _, found := s.inventory[p.ID]; !found {
		s.inventory[p.ID] = set{}
	}
	s.inventory[p.ID][p.ItemID] = struct{}{}

	return nil
}

// Remove an item ID from a character's inventory
func (s *inventorysrvc) Remove(ctx context.Context, p *inventory.RemovePayload) error {
	s.logger.Print("inventory.remove")

	itemSet := s.inventory[p.ID]
	if itemSet == nil {
		return &inventory.NotFound{
			ID:      p.ID,
			Message: "no character inventory found for that ID",
		}
	}
	// This is idempotent and will have no affect if that ID does not exist in
	// the inventory.
	delete(itemSet, p.ItemID)

	return nil
}
