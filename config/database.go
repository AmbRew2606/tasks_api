package config

import (
	"log"
	"os"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func LoadConfig() *Config {
	host := os.Getenv("DB_HOST")
	if host == "" {
		log.Fatal("DB_HOST отсутствует")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		log.Fatal("DB_PORT отсутствует")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		log.Fatal("DB_USER отсутствует")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		log.Fatal("DB_PASSWORD отсутствует")
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		log.Fatal("DB_NAME отсутствует")
	}

	return &Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbname,
	}
}
