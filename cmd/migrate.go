package cmd

import (
	"go-modules-api/config"
	"go-modules-api/internal/models"
	"go-modules-api/utils"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func newMigrateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Run database migrations",
		Long:  `This command runs all database migrations using GORM AutoMigrate.`,
		Run: func(cmd *cobra.Command, args []string) {
			utils.InitLogger()
			log := utils.Logger.Named("migrations")

			log.Info("Starting database migrations")

			// Load environment variables
			config.Load(log)

			// Connect to database
			config.ConnectDatabase()

			// Run migrations
			err := config.DB.AutoMigrate(
				&models.HubClient{},
				&models.Role{},
				&models.Module{},
				&models.ModulePermission{},
				&models.EntityRegister{},
				&models.ModuleEntityRegister{},
				&models.Product{},
				&models.EntityRegisterProduct{},
				&models.MenuItem{},
				&models.MenuItemPermission{},
			)
			if err != nil {
				log.Fatal("Migration failed", zap.Error(err))
			}

			log.Info("Database migration completed successfully!")
		},
	}

	return cmd
}
