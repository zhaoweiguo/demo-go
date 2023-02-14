package ioutil

import (
	"github.com/bmizerany/assert"
	"io/ioutil"
	"os"
	"testing"
)

// ioutil.WriteFile会直接覆盖原文件中的内容
func TestWriteFile(t *testing.T) {
	filename := "ignore.txt"
	ioutil.WriteFile(filename, []byte("abc"), os.ModePerm)
	ioutil.WriteFile(filename, []byte("cde"), os.ModePerm)
	c, _ := ioutil.ReadFile(filename)
	assert.Equal(t, string(c), "cde")
}

// ioutil.WriteFile会直接覆盖原文件中的内容
func TestWriteFile2(t *testing.T) {
	filename := "ignore.txt"
	ioutil.WriteFile(filename, []byte("abc\nabc"), os.ModePerm)
	ioutil.WriteFile(filename, []byte("cde"), os.ModePerm)
	c, _ := ioutil.ReadFile(filename)
	assert.Equal(t, string(c), "cde")
}
