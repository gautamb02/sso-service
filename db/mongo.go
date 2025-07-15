package db

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	cnf "github.com/gautamb02/sso-service/confreader"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoClient wraps MongoDB client and database
type MongoClient struct {
	Client *mongo.Client
	DB     *mongo.Database
	cfg    cnf.MongoConfig
}

// NewMongoClient initializes and returns a new MongoClient
func NewMongoClient(cfg cnf.MongoConfig) (*MongoClient, error) {
	uri := buildMongoURI(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("mongo connection error: %w", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("mongo ping error: %w", err)
	}

	db := client.Database(cfg.Database)
	log.Printf("Connected to MongoDB [%s] at %s:%d (DB: %s)", cfg.Name, cfg.Host, cfg.Port, cfg.Database)

	return &MongoClient{
		Client: client,
		DB:     db,
		cfg:    cfg,
	}, nil
}

// Disconnect closes the MongoDB connection
func (m *MongoClient) Disconnect(ctx context.Context) error {
	if err := m.Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("mongo disconnect error: %w", err)
	}
	log.Printf("Disconnected from MongoDB [%s]", m.cfg.Name)
	return nil
}

// Helper to build connection URI
func buildMongoURI(cfg cnf.MongoConfig) string {
	credentials := ""
	if cfg.User != "" && cfg.Password != "" {
		credentials = fmt.Sprintf("%s:%s@", url.QueryEscape(cfg.User), url.QueryEscape(cfg.Password))
	}
	return fmt.Sprintf("mongodb://%s%s:%d", credentials, cfg.Host, cfg.Port)
}
