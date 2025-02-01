package config

import (
	"fmt"
	"go.uber.org/zap"

	"go-modules-api/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global database instance
var DB *gorm.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase() {
	log := utils.Logger.Named("database")

	// Build DSN connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		Env.DbHost,
		Env.DbUser,
		Env.DbPass,
		Env.DbName,
		Env.DbPort,
	)

	log.Info("Connecting to the database", zap.String("host", Env.DbHost), zap.String("db_name", Env.DbName))

	// Open database connection
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", zap.Error(err))
	}

	DB = database
	log.Info("Database connection established successfully!")

}
