package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"bids-app-with-redis/internal/config"
	"bids-app-with-redis/internal/services/redis"
)

func main() {
	// Setup configuration
	cfg := config.NewConfig()
	fmt.Println("Starting Redis test application...")

	// Initialize Redis service
	redisService, err := redis.NewService(
		cfg.RedisHost,
		cfg.RedisPort,
		cfg.RedisDB,
		cfg.RedisPasswd,
	)
	if err != nil {
		log.Fatalf("Failed to initialize Redis service: %v", err)
	}

	// Test Redis connection
	ctx := context.Background()
	if err := redisService.Ping(ctx); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis successfully")

	// Test storing a bid
	bidID := "bid123"
	userID := "user456"
	itemID := "item789"
	amount := 100.50

	fmt.Println("Storing test bid in Redis...")
	err = redisService.StoreBid(ctx, bidID, userID, itemID, amount)
	if err != nil {
		log.Fatalf("Failed to store bid: %v", err)
	}

	// Test retrieving the bid
	fmt.Println("Retrieving test bid from Redis...")
	bid, err := redisService.GetBid(ctx, bidID)
	if err != nil {
		log.Fatalf("Failed to retrieve bid: %v", err)
	}

	fmt.Println("Bid data retrieved:")
	for k, v := range bid {
		fmt.Printf("  %s: %s\n", k, v)
	}

	// Setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("\nPress Ctrl+C to exit")
	<-quit

	// Close Redis connection
	if err := redisService.Close(); err != nil {
		log.Fatalf("Error closing Redis connection: %v", err)
	}

	fmt.Println("Redis connection closed, application exited properly")
}
