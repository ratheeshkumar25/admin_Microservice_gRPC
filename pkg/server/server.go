package server

import (
	"log"
	"net"

	"github.com/ratheeshkumar/microservice_grpc_adminV1/pkg/handler"
	pb "github.com/ratheeshkumar/microservice_grpc_adminV1/pkg/pb"
	"google.golang.org/grpc"
)

func NewGrpcServer(handler *handler.AdminHandler) {
	log.Println("connecting to gRPC server")
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal("error creating listener on port 8083")
	}

	grp := grpc.NewServer()
	pb.RegisterAdminServiceServer(grp, handler)

	log.Printf("listening on gRPC server 8083")
	if err := grp.Serve(lis); err != nil {
		log.Fatal("error connecting to gRPC server")

	}
}
