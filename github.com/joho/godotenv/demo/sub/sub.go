package sub

import (
	"github.com/joho/godotenv"
	"testing"
)

func GetConf(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		t.Fatal("Error loading .env file")
	}
}
