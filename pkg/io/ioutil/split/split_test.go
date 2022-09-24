package split

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestNormal(t *testing.T) {
	data, err := ioutil.ReadFile("./testdata/normal.txt")
	if err != nil {
		panic(err.Error())
	}

	citys := strings.Split(string(data), "--")

	for _, city := range citys {
		if city == "" {
			continue
		}
		fmt.Println(city)
	}
}

func TestCRLF(t *testing.T) {
	data, err := ioutil.ReadFile("./testdata/crlf.txt")
	if err != nil {
		panic(err.Error())
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		fmt.Println(line)
	}
}
