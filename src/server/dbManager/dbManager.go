package dbManager

import (
	"context"
	"log/slog"

	pbManagement "go-sdap/src/proto/management"
	"go-sdap/src/server/helper"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbManager struct {
	logger *slog.Logger
}

var (
	clientOptions = options.Client().ApplyURI("mongodb://admin:admin@localhost:27017")
	dbClient      *mongo.Client
	db            *mongo.Database
	ctx           context.Context
)

func New(logger *slog.Logger) *DbManager {
	var err error
	dbClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Error("failed to connect to database", "error", err)
	}

	db = dbClient.Database("sdap")

	return &DbManager{
		logger: logger,
	}
}

func (d *DbManager) AddUsers(users []*pbManagement.User) pbManagement.Status {
	if db == nil {
		return pbManagement.Status_STATUS_ERROR
	}

	usersCollection := db.Collection("users")

	bsonArray := make([]interface{}, 0)
	for _, user := range users {

		bsonDoc, err := helper.ProtoToBSON(user)

		if err != nil {
			d.logger.Error("Error converting user to bson", "error", err)
			continue
		}

		bsonArray = append(bsonArray, bsonDoc)
	}

	_, err := usersCollection.InsertMany(ctx, bsonArray)
	if err != nil {
		d.logger.Error("Error inserting to database", "error", err)

		return pbManagement.Status_STATUS_ERROR
	}

	return pbManagement.Status_STATUS_OK
}

func (d *DbManager) Disconnect() {
	if dbClient != nil {
		dbClient.Disconnect(ctx)
	}
}

func (d *DbManager) Ping() error {
	if dbClient != nil {
		return dbClient.Ping(ctx, nil)
	}
	return nil
}

func (d *DbManager) Reconnect() error {
	if dbClient != nil {
		dbClient.Disconnect(ctx)

		var err error
		dbClient, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			d.logger.Error("failed to reconnect to database", "error", err)
		}
		return err
	}
	return nil
}
