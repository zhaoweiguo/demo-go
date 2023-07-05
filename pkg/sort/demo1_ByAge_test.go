package main

import (
	"fmt"
	"sort"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface for []Person based on the Age field.
type ByAge []Person

func (p ByAge) Len() int           { return len(p) }
func (p ByAge) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByAge) Less(i, j int) bool { return p[i].Age > p[j].Age }

func TestByAge(t *testing.T) {
	people := []Person{
		{"Bob", 13},
		{"Joe", 123},
		{"Lucy", 3},
		{"Gordon", 83},
	}

	fmt.Println(people)

	// Stable sorts data while keeping the original order of equal elements.
	sort.Stable(ByAge(people))
	fmt.Println(people)

	// The sort is not guaranteed to be stable.
	sort.Sort(ByAge(people))
	fmt.Println(people)

}
