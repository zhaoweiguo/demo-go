package main

import(
	"fmt"
)

func main() {
	tail1 := "v1"
	tail2 := "v2"
	head := []string{"head1", "head2", "head3"}

	result := append(head, tail1,tail2)
	fmt.Printf("result1:%s\n", result)
}
