package interfaces

import "github.com/ratheeshkumar/microservice_grpc_adminV1/pkg/entities"

type AdminRepoInter interface {
	FindAdmin(user string) (*entities.Admin, error)
}
