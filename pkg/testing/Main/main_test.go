package Main

import (
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Main is starting...")
	m.Run()
	log.Println("Main is ending...")
}

func TestAbc(t *testing.T) {
	log.Println("Abc is starting...")
}
