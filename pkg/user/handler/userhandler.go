package handler

import (
	"context"
	"log"

	adminpb "github.com/ratheeshkumar/microservice_grpc_adminV1/pkg/pb"
	userpb "github.com/ratheeshkumar/microservice_grpc_adminV1/pkg/user/pb"
)

func CreateUserHandler(client userpb.UserServiceClient, p *adminpb.User) (*userpb.CommonResponse, error) {
	ctx := context.Background()

	response, err := client.CreateUser(ctx, &userpb.UserCreate{
		Username: p.Username,
		Email:    p.Email,
		Password: p.Password,
		Role:     "user",
	})

	if err != nil {
		log.Printf("error while creating user")
		return nil, err
	}
	return response, nil
}

func FindAllUserHandler(client userpb.UserServiceClient, p *adminpb.NoParam) (*userpb.Users, error) {
	ctx := context.Background()
	response, err := client.FindAllUsers(ctx, &userpb.FetchAll{})
	if err != nil {
		log.Printf("error getting users")
		return nil, err
	}
	return response, nil
}

func EditUserHandler(client userpb.UserServiceClient, p *adminpb.User) (*userpb.UserCreate, error) {
	ctx := context.Background()

	response, err := client.EditUser(ctx, &userpb.UserCreate{
		Username: p.Username,
		Email:    p.Email,
		Password: p.Password,
		Role:     p.Role,
	})
	if err != nil {
		log.Printf("error editing user details")
		return nil, err
	}
	return response, nil
}

func DeleteUserHandler(client userpb.UserServiceClient, p *adminpb.UserID) (*userpb.CommonResponse, error) {
	ctx := context.Background()

	response, err := client.DeleteUser(ctx, &userpb.ID{
		Id: p.Id,
	})
	if err != nil {

		log.Printf("error while deleting user")
		return nil, err
	}
	return response, nil
}

func FindUserByIDHandler(client userpb.UserServiceClient, p *adminpb.UserID) (*userpb.UserDetail, error) {
	ctx := context.Background()

	response, err := client.FindUserByID(ctx, &userpb.ID{
		Id: p.Id,
	})
	if err != nil {
		log.Printf("error while finding user")
		return nil, err
	}
	return response, nil
}

func FindUserByEmailHandler(client userpb.UserServiceClient, p *adminpb.UserRequest) (*userpb.UserDetail, error) {
	ctx := context.Background()

	response, err := client.FindUserByEmail(ctx, &userpb.Email{
		Email: p.Email,
	})
	if err != nil {
		log.Printf("error while finding user")
		return nil, err
	}
	return response, nil
}
