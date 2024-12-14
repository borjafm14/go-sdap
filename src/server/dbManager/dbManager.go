package dbManager

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	pbManagement "go-sdap/src/proto/management"
	"go-sdap/src/server/helper"

	"go.mongodb.org/mongo-driver/bson"
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

func (d *DbManager) DeleteUsers(usernames []string) pbManagement.Status {
	if db == nil {
		return pbManagement.Status_STATUS_ERROR
	}

	if len(usernames) == 0 {
		d.logger.Warn("No usernames provided for deletion")
		return pbManagement.Status_STATUS_ERROR
	}

	usersCollection := db.Collection("users")

	filter := bson.M{"username": bson.M{"$in": usernames}}

	_, err := usersCollection.DeleteMany(ctx, filter)
	if err != nil {
		d.logger.Error("Error deleting users from database", "error", err)
		return pbManagement.Status_STATUS_ERROR
	}

	d.logger.Info("Users deleted successfully")

	return pbManagement.Status_STATUS_OK
}

func (d *DbManager) GetUser(username string) (pbManagement.Status, *pbManagement.User) {
	if db == nil {
		return pbManagement.Status_STATUS_ERROR, nil
	}

	usersCollection := db.Collection("users")

	filter := bson.M{"username": username}
	var userBson bson.M
	err := usersCollection.FindOne(ctx, filter).Decode(&userBson)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			d.logger.Error("Username not found", "username", username)
		} else {
			d.logger.Error("GetUser error", "error", err)
		}

		return pbManagement.Status_STATUS_USER_NOT_FOUND, nil
	}

	user := &pbManagement.User{}
	err = helper.BSONToProto(userBson, user)

	if err != nil {
		d.logger.Error("Error converting user from bson", "error", err)
		return pbManagement.Status_STATUS_ERROR, nil
	}

	return pbManagement.Status_STATUS_OK, user

}

func (d *DbManager) AddUsers(users []*pbManagement.User) pbManagement.Status {
	if db == nil {
		return pbManagement.Status_STATUS_ERROR
	}

	usersCollection := db.Collection("users")

	bsonArray := make([]interface{}, 0)
	for _, user := range users {

		if user.FirstName == nil || user.LastName == nil {
			d.logger.Error("User missing firstName or lastName", "user", user)
			continue
		}

		// Generate username by getting the first three characters from first and last name
		// converting to lower case and discarding special characters

		firstName := helper.SanitizeName(*user.FirstName)
		lastName := helper.SanitizeName(*user.LastName)

		usernameBase := firstName[:min(3, len(firstName))] + lastName[:min(3, len(lastName))]
		usernameBase = strings.ToLower(usernameBase)

		var highestUsername string
		filter := bson.M{"username": bson.M{"$regex": fmt.Sprintf("^%s\\d*$", usernameBase)}}
		opts := options.FindOne().SetSort(bson.D{{Key: "username", Value: -1}})
		err := usersCollection.FindOne(ctx, filter, opts).Decode(&highestUsername)

		var username string
		if err == nil {
			// Username already exists, increment the number at the end
			username = helper.GenerateNextUsername(usernameBase, highestUsername)
		} else {
			// Username doesn't exist, use it
			username = usernameBase
		}

		user.Username = &username

		filter = bson.M{"username": *user.Username}
		count, err := usersCollection.CountDocuments(ctx, filter)
		if err != nil {
			d.logger.Error("Error checking for existing user", "username", user.Username, "error", err)
			continue
		}

		if count > 0 {
			d.logger.Warn("User already exists", "username", user.Username)
			continue
		}

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

func (d *DbManager) ModifyUsers(usernames []string, filters []*pbManagement.Filter) pbManagement.Status {
	if db == nil {
		return pbManagement.Status_STATUS_ERROR
	}

	updateFields := bson.M{}
	for _, filter := range filters {
		if filter == nil {
			continue
		}

		fieldName, err := helper.CharacteristicToJSON(filter.Characteristic)
		if err != nil {
			d.logger.Error("Unknown characteristic in filter", "characteristic", filter.Characteristic)
			continue
		}

		updateFields[fieldName] = filter.Value
	}

	if len(updateFields) == 0 {
		d.logger.Warn("No valid updates provided")
		return pbManagement.Status_STATUS_ERROR
	}

	query := bson.M{"username": bson.M{"$in": usernames}}
	update := bson.M{"$set": updateFields}

	usersCollection := db.Collection("users")
	result, err := usersCollection.UpdateMany(ctx, query, update)
	if err != nil {
		d.logger.Error("Error updating users in database", "error", err)
		return pbManagement.Status_STATUS_ERROR
	}

	d.logger.Info("Users updated", "matchedCount", result.MatchedCount, "modifiedCount", result.ModifiedCount)
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
