package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	AppEnv     string
	AppPort    string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
	DbSsl      string
}

func LoadConfig() *Config {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		logrus.Fatalf("DB_HOST environment variable is not set")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		logrus.Fatalf("DB_PORT environment variable is not set")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		logrus.Fatalf("DB_NAME environment variable is not set")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		logrus.Fatalf("DB_USER environment variable is not set")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		logrus.Fatalf("DB_PASSWORD environment variable is not set")
	}

	dbSsl := os.Getenv("DB_SSL")
	if dbSsl == "" {
		dbSsl = "verify-full"
	}

	return &Config{
		AppEnv:     appEnv,
		AppPort:    appPort,
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbName:     dbName,
		DbUser:     dbUser,
		DbPassword: dbPassword,
		DbSsl:      dbSsl,
	}
}
