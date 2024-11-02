package sdapClient

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	pb "go-sdap/src/proto/sdap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type sdapClient struct {
	addr   string
	port   int
	secure bool
	token  string
	logger *slog.Logger
}

var (
	conn       *grpc.ClientConn
	err        error
	client     pb.OperationClient
	ctx        context.Context
	cancelFunc context.CancelFunc
)

func New(logger *slog.Logger) *sdapClient {
	return &sdapClient{
		addr:   "",
		port:   -1,
		secure: false,
		token:  "",
		logger: logger,
	}
}

func (s *sdapClient) Connect(hostname string, port int, secure bool) (string, pb.Status, error) {
	s.addr = hostname
	s.port = port
	s.secure = secure

	conn, err = grpc.NewClient(s.addr+":"+strconv.Itoa(s.port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.logger.Error("Connect error", slog.String("err", err.Error()))

		return "", pb.Status_STATUS_ERROR, err
	}

	client = pb.NewOperationClient(conn)
	ctx, cancelFunc = context.WithTimeout(context.Background(), time.Second*60)

	sessionResponse, err := client.Connect(ctx, &pb.SessionRequest{
		Hostname: s.addr,
	})

	s.token = sessionResponse.Token

	go s.HandleShutdown()

	return sessionResponse.Token, sessionResponse.Status, err
}

func (s *sdapClient) Authenticate(username string, pass string) (*pb.User, pb.Status, error) {
	authenticateRequest := &pb.AuthenticateRequest{
		Token:    s.token,
		Username: username,
		Password: pass,
	}

	authenticateResponse, err := client.Authenticate(ctx, authenticateRequest)

	return authenticateResponse.User, authenticateResponse.Status, err
}

func (s *sdapClient) GetCharacteristics(username *string, characteristics []pb.Characteristic) (*pb.User, pb.Status, error) {
	characteristicsRequest := &pb.CharacteristicsRequest{
		Token:           s.token,
		Username:        username,
		Characteristics: characteristics,
	}

	characteristicsResponse, err := client.GetCharacteristics(ctx, characteristicsRequest)

	return characteristicsResponse.User, characteristicsResponse.Status, err
}

func (s *sdapClient) GetMemberOf(username *string) ([]string, pb.Status, error) {
	memberOfRequest := &pb.MemberOfRequest{
		Token:    s.token,
		Username: username,
	}

	memberOfResponse, err := client.GetMemberOf(ctx, memberOfRequest)

	return memberOfResponse.MemberOf, memberOfResponse.Status, err
}

func (s *sdapClient) Disconnect() error {
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

func (s *sdapClient) HandleShutdown() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	s.logger.Debug("HandleShutdown calling Disconnect for token " + s.token)
	s.Disconnect()
}
