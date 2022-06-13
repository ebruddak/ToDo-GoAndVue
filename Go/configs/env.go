package configs

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

func EnvMongoConnection() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("error .env")
	}

	mongoIRU := os.Getenv("MONGOURI")
	return mongoIRU
}