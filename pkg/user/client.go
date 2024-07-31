package user

import (
	"log"

	userpb "github.com/ratheeshkumar/microservice_grpc_adminV1/pkg/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial() (userpb.UserServiceClient, error) {
	grpc, err := grpc.NewClient(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error dialing to grpc client :8082")
		return nil, err
	}
	log.Printf("successfully connected to client :8082")
	return userpb.NewUserServiceClient(grpc), nil
}
