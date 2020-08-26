package config

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

func checkExists(exists bool) {
	if !exists {
		log.Fatalln("Error reading environment variables")
	}
}

// GetDBConfig returns a DBConfig struct after extracting the environment variables
func GetDBConfig() *DBConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error reading environment variables")
	}
	dbHost, exists := os.LookupEnv("DB_HOST")
	checkExists(exists)
	dbName, exists := os.LookupEnv("DB_NAME")
	checkExists(exists)
	dbUser, exists := os.LookupEnv("DB_USER")
	checkExists(exists)
	dbPassword, exists := os.LookupEnv("DB_PASSWORD")
	checkExists(exists)
	dbPort, exists := os.LookupEnv("DB_PORT")
	checkExists(exists)
	return &DBConfig{
		DBHost:     dbHost,
		DBName:     dbName,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBPort:     dbPort,
	}
}

// GetSecretKey returns the secret key stored as an enviroment variable
// required to create an authentication token
func GetSecretKey() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", errors.New("Error loading enviroment variable")
	}
	key, exists := os.LookupEnv("API_SECRET")
	if !exists {
		return "", errors.New("Environment variable not found")
	}
	return key, nil
}

// GetRedisAddr returns the address of the redis service
func GetRedisAddr() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", errors.New("Error loading enviroment variable")
	}
	addr, exists := os.LookupEnv("REDIS_ADDR")
	if !exists {
		return "", errors.New("Environment variable not found")
	}
	return addr, nil
}
