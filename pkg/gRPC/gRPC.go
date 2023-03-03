package gRPC

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func ServeGRPC(service string, port string, server *grpc.Server) {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(service, "Failed to listen on port", port, err)
		panic(err)
	}

	log.Fatal(service, "Starting gRPC server on port", port)
	err = server.Serve(lis)
	if err != nil {
		log.Fatal(service, "Failed to serve gRPC server", err)
		panic(err)
	}
}
