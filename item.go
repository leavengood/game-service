package game

import (
	"context"
	"log"

	"github.com/google/uuid"

	item "game-service/gen/item"
	itemview "game-service/gen/item/views"
)

// itemsrvc manages an in-memory list of items
type itemsrvc struct {
	logger *log.Logger
	items  item.StoredItemCollection
}

// NewItem returns the item service implementation.
func NewItem(logger *log.Logger) item.Service {
	return &itemsrvc{
		logger: logger,
		items: item.StoredItemCollection{
			{
				ID:         uuid.NewString(),
				Name:       "Shield of Tarthus",
				Damage:     0,
				Healing:    0,
				Protection: 10,
			},
			{
				ID:         uuid.NewString(),
				Name:       "Golden Ring of Healing",
				Damage:     0,
				Healing:    100,
				Protection: 5,
			},
			{
				ID:         uuid.NewString(),
				Name:       "Heavy Axe",
				Damage:     60,
				Healing:    0,
				Protection: 5,
			},
		},
	}
}

// List all items
func (s *itemsrvc) List(ctx context.Context) (item.StoredItemCollection, error) {
	s.logger.Print("item.list")

	return s.items, nil
}

// Show item by ID
// The "view" return value must have one of the following views
//   - "default"
//   - "tiny"
func (s *itemsrvc) Show(ctx context.Context, p *item.ShowPayload) (res *item.StoredItem, view string, err error) {
	s.logger.Print("item.show")

	if p.View != nil {
		view = *p.View
	} else {
		view = "default"
	}

	for _, item := range s.items {
		if item.ID == p.ID {
			res = item
			break
		}
	}

	if res == nil {
		return res, view, itemNotFound(p.ID)
	}

	return res, view, nil
}

// Add new item and return its ID
func (s *itemsrvc) Add(ctx context.Context, it *item.Item) (string, error) {
	s.logger.Print("item.add")

	if err := s.checkNameUniqueness(it.Name); err != nil {
		return "", err
	}

	id := uuid.NewString()
	stored := &item.StoredItem{
		ID:          id,
		Name:        it.Name,
		Description: it.Description,
		Damage:      it.Damage,
		Healing:     it.Healing,
		Protection:  it.Protection,
	}

	if err := itemview.ValidateStoredItemView(itemToView(stored)); err != nil {
		return "", err
	}

	s.items = append(s.items, stored)

	return id, nil
}

func (s *itemsrvc) Update(ctx context.Context, u *item.UpdatePayload) error {
	var existing *item.StoredItem
	for _, item := range s.items {
		if item.ID == u.ID {
			s.logger.Printf("found an item with ID %s", u.ID)
			existing = item
			break
		}
	}
	if existing == nil {
		return itemNotFound(u.ID)
	}

	// Perform the update
	newItem := u.Item
	if existing.Name != "" {
		if err := s.checkNameUniqueness(u.Item.Name); err != nil {
			return err
		}
		existing.Name = newItem.Name
	}
	if newItem.Description != nil {
		existing.Description = newItem.Description
	}
	existing.Damage = newItem.Damage
	existing.Healing = newItem.Healing
	existing.Protection = newItem.Protection

	return nil
}

func itemToView(it *item.StoredItem) *itemview.StoredItemView {
	return &itemview.StoredItemView{
		ID:          stringToPointer(it.ID),
		Name:        stringToPointer(it.Name),
		Description: it.Description,
		Damage:      &it.Damage,
		Healing:     &it.Healing,
		Protection:  &it.Protection,
	}
}

func stringToPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// Remove an item
func (s *itemsrvc) Remove(ctx context.Context, p *item.RemovePayload) error {
	s.logger.Printf("item.remove ID %s", p.ID)

	newItems := make(item.StoredItemCollection, 0, len(s.items))
	found := false
	for _, item := range s.items {
		if item.ID == p.ID {
			s.logger.Printf("found an item with ID %s", p.ID)
			found = true
		} else {
			newItems = append(newItems, item)
		}
	}

	if !found {
		return itemNotFound(p.ID)
	}

	s.items = newItems

	return nil
}

func (s *itemsrvc) checkNameUniqueness(name string) error {
	// This is not very efficient but in real code we would have a database with indexes
	for _, i := range s.items {
		if i.Name == name {
			return &item.NameTaken{
				Message: "an item with that name already exists",
				Name:    name,
			}
		}
	}
	return nil
}

func itemNotFound(id string) error {
	return &item.NotFound{
		Message: "could not find an item with this ID",
		ID:      id,
	}
}
