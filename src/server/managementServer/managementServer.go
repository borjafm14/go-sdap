package managementServer

import (
	"context"
	pb "go-sdap/src/proto/management"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"
)

type managementServer struct {
	pb.UnimplementedManagementServer
	logger *slog.Logger
}

var (
	clientOptions = options.Client().ApplyURI("mongodb://admin:admin@localhost:27017")
	dbClient      *mongo.Client
	db            *mongo.Database
)

func New(logger *slog.Logger) *managementServer {
	return &managementServer{
		logger: logger,
	}
}

func (s *managementServer) Connect(ctx context.Context, in *pb.SessionRequest) (*pb.SessionResponse, error) {
	logger := s.logger.With("RPC", "Connect")
	logger.Info("Incoming request", "req", in)

	var err error
	dbClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Error("failed to connect to database", "error", err)
	}

	db = dbClient.Database("sdap")

	return &pb.SessionResponse{
		Token:  "1234",
		Status: pb.Status_STATUS_OK,
	}, nil
}

func (s *managementServer) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	logger := s.logger.With("RPC", "GetUser")
	logger.Info("Incoming request", "req", in)

	var user *pb.User
	return &pb.UserResponse{
		User:   user,
		Status: pb.Status_STATUS_OK,
	}, nil
}

func (s *managementServer) ListUsers(ctx context.Context, in *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	logger := s.logger.With("RPC", "ListUsers")
	logger.Info("Incoming request", "req", in)

	var users []*pb.User
	return &pb.ListUsersResponse{
		Users:  users,
		Status: pb.Status_STATUS_OK,
	}, nil
}

func (s *managementServer) ModifyUsers(ctx context.Context, in *pb.ModifyUsersRequest) (*pb.ModifyUsersResponse, error) {
	logger := s.logger.With("RPC", "ModifyUsers")
	logger.Info("Incoming request", "req", in)

	var users []*pb.User
	return &pb.ModifyUsersResponse{
		Users:  users,
		Status: pb.Status_STATUS_OK,
	}, nil
}

func (s *managementServer) AddUsers(ctx context.Context, in *pb.AddUsersRequest) (*pb.AddUsersResponse, error) {
	logger := s.logger.With("RPC", "AddUsers")
	logger.Info("Incoming request", "req", in)

	if db != nil {
		usersCollection := db.Collection("users")

		bsonArray := make([]interface{}, 0)
		for _, user := range in.Users {

			logger.Info("Contenido del usuario antes de Marshal", "user", user)

			// protobuf to json
			userJSON, err := protojson.Marshal(user)
			if err != nil {
				logger.Error("Error converting to JSON", "error", err)
				continue
			}

			// json to bson
			var bsonDoc bson.M
			if err := bson.UnmarshalExtJSON(userJSON, true, &bsonDoc); err != nil {
				logger.Error("Error converting to BSON", "error", err)
			}

			bsonArray = append(bsonArray, bsonDoc)
		}

		_, err := usersCollection.InsertMany(ctx, bsonArray)
		if err != nil {
			logger.Error("Error inserting to database", "error", err)
		}
	}

	return &pb.AddUsersResponse{
		Status: pb.Status_STATUS_OK,
	}, nil
}

func (s *managementServer) DeleteUsers(ctx context.Context, in *pb.DeleteUsersRequest) (*pb.DeleteUsersResponse, error) {
	logger := s.logger.With("RPC", "DeleteUsers")
	logger.Info("Incoming request", "req", in)

	return &pb.DeleteUsersResponse{
		Status: pb.Status_STATUS_OK,
	}, nil
}

func (s *managementServer) Disconnect(ctx context.Context, in *pb.DisconnectRequest) (*emptypb.Empty, error) {
	logger := s.logger.With("RPC", "Disconnect")
	logger.Info("Incoming request", "req", in)

	if dbClient != nil {
		dbClient.Disconnect(ctx)
	}

	return &emptypb.Empty{}, nil
}
