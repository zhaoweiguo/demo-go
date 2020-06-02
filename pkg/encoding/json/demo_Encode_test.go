package json

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestEncode(t *testing.T) {
	user := User{23, "zhaoweiguo", "admin@zhaoweiguo.com"}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	err := enc.Encode(user)

	enc.SetIndent("", "--->")
	err = enc.Encode(user)

	if err != nil {
		log.Panic(err)
	}
}

type User struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
}
