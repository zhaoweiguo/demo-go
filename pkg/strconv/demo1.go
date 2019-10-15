package main

import(
	"strconv"
	"fmt"
	"log"
)

func main() {
	test := func(s string) {
		t, err := strconv.Unquote(s)  // 
		if err != nil {
			fmt.Printf("Unquote(%#v): %v\n", s, err)
		} else {
			fmt.Printf("Unquote(%#v) = %v\n", s, t)
		}
	}

	s := `cafe\u0301`
	// If the string doesn't have quotes, it can't be unquoted.
	test(s) // invalid syntax
	test("`" + s + "`")
	test(`"` + s + `"`)

	test(`'\u00e9'`)

	strA := "123"
	intA, err := strconv.ParseInt(strA, 10, 0)
	if (err != nil ) {
		log.Println("[error]", err)
		panic(err)
	}
	log.Println("intA", intA)

}
