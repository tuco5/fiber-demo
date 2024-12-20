package config

import (
	"fmt"

	"github.com/tuco5/fiber-demo/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	helper.ErrorPanic(err)

	fmt.Println("Connected to DB")
	return db
}
