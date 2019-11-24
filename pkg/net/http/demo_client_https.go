package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//url := "https://passport.zhaoweiguo.com/interserver/authen/1.2/getaccountid?lpsust=&realm="
	url := "https://gitcodecloud.zhaoweiguo.com.cn"
	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("==============")
	log.Println(body)

}
