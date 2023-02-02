package main

import (
	"io/fs"
	"log"
	"os"
	"strings"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	path := "/Users/zhaoweiguo/0backup/3gitdoc/0shares/files"
	fs.WalkDir(os.DirFS(path), ".", func(path string, dir fs.DirEntry, e error) error {
		//log.Println("----", dir.Name())
		if dir == nil { // path is a file other than a dir
			log.Println(path)
		} else if dir.IsDir() && path != "." && strings.HasPrefix(path, ".") { //  skip hidden folders
			log.Println(path)
		} else {
			log.Println(path, dir.Name())
		}
		return nil
	})
}
