package main

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	byt, _ := ioutil.ReadFile("./priv/archive.zip")
	num := int64(len(byt))
	zipReader, _ := zip.NewReader(bytes.NewReader(byt), num)
	for _, _file := range zipReader.File {
		log.Println(_file.Name, _file.FileInfo().Name(), _file.UncompressedSize64, _file.FileInfo().Size())
		if _file.FileInfo().Name() != "a.bin" {
			continue
		}
		f, _ := _file.Open()

		//desfile, err1 := os.OpenFile("a.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		//defer desfile.Close()
		//if err1 == nil {
		//	io.CopyN(desfile, f, int64(_file.UncompressedSize64))
		//}

		c := make([]byte, _file.UncompressedSize64)
		i, err := f.Read(c)
		log.Println(i, err)
		log.Println(string(c))
	}
}
