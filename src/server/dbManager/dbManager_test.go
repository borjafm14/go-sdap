package dbManager

import (
	"context"
	"os"
	"testing"

	"log/slog"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestNew(t *testing.T) {
	// Mock logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Mock MongoDB client
	clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	ctx = context.TODO()

	// Test successful connection
	dbManager := New(logger)
	if dbManager == nil {
		t.Fatal("Expected dbManager to be non-nil")
	}

	if dbClient == nil {
		t.Fatal("Expected dbClient to be non-nil")
	}

	if db == nil {
		t.Fatal("Expected db to be non-nil")
	}

	// Test failed connection
	clientOptions = options.Client().ApplyURI("mongodb://invalid:invalid@localhost:27017")
	dbManager = New(logger)
	if dbManager == nil {
		t.Fatal("Expected dbManager to be non-nil even on failed connection")
	}
}
