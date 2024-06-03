package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

var required_envs = [...]string{
	"db_host",
	"db_port",
	"db_user",
	"db_password",
	"db_name",
}

func LoadConfig() {
	log.Println("Loading config...")
	_, filename, _, _ := runtime.Caller(0)
	configDir := filepath.Dir(filename)

	envPath := filepath.Join(configDir, ".env")

	err := godotenv.Load(envPath)

	if err != nil {
		log.Fatalf("Error loading .env file. REMINDER: .env file should be in the config directory")
	}

	for _, env := range required_envs {
		if os.Getenv(env) == "" {
			log.Println("ERROR: not all required .env values presents in the .env file")
			log.Printf("Required values: %v", required_envs)
			log.Fatalf("%s is not set in the .env file.", env)
		}
	}
}
