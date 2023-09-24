package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Port      string
	Templates string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env is not found")
	}

	var exists bool

	Port, exists = os.LookupEnv("PORT")
	if !exists {
		log.Fatal("Port not found")
	}

	Templates, exists = os.LookupEnv("TEMPLATES")
	if !exists {
		log.Fatal("Templates not found")
	}
}
