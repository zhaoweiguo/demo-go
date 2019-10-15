package main

import (
	"log"
	"os"
)

// FileReader
type FileReader struct {
	f *os.File
}

// NewFileReader *FileReader 构造函数，第二个参数是cleanup function
func NewFileReader(filePath string) (*FileReader, func(), error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	fr := &FileReader{
		f: f,
	}
	fn := func() {
		log.Println("cleanup")
		fr.f.Close()
	}
	return fr, fn, nil
}

func main() {
	_, cleanup, err := InitializeFileReader("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()
}

/*
// go run .
func main() {
	_, cleanup, err := NewFileReader("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()
	// do something

}

*/
