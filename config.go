package main

import (
	"fmt"
	"os"
)

// Global var ?
var Envs = initConfig()

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWTSecret  string
}

// var Envs =
func initConfig() Config {
	return Config{
		Port:       getEnv("PORT", "3306"),
		DBUser:     getEnv("DB_USER", "min23asdw"),
		DBPassword: getEnv("DB_PASSWORD", "#Min087306"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "projectmanager"),
		JWTSecret:  getEnv("JWT_SECRET", "random"),
	}
}

// ENV lookup
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
