package server

import (
	"google.golang.org/grpc"
	"net"
	"promotioner/cmd/authService/manager"
	auth "promotioner/protobuf"
)

type Server struct {
	SessionManager manager.SessionManagerInterface
	Port           string
}

func NewServer(port string, manager manager.SessionManagerInterface) *Server {
	return &Server{SessionManager: manager, Port: port}
}

func (s Server) Run() error {
	lis, err := net.Listen("tcp", s.Port)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	auth.RegisterAuthServiceServer(server, NewAuthService(s.SessionManager))
	err = server.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
