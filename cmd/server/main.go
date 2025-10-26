package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/geoffreyhinton/crypto_trading_v2/internal/api"
	"github.com/geoffreyhinton/crypto_trading_v2/internal/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Initialize logger
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
	// Load configuration
	config, err := loadConfig()
	if err != nil {
		logger.WithError(err).Fatal("Failed to load configuration")
	}

	// Initialize database
	db, err := initDatabase(config.DatabaseURL, logger)
	if err != nil {
		logger.Fatalf("Failed to initialize database: %v", err)
	}
	apiServer := api.NewServer(db, logger)

	// Start the server in a goroutine
	go func() {
		logger.Info("Starting API server on :8080")
		if err := apiServer.Start(":8080"); err != nil {
			logger.WithError(err).Fatal("Failed to start API server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")
}

type Config struct {
	DatabaseURL string
}

func loadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")
	// Set defaults
	viper.SetDefault("database_url", "postgres://user:password@localhost/crypto_exchange_v2?sslmode=disable")
	// Read from environment variables
	viper.AutomaticEnv()
	// Try to read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}
	config := &Config{
		DatabaseURL: viper.GetString("database_url"),
	}
	return config, nil
}
func initDatabase(databaseUrl string, logger *logrus.Logger) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		logger.WithError(err).Error("Failed to connect to database")
		return nil, err
	}
	err = models.AutoMigrate(db)
	if err != nil {
		logger.WithError(err).Error("Failed to migrate database")
		return nil, err
	}
	logger.Info("Database migrations completed successfully")
	return db, nil
}
