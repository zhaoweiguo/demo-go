package main

import (
	"fmt"
	"reflect"
)

// Bird is a nice type.
type Bird struct {
	Name           string
	LifeExpectance int
}

// Fly is fly.
func (b *Bird) Fly() {
	fmt.Println("I am flying...")
}

func main() {
	sparrow := &Bird{"sparrow", 3}
	s := reflect.ValueOf(sparrow).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
