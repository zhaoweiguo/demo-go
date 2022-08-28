package bytes

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestNormal(t *testing.T) {
	//定义零值Buffer类型变量b
	var b bytes.Buffer
	//使用Write方法为写入字符串
	b.Write([]byte("你好"))

	//这个是把一个字符串拼接到Buffer里
	fmt.Fprint(&b, ",", "http://www.flysnow.org")
	b.Write([]byte(" --->"))
	//把Buffer里的内容打印到终端控制台
	b.WriteTo(os.Stdout)
}

func TestAddHeader(t *testing.T)  {
	body := []byte("xxxxxxxxxxxx")
	sub := []byte("aaaa")
	a := append(sub, body...)
	log.Printf("%s\n", string(a))
}

func TestStringByte(t *testing.T)  {
	num := 17
	str := fmt.Sprintf("%04X", num)
	body := []byte(str)
	log.Println(str, body)
}




