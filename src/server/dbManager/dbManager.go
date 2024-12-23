package dbManager

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	pbManagement "go-sdap/src/proto/management"
	pbSdap "go-sdap/src/proto/sdap"
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

		// Generate random password
		password := helper.GeneratePassword()
		user.Password = &password

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

func (d *DbManager) ChangePassword(username string, old_password string, new_password string) pbSdap.Status {
	if db == nil {
		return pbSdap.Status_STATUS_ERROR
	}

	// check new password requirements
	if !helper.ValidatePassword(new_password) {
		d.logger.Warn("New password does not meet requirements", "password", new_password)
		return pbSdap.Status_STATUS_ERROR
	}

	usersCollection := db.Collection("users")

	filter := bson.M{"username": username, "password": old_password}
	update := bson.M{"$set": bson.M{"password": new_password}}

	result, err := usersCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		d.logger.Error("Error updating password in database", "error", err)
		return pbSdap.Status_STATUS_ERROR
	}

	if result.ModifiedCount == 0 {
		d.logger.Warn("Password not updated, username or old password is incorrect", "username", username)
		return pbSdap.Status_STATUS_ERROR
	}

	d.logger.Info("Password updated successfully", "username", username)
	return pbSdap.Status_STATUS_OK

}

func (d *DbManager) Authenticate(username string, password string) (*pbSdap.User, pbSdap.Status) {
	if db == nil {
		return nil, pbSdap.Status_STATUS_ERROR
	}

	usersCollection := db.Collection("users")

	// TODO get more or less user fields depending on user's sdap_role

	filter := bson.M{"username": username, "password": password}
	var userBson bson.M
	err := usersCollection.FindOne(ctx, filter).Decode(&userBson)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			d.logger.Error("Username not found", "username", username)
		} else {
			d.logger.Error("Authenticate error", "error", err)
		}

		return nil, pbSdap.Status_STATUS_ERROR
	}

	user := &pbSdap.User{}
	err = helper.BSONToProto(userBson, user)

	if err != nil {
		d.logger.Error("Error converting user from bson", "error", err)
		return nil, pbSdap.Status_STATUS_ERROR
	}

	return user, pbSdap.Status_STATUS_OK
}

func (d *DbManager) ListUsers(username *string, filters []*pbManagement.Filter) ([]*pbManagement.User, pbManagement.Status) {
	if db == nil {
		return nil, pbManagement.Status_STATUS_ERROR
	}

	if username != nil {
		// if username is provided, return only that user
		status, user := d.GetUser(*username)
		return []*pbManagement.User{user}, status
	}

	// else, search users that match filter

	searchFields := bson.M{}
	for _, filter := range filters {
		if filter == nil {
			continue
		}

		fieldName, err := helper.ManagementCharacteristicToJSON(filter.Characteristic)
		if err != nil {
			d.logger.Error("Unknown characteristic in filter", "characteristic", filter.Characteristic)
			continue
		}

		searchFields[fieldName] = filter.Value
	}

	if len(searchFields) == 0 {
		d.logger.Warn("No valid filters provided")
		return nil, pbManagement.Status_STATUS_ERROR
	}

	usersCollection := db.Collection("users")

	cursor, err := usersCollection.Find(ctx, searchFields)
	if err != nil {
		d.logger.Error("Error querying database", "error", err)
		return nil, pbManagement.Status_STATUS_ERROR
	}

	users := make([]*pbManagement.User, 0)
	for cursor.Next(ctx) {
		var userBson bson.M
		err := cursor.Decode(&userBson)
		if err != nil {
			d.logger.Error("Error decoding user", "error", err)
			continue
		}

		user := &pbManagement.User{}
		err = helper.BSONToProto(userBson, user)
		if err != nil {
			d.logger.Error("Error converting user from bson", "error", err)
			continue
		}

		users = append(users, user)
	}

	return users, pbManagement.Status_STATUS_OK
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

		fieldName, err := helper.ManagementCharacteristicToJSON(filter.Characteristic)
		if err != nil {
			d.logger.Error("Unknown characteristic in filter", "characteristic", filter.Characteristic)
			continue
		}

		if fieldName == "password" && !helper.ValidatePassword(filter.Value) {
			d.logger.Warn("New password does not meet requirements", "password", filter.Value)
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

func (d *DbManager) GetCharacteristics(username string, characteristics []pbSdap.Characteristic) (*pbSdap.User, pbSdap.Status) {
	if db == nil {
		return nil, pbSdap.Status_STATUS_ERROR
	}

	usersCollection := db.Collection("users")

	projection := bson.M{}
	for _, characteristic := range characteristics {
		fieldName, err := helper.SdapCharacteristicToJSON(characteristic)
		if err != nil {
			d.logger.Error("Unknown characteristic", "characteristic", characteristic)
			continue
		}
		projection[fieldName] = 1
	}
	projection["username"] = 1 // Ensure username is always included

	if len(projection) == 0 {
		d.logger.Warn("No valid characteristics provided")
		return nil, pbSdap.Status_STATUS_ERROR
	}

	filter := bson.M{"username": username}
	opts := options.FindOne().SetProjection(projection)

	var userBson bson.M
	err := usersCollection.FindOne(ctx, filter, opts).Decode(&userBson)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			d.logger.Error("Username not found", "username", username)
			return nil, pbSdap.Status_STATUS_USER_NOT_FOUND
		}
		d.logger.Error("Error querying database", "error", err)
		return nil, pbSdap.Status_STATUS_ERROR
	}

	user := &pbSdap.User{}
	err = helper.BSONToProto(userBson, user)
	if err != nil {
		d.logger.Error("Error converting user from bson", "error", err)
		return nil, pbSdap.Status_STATUS_ERROR
	}

	return user, pbSdap.Status_STATUS_OK
}

func (d *DbManager) GetMemberOf(username string) ([]string, pbSdap.Status) {
	if db == nil {
		return nil, pbSdap.Status_STATUS_ERROR
	}

	usersCollection := db.Collection("users")

	filter := bson.M{"username": username}
	projection := bson.M{"member_of": 1, "_id": 0}

	var result bson.M
	err := usersCollection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			d.logger.Error("Username not found", "username", username)
			return nil, pbSdap.Status_STATUS_USER_NOT_FOUND
		}
		d.logger.Error("Error querying database", "error", err)
		return nil, pbSdap.Status_STATUS_ERROR
	}

	memberOfBson, ok := result["member_of"].(bson.A)
	if !ok {
		d.logger.Warn("Field 'member_of' not found or invalid type", "username", username)
		return nil, pbSdap.Status_STATUS_ERROR
	}

	memberOf := make([]string, len(memberOfBson))
	for i, v := range memberOfBson {
		memberOf[i], ok = v.(string)
		if !ok {
			d.logger.Warn("Invalid type in 'member_of' array", "username", username)
			return nil, pbSdap.Status_STATUS_ERROR
		}
	}
	if !ok {
		d.logger.Warn("Field 'member_of' not found or invalid type", "username", username)
		return nil, pbSdap.Status_STATUS_ERROR
	}

	return memberOf, pbSdap.Status_STATUS_OK
}

func (d *DbManager) ChangeUsername(oldUsername, newUsername string) pbManagement.Status {
	if db == nil {
		return pbManagement.Status_STATUS_ERROR
	}

	usersCollection := db.Collection("users")

	// check if new username already exists
	filter := bson.M{"username": newUsername}
	count, err := usersCollection.CountDocuments(ctx, filter)
	if err != nil {
		d.logger.Error("Error checking for existing user", "username", newUsername, "error", err)
		return pbManagement.Status_STATUS_ERROR
	}
	if count > 0 {
		d.logger.Warn("New username already exists", "username", newUsername)
		return pbManagement.Status_STATUS_ERROR
	}

	filter = bson.M{"username": oldUsername}
	update := bson.M{"$set": bson.M{"username": newUsername}}

	_, err = usersCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		d.logger.Error("Error updating username in database", "error", err)
		return pbManagement.Status_STATUS_ERROR
	}

	d.logger.Info("Username updated successfully", "oldUsername", oldUsername, "newUsername", newUsername)
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
