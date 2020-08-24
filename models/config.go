package models

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DBConfig struct contains the database configuration
type DBConfig struct {
	DBHost     string
	DBName     string
	DBUser     string
	DBPassword string
	DBPort     string
}

// GetDBConfig returns a DBConfig struct after extracting the environment variables
func GetDBConfig() *DBConfig {
	err := godotenv.Load(".env")
	if err != nil {
		//log.Fatalln("Error reading environment variables")
		log.Fatalln(err.Error())
		return nil
	}
	return &DBConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBPort:     os.Getenv("DB_PORT"),
	}
}

// GetSecretKey returns the secret key stored as an enviroment variable
// required to create an authentication token
func GetSecretKey() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", errors.New("Error loading enviroment variable")
	}
	return os.Getenv("API_SECRET"), nil
}
