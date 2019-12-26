package main

import (
	"go/build"
	"log"
	"os"
)

func main() {
	static()
	doimport()
}

func static() {
	log.Println("===static==========================")
	log.Println(build.Default.GOROOT, build.Default.GOPATH, build.Default.GOOS, build.Default.GOARCH)
	log.Println(build.Default.BuildTags, build.Default.CgoEnabled, build.Default.Compiler, build.Default.InstallSuffix)
	log.Println(build.Default.ReleaseTags)
	log.Println(build.Default.SrcDirs())
}

func doimport() {
	log.Println("===do import==========================")
	p, _ := build.Default.Import("github.com/zhaoweiguo/demo-go/algorithm/sort/utils", "", build.FindOnly)
	log.Println(p.Name, p.Imports)
	log.Println(p.BinDir, p.Dir)
	log.Println(p.AllTags)

	f, _ := os.Open(p.Dir)
	defer f.Close()

	fileInfos, _ := f.Readdir(0)
	for _, fileInfo := range fileInfos {
		log.Println(fileInfo.Name(), fileInfo.Size(), fileInfo.IsDir())
	}
}
