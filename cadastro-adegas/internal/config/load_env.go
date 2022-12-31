package config

import (
	"os"

	dotenv "github.com/joho/godotenv"
)

type SqsConfig struct {
	SnsQueueName string
	AwsAccessKey string
	AwsSecretKey string
	AwsRegion    string
}

type DbConfig struct {
	DbHost     string
	DbName     string
	DbPort     string
	DbUser     string
	DbPassword string
}

var (
	DbConf  *DbConfig
	SnsConf *SqsConfig
)

func LoadEnv() {
	dotenv.Load()

	SnsConf = &SqsConfig{
		SnsQueueName: os.Getenv("SNS_QUEUE_NAME"),
		AwsAccessKey: os.Getenv("SNS_ACCESS_KEY"),
		AwsSecretKey: os.Getenv("SNS_SECRET_KEY"),
		AwsRegion:    os.Getenv("SNS_REGION"),
	}

	DbConf = &DbConfig{
		DbName:     os.Getenv("DB_NAME"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
	}

}
