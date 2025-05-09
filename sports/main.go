package main

import (
	"log"
	"net"

	"github.com/SylvanSol/Entain_Test/sports/proto/sports"
	"github.com/SylvanSol/Entain_Test/sports/service"
	"google.golang.org/grpc"
)

func main() {
	const port = ":9100"

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	sports.RegisterSportsServer(grpcServer, service.NewSportsService())

	log.Printf("Sports gRPC server listening on %s", port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
