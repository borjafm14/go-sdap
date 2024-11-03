package sdapClient

import (
	"context"
	"log/slog"
	"strconv"
	"time"

	pb "go-sdap/src/proto/sdap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type managementClient struct {
	addr   string
	port   int
	secure bool
	token  string
	logger *slog.Logger
}

var (
	conn       *grpc.ClientConn
	err        error
	client     pb.ManagementClient
	ctx        context.Context
	cancelFunc context.CancelFunc
)

func New() *managementClient {
	return &managementClient{
		addr:   "",
		port:   -1,
		secure: false,
		token:  "",
		logger: slog.Default(),
	}
}

func (s *managementClient) Connect(hostname string, port int, secure bool) (string, pb.Status, error) {
	s.addr = hostname
	s.port = port
	s.secure = secure

	conn, err = grpc.NewClient(s.addr+":"+strconv.Itoa(s.port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.logger.Error("Connect error", slog.String("err", err.Error()))

		return "", pb.Status_STATUS_ERROR, err
	}

	client = pb.NewManagementClient(conn)
	ctx, cancelFunc = context.WithTimeout(context.Background(), time.Second*60)

	sessionResponse, err := client.Connect(ctx, &pb.SessionRequest{
		Hostname: s.addr,
	})

	s.token = sessionResponse.Token

	return sessionResponse.Token, sessionResponse.Status, err
}

func (s *managementClient) GetUser(username string) (*pb.User, pb.Status, error) {
	userRequest := &pb.UserRequest{
		Token:    s.token,
		Username: username,
	}

	userResponse, err := client.GetUser(ctx, userRequest)

	return userResponse.User, userResponse.Status, err
}

func (s *managementClient) ListUsers(username *string, filter []*pb.Filter) ([]*pb.User, pb.Status, error) {
	listUsersRequest := &pb.ListUsersRequest{
		Token:    s.token,
		Username: username,
		Filter:   filter,
	}

	listUsersResponse, err := client.ListUsers(ctx, listUsersRequest)

	return listUsersResponse.Users, listUsersResponse.Status, err
}

func (s *managementClient) ModifyUsers(usernames []string, filter []*pb.Filter) ([]*pb.User, pb.Status, error) {
	modifyUsersRequest := &pb.ModifyUsersRequest{
		Token:     s.token,
		Usernames: usernames,
		Filter:    filter,
	}

	modifyUsersResponse, err := client.ModifyUsers(ctx, modifyUsersRequest)

	return modifyUsersResponse.Users, modifyUsersResponse.Status, err
}

func (s *managementClient) AddUsers(users []*pb.User) ([]*pb.User, pb.Status, error) {
	addUsersRequest := &pb.AddUsersRequest{
		Token: s.token,
		Users: users,
	}

	addUsersResponse, err := client.AddUsers(ctx, addUsersRequest)

	return addUsersResponse.Users, addUsersResponse.Status, err
}

func (s *managementClient) DeleteUsers(usernames []string) error {
	deleteUsersRequest := &pb.DeleteUsersRequest{
		Token:     s.token,
		Usernames: usernames,
	}

	_, err := client.DeleteUsers(ctx, deleteUsersRequest)

	return err
}

func (s *managementClient) Disconnect() error {
	_, err := client.Disconnect(ctx, &pb.DisconnectRequest{
		Token: s.token,
	})

	if err != nil {
		s.logger.Error("Disconnect error", slog.String("err", err.Error()))
	}

	if cancelFunc != nil {
		cancelFunc()
	}
	if conn != nil {
		conn.Close()
	}

	return err
}
