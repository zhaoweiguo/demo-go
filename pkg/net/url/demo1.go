package main

import(
	"net/url"
	"fmt"
)

func main() {
	u, err := url.Parse("http://www.baidu.com/search?q=dotnet")
	if err != nil {
		panic(err)
	}

	fmt.Printf("url.string(): %s\n", u.String())

}
