package factory

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

type ServerOptions struct {
	Network string
	Address string
}

func MakeServer(register func(server *grpc.Server), opts ServerOptions) error {
	listener, err := net.Listen(opts.Network, opts.Address)

	if err != nil {
		return err
	}

	server := grpc.NewServer()

	register(server)

	log.Println("Service started")

	err = server.Serve(listener)

	if err != nil {
		return err
	}

	return nil
}
