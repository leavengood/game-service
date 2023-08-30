package game

import (
	"context"
	front "game-service/gen/front"
	item "game-service/gen/item"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	itemc "game-service/gen/grpc/item/client"
)

// front service example implementation.
// The example methods log the requests and return zero values.
type frontsrvc struct {
	logger     *log.Logger
	grpcClient *grpc.ClientConn
}

// NewFront returns the front service implementation.
func NewFront(logger *log.Logger) front.Service {
	// TODO: Parameterize this URL
	// TODO: Return an error!
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return &frontsrvc{
		logger:     logger,
		grpcClient: conn,
	}
}

// List all items
func (s *frontsrvc) ListItems(ctx context.Context) (front.StoredItemCollection, string, error) {
	s.logger.Print("front.list")

	c := itemc.NewClient(s.grpcClient)
	endpoint := c.List()
	anyItems, err := endpoint(context.Background(), nil)
	if err != nil {
		return nil, "", err
	}
	items := anyItems.(item.StoredItemCollection)
	result := make(front.StoredItemCollection, 0, len(items))
	for _, it := range items {
		s.logger.Printf("item: %#v\n", it)
		result = append(result, &front.StoredItem{
			ID:          it.ID,
			Name:        it.Name,
			Description: it.Description,
			Damage:      it.Damage,
			Healing:     it.Healing,
			Protection:  it.Protection,
		})
	}

	return result, "default", nil
}
