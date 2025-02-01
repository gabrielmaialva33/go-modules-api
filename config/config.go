package config

import (
	"os"
	"path/filepath"
	"strconv"

	"go-modules-api/utils"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

const envFilePath = ".env"

type Config struct {
	AppHost     string `envconfig:"APP_HOST" default:""`
	AppPort     int    `envconfig:"APP_PORT" default:"3000"`
	UsePublicIP bool   `envconfig:"USE_PUBLIC_IP" default:"false"`

	DbHost string `envconfig:"DB_HOST" default:"localhost"`
	DbPort string `envconfig:"DB_PORT" default:"5432"`
	DbUser string `envconfig:"DB_USER" default:"postgres"`
	DbPass string `envconfig:"DB_PASSWORD" default:"postgres"`
	DbName string `envconfig:"DB_NAME" default:"postgres"`
}

var Env = &Config{}

func (c *Config) loadEnvFile(log *zap.Logger) {
	envPath := filepath.Clean(envFilePath)
	log.Sugar().Infof("Trying to load ENV vars from %s", envPath)

	if err := godotenv.Load(envPath); err != nil {
		if os.IsNotExist(err) {
			log.Sugar().Errorf("ENV file not found: %s. Please create .env file.", envPath)
		} else {
			log.Fatal("Unknown error while parsing env file.", zap.Error(err))
		}
	}
}

func (c *Config) configureAppHost(ip string, ipBlocked bool, log *zap.Logger) {
	if c.AppHost == "" {
		c.AppHost = "http://" + ip + ":" + strconv.Itoa(c.AppPort)

		if c.UsePublicIP {
			if ipBlocked {
				log.Sugar().Warn("Can't get public IP, using local IP")
			} else {
				log.Sugar().Warn("Using a public IP. Be aware of the security risks while exposing your IP to the internet.")
				log.Sugar().Warn("Use 'APP_HOST' variable to set a domain name")
			}
		}

		log.Sugar().Infof("APP_HOST not set, automatically set to %s", c.AppHost)
	}
}

func (c *Config) initializeConfig(log *zap.Logger) {
	c.loadEnvFile(log)

	if err := envconfig.Process("", c); err != nil {
		log.Fatal("Error while parsing env variables", zap.Error(err))
	}

	ip, err := utils.GetIP(c.UsePublicIP)
	ipBlocked := err != nil
	if ipBlocked {
		log.Error("Error while getting IP", zap.Error(err))
		ip = "localhost"
	}

	c.configureAppHost(ip, ipBlocked, log)
}

func Load(log *zap.Logger) {
	log = log.Named("config")
	defer log.Info("Configuration loaded successfully")

	Env.initializeConfig(log)
}
