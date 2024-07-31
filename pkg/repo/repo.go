package repo

import (
	"log"

	"github.com/ratheeshkumar/microservice_grpc_adminV1/pkg/entities"
	inter "github.com/ratheeshkumar/microservice_grpc_adminV1/pkg/repo/interfaces"
	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

// FindAdmin implements interfaces.AdminRepoInter.
func (a *AdminRepo) FindAdmin(user string) (*entities.Admin, error) {
	var admin entities.Admin
	log.Printf("Querying admin with username: %s", user)
	err := a.DB.Where("username=?", user).First(&admin).Error
	return &admin, err
}

func NewAdminRepo(db *gorm.DB) inter.AdminRepoInter {
	return &AdminRepo{
		DB: db,
	}
}
