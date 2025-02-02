package cmd

import (
	"strconv"

	"go-modules-api/config"
	"go-modules-api/internal/factories"
	"go-modules-api/utils"

	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/cobra"
)

func newSeedCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "seed [number]",
		Short:   "Run database seeds",
		Long:    `This command runs all database seeds to populate the database with fake data.`,
		Aliases: []string{"s", "seeds"},
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Default number of records to seed
			numRecords := 10

			// If an argument is provided, override the default
			if len(args) > 0 {
				n, err := strconv.Atoi(args[0])
				if err != nil {
					log.Fatalf("Invalid number of records: %v", err)
				}
				numRecords = n
			}

			utils.InitLogger()
			logger := utils.Logger.Named("seeds")

			// Format the message before passing it to the logger
			logger.Info("Starting database seeding with " + strconv.Itoa(numRecords) + " records")

			// Load environment variables
			config.Load(logger)

			// Connect to database
			config.ConnectDatabase()

			// Generate and save fake data
			for i := 0; i < numRecords; i++ {
				hubClient := factories.HubClientFactory()
				config.DB.Create(hubClient)
			}

			logger.Info("Database seeding completed successfully!")
		},
	}

	return cmd
}
