package cmd

import (
	"log"
	"modules/v2/modules/service"
	"modules/v2/pkg/plugin/v1/greeter"
	"net"

	"google.golang.org/grpc"
)

const host = "0.0.0.0:5000"

type Server struct {
}

func (s *Server) Listen() error {
	_listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	_grpc := grpc.NewServer()
	greeter.RegisterGreeterServer(_grpc, &service.Greeter{})
	log.Printf("grpc listening on %s\n", _listener.Addr())
	return _grpc.Serve(_listener)
}

func NewServer() *Server {
	return &Server{}
}
