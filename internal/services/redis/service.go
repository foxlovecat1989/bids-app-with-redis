package redis

import (
	"context"
	"time"
)

// Service provides application-specific Redis functionality
type Service struct {
	client *Client
}

// NewService creates a new Redis service
func NewService(host string, port int, db int, password string) (*Service, error) {
	client, err := NewClient(Config{
		Host:     host,
		Port:     port,
		DB:       db,
		Password: password,
	})

	if err != nil {
		return nil, err
	}

	return &Service{
		client: client,
	}, nil
}

// Ping checks the Redis connection
func (s *Service) Ping(ctx context.Context) error {
	_, err := s.client.Get(ctx, "ping-check")
	// Ignoring redis.Nil error, which is expected if the key doesn't exist
	if err != nil && err.Error() != "redis: nil" {
		return err
	}
	return nil
}

// StoreBid stores a bid in Redis
func (s *Service) StoreBid(ctx context.Context, bidID, userID, itemID string, amount float64) error {
	key := "bid:" + bidID

	// Use a hash to store the bid data
	err := s.client.HSet(ctx, key,
		"user_id", userID,
		"item_id", itemID,
		"amount", amount,
		"timestamp", time.Now().Unix(),
	)

	if err != nil {
		return err
	}

	// Store in a sorted set for the item to track highest bids
	itemBidsKey := "item:" + itemID + ":bids"
	err = s.client.Set(ctx, itemBidsKey+":"+bidID, bidID, 0)
	if err != nil {
		return err
	}

	// Add to user's bids
	userBidsKey := "user:" + userID + ":bids"
	err = s.client.Set(ctx, userBidsKey+":"+bidID, bidID, 0)
	if err != nil {
		return err
	}

	return nil
}

// GetBid retrieves a bid from Redis
func (s *Service) GetBid(ctx context.Context, bidID string) (map[string]string, error) {
	key := "bid:" + bidID
	return s.client.HGetAll(ctx, key)
}

// Close closes the Redis connection
func (s *Service) Close() error {
	return s.client.Close()
}
