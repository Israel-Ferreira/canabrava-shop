package config

import (
	"os"

	dotenv "github.com/joho/godotenv"
)

type SqsConfig struct {
	SqsQueueName string
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
	SqsConf *SqsConfig
)

func LoadEnv() {
	dotenv.Load()

	SqsConf = &SqsConfig{
		SqsQueueName: os.Getenv("SQS_QUEUE_NAME"),
		AwsAccessKey: os.Getenv("SQS_ACCESS_KEY"),
		AwsSecretKey: os.Getenv("SQS_SECRET_KEY"),
		AwsRegion:    os.Getenv("SQS_REGION"),
	}

	DbConf = &DbConfig{
		DbName:     os.Getenv("DB_NAME"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
	}

}
