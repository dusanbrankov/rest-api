package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	appConfig *AppConfig
	once      sync.Once
)

// AppConfig holds the application configuration
type AppConfig struct {
	DatabaseConfig *DatabaseConfig
	Enivronment    string
	Infolog        *log.Logger
	IsProduction   bool
	PublicHost     string
	SMTPConfig     *SMTPConfig
	SessionKey     string
}

type DatabaseConfig struct {
	Name     string
	User     string
	Password string
}

type SMTPConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

var App = GetAppConfig()

// GetAppConfig returns the singleton AppConfig instance
func GetAppConfig() *AppConfig {
		if os.Getenv("ENVIRONMENT") != "production" {
			if err := godotenv.Load(); err != nil {
				log.Fatalf("error loading .env file: %v", err)
			}
		}

		appConfig = &AppConfig{
			DatabaseConfig: &DatabaseConfig{
				Name:     os.Getenv("DB_NAME"),
				User:     os.Getenv("DB_USER"),
				Password: os.Getenv("DB_PASSWORD"),
			},
			Enivronment:    os.Getenv("ENVIRONMENT"),
			Infolog:        log.New(os.Stdout, "info: ", log.LstdFlags),
			IsProduction:   os.Getenv("ENVIRONMENT")  == "production",
			PublicHost:     os.Getenv("PUBLIC_HOST"),
			SessionKey:     os.Getenv("SESSION_KEY"),
			SMTPConfig:     &SMTPConfig{
				Host:     os.Getenv("SMTP_HOST"),
				Port:     os.Getenv("SMTP_PORT"),
				Username: os.Getenv("SMTP_USERNAME"),
				Password: os.Getenv("SMTP_PASS"),
			},
		}

	return appConfig
}

