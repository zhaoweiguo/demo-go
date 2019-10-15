package main

import(
	"fmt"
	"os"
)

type ApiError struct{
	Msg string
	Err error
}

func (e *ApiError) Error() {
	fmt.Println(e.Err, e.Msg)
}

func main() {

	file, err := os.Open("/tmp/abc")

	fmt.Println(file)

	newerr := &ApiError{"abc", err}
	newerr.Error()

}


