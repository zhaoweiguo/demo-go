package pointer

import (
	"fmt"
	"testing"
)

func TestConst(t *testing.T) {
}

func TestVariable(t *testing.T) {
	cases := []struct {
		name string
	}{
		{"NO1. "},
		{"NO2. "},
		{"NO3. "},
		{"NO4. "},
	}
	for _, c := range cases {
		a := "123"
		fmt.Printf("[%s], %p \n", c.name, &c.name)
		//fmt.Printf("[%s], %p \n", c, &c)
		fmt.Printf("[%s], %p \n", a, &a)
	}

}
