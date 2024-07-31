package interfaces

import pb "github.com/ratheeshkumar/microservice_grpc_adminV1/pkg/pb"

type AdminServiceInter interface {
	AdminLoginService(p *pb.AdminRequest) (*pb.AdminResponse, error)
	CreateUserService(p *pb.User) (*pb.AdminResponse, error)
	EditUserService(p *pb.User) (*pb.User, error)
	DeleteUserService(p *pb.UserID) (*pb.AdminResponse, error)
	FindUserByEmailService(p *pb.UserRequest) (*pb.User, error)
	FIndAllUserService(p *pb.NoParam) (*pb.UserList, error)
	FindUserByIDService(p *pb.UserID) (*pb.User, error)
}
