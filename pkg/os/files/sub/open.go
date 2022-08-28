package sub

import (
	"log"
	"os"
)

func Open()  {
	f, err := os.Open("pkg/os/files/sub/a.txt")
	if err != nil {
		panic(err)
	}
	log.Println(f.Name())
}

