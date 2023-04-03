package sha256

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

/*
sha256是golang内置的sha256的散列标准库，可以让我们很容易的生成对应数据的散列值
*/

func TestDemo1(t *testing.T) {
	log.Println("===demo1============================")
	// 第一种调用方法
	sum := sha256.Sum256([]byte("hello world\n"))
	log.Printf("%x\n", sum)

	// 第二种调用方法
	h := sha256.New()
	h.Write([]byte("hello world\n"))
	log.Printf("%x\n", h.Sum(nil))
	log.Printf("%X\n", h.Sum(nil))
}

// 对文件加密
func TestDemo2(t *testing.T) {
	log.Println("===demo2_user============================")
	f, err := os.Open("pkg/crypto/sha256/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	log.Printf("%x\n", h.Sum(nil))
}

func TestDemo3(t *testing.T) {

	secret := "SECc4e8395e49471c7260be9057b9a2926b0082221ddb578786752bebb46f183a6a"

	timestamp := "1612420355000"
	msg := fmt.Sprintf("%s\n%s", timestamp, secret)

	h := sha256.New()
	h.Write([]byte(msg))

	log.Printf("%x", h.Sum(nil))

}
