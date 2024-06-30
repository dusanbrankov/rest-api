package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/dusanbrankov/rest-api/config"
	"github.com/dusanbrankov/rest-api/sqlc"
	"github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Queries *sqlc.Queries
}

var (
	app      *config.AppConfig
	database *sql.DB
	instance *DBConfig
	queries  *sqlc.Queries
	mySQLConfig *mysql.Config
)

// Initialize sets up the application configuration, database
// connection, and SQL queries instance
// This function should be called at the start of the
// application to ensure all necessary configurations and
// connections are properly initialized.
func Initialize() *sql.DB {
	app = config.App
	database = conn()
	instance = &DBConfig{}

	queries = sqlc.New(database)
	instance = &DBConfig{
		Queries: queries,
	}

	return database
}

// NewMySQLStorage creates and returns a new database connection
// using the configuration obtained from getMySQLConfig
func NewMySQLStorage() (*sql.DB) {
	cfg := getMySQLConfig()

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	return db
}

// getMySQLConfig returns a MySQL configuration object
// If the configuration has not been created yet, it initializes it
// and stores it for future reuse.
func getMySQLConfig() *mysql.Config {
	if mySQLConfig == nil {
		mySQLConfig = &mysql.Config{
			User:                 app.DatabaseConfig.User,
			Passwd:               app.DatabaseConfig.Password,
			Net:                  "tcp",
			Addr:                 "127.0.0.1:3306",
			DBName:               app.DatabaseConfig.Name,
			AllowNativePasswords: true,
			ParseTime:            true,
		}
	}
	return mySQLConfig
}

// getDBConfig returns the singleton instance of DBConfig
func GetDBConfig() *DBConfig {
	if instance == nil {
		Initialize()
	}
	return instance
}

// conn connects to the database
func conn() *sql.DB {
	db := NewMySQLStorage()

	maxConns := 12
	db.SetConnMaxLifetime(time.Minute*3)
	db.SetMaxOpenConns(maxConns)
	db.SetMaxIdleConns(maxConns)

	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatalf("failed to ping database: %v", err)
	}

	return db
}

