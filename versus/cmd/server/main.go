package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/chessclub2205-dev/versus-service/internal/api"
	"github.com/chessclub2205-dev/versus-service/internal/db"
	"github.com/chessclub2205-dev/versus-service/internal/match"
	"github.com/chessclub2205-dev/versus-service/internal/payments"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	redisURL := os.Getenv("REDIS_URL")
	if dsn == "" || redisURL == "" {
		log.Fatal("DATABASE_URL and REDIS_URL environment variables required")
	}

	// Postgres
	pg, err := db.NewSQL(dsn)
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	// Redis
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("redis parse url: %v", err)
	}
	rdb := redis.NewClient(opt)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("redis ping: %v", err)
	}

	// core components
	pm := payments.NewManager(pg)
	mm := match.NewMatchmaker(rdb, pg)

	// API server
	router := gin.Default()
	api.RegisterRoutes(router, pm, mm)

	addr := ":8080"
	log.Printf("versus-service listening on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
