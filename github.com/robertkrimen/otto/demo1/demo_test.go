package demo1

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"

	"github.com/robertkrimen/otto"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func TestDemo(t *testing.T) {
	vm := otto.New()
	vm.Run(`
    abc = 2 + 2;
    console.log("The value of abc is " + abc); // 4
`)
}

func TestDemoFunc(t *testing.T) {
	vm := otto.New()
	value, _ := vm.Call(`[ 1, 2, 3, undefined, 4 ].concat`, nil, 5, 6, 7, "abc")
	log.Println(value.String())
}

func TestDemoFile(t *testing.T) {
	vm := otto.New()
	jscontent, err := ioutil.ReadFile("./demo.js")
	assert.Nil(t, err)
	v, err := vm.Run(string(jscontent))
	assert.Nil(t, err)
	assert.NotNil(t, v)
	v2, err := vm.Call(`jsonParse`, nil, `{"name":"Surpass","age":28,"location":"Shanghai","from":"Wuhan","to":"Nanjing"}`)
	assert.Nil(t, err)
	log.Println(v2)
}
