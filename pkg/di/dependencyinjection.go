package di

import (
	"log"

	"github.com/ratheeshkumar/microservice_grpc_adminV1/config"
	"github.com/ratheeshkumar/microservice_grpc_adminV1/pkg/db"
	"github.com/ratheeshkumar/microservice_grpc_adminV1/pkg/handler"
	"github.com/ratheeshkumar/microservice_grpc_adminV1/pkg/repo"
	"github.com/ratheeshkumar/microservice_grpc_adminV1/pkg/server"
	"github.com/ratheeshkumar/microservice_grpc_adminV1/pkg/services"
	"github.com/ratheeshkumar/microservice_grpc_adminV1/pkg/user"
)

func Init() {
	config.LoadConfig()
	db := db.ConnectDB()
	client, err := user.ClientDial()
	if err != nil {
		log.Fatalf("something went wrong %v", err)
	}
	adminRepo := repo.NewAdminRepo(db)
	adminService := services.NewAdminService(adminRepo, client)
	adminHandler := handler.NewAdminHandler(adminService)
	server.NewGrpcServer(adminHandler)
}
