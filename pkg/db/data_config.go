package db

import (
	"log"
	"os"

	"github.com/ratheeshkumar/microservice_grpc_adminV1/pkg/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn, ok := os.LookupEnv("DSN")

	if !ok {
		log.Fatal("loading database env")
	}
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting database:%v", err.Error())
	}
	DB.AutoMigrate(&entities.Admin{})

	// Check if admin already exists
	var existingAdmin entities.Admin
	err = DB.Where("username = ?", "admin").First(&existingAdmin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Fatalf("error checking existing admin: %v", err.Error())
	}

	if err == gorm.ErrRecordNotFound {
		// Create an admin
		admin := entities.Admin{
			Username: "admin",
			Password: "admin123",
		}

		result := DB.Create(&admin)
		if result.Error != nil {
			log.Fatalf("failed to create admin: %v", result.Error)
		}

		log.Println("Admin created successfully")
	} else {
		log.Println("Admin already exists")
	}

	return DB
}
