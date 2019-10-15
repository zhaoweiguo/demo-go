package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
)

type UserLine struct {
	User_Id  string
	Create_Time string
	Context struct{
		Name string
	}
}


func main() {

	fi, err := os.Open("./pkg/os/data/user_offline.txt")
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		log.Println(string(line))

		var userline = UserLine{}
		err = json.Unmarshal(line, &userline)
		if err != nil {
			log.Println("error:", err)
		}

		log.Println(userline)

		log.Println("--")
	}
}

