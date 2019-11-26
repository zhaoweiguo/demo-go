package main

import (
	"fmt"
	"unicode"
)

func main() {
	const lcG = 'g'
	fmt.Printf("%#U\n", unicode.To(unicode.UpperCase, lcG))
	fmt.Printf("%#U\n", unicode.To(unicode.LowerCase, lcG))
	fmt.Printf("%#U\n", unicode.To(unicode.TitleCase, lcG))

	const ucG = 'G'
	fmt.Printf("%#U\n", unicode.To(unicode.UpperCase, ucG))
	fmt.Printf("%#U\n", unicode.To(unicode.LowerCase, ucG))
	fmt.Printf("%#U\n", unicode.To(unicode.TitleCase, ucG))

	fmt.Printf("%#U\n", unicode.ToLower(ucG))
	fmt.Printf("%#U\n", unicode.ToTitle(ucG))
	fmt.Printf("%#U\n", unicode.ToTitle(lcG))
	fmt.Printf("%#U\n", unicode.ToUpper(lcG))

	t := unicode.TurkishCase
	const lci = 'i'
	fmt.Printf("%#U\n", t.ToLower(lci))
	fmt.Printf("%#U\n", t.ToTitle(lci))
	fmt.Printf("%#U\n", t.ToUpper(lci))

	const uci = 'İ'
	fmt.Printf("%#U\n", t.ToLower(uci))
	fmt.Printf("%#U\n", t.ToTitle(uci))
	fmt.Printf("%#U\n", t.ToUpper(uci))

}
