package command

import (
	"log"
	"os"
)

func SetEnv(name, value string) {
	log.Printf("setting environment variable %s=%s", name, value)
	err := os.Setenv(name, value)
	if err != nil {
		log.Fatalf("failed setting environment variables")
	}
}
