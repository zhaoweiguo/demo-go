package regexp

import (
	"fmt"
	"regexp"
	"testing"
)

func TestDemo1(t *testing.T) {
	re, _ := regexp.Compile("a(x*)b(y*)c")
	str := "axxbyyyccc   axxbyycc"
	matchonestr := "aaaaaaxxbdsafdsc"
	notmatchstr := "not match"
	index := -1

	// 得到从左数第一个匹配相关
	// 1. 前面2个数为匹配成功的下标
	// 2. 后面4个数的为两个()内的下标数即:(x*)与(y*)
	rtn1 := re.FindStringSubmatchIndex(str) // [0 8 1 3 4 7]
	fmt.Printf("rtn1:%v\n", rtn1)
	notmatchrtn1 := re.FindStringSubmatchIndex(notmatchstr) // []
	fmt.Printf("not match rtn1:%v\n", notmatchrtn1)
	matchonertn1 := re.FindStringSubmatchIndex(matchonestr) // []
	fmt.Printf("match once rtn1:%v\n", matchonertn1)

	rtn2 := re.FindStringSubmatch(str)
	fmt.Printf("rtn2:%v\n", rtn2) // rtn2:[axxbyyyc xx yyy]
	rtn3 := re.FindStringIndex(str)
	fmt.Printf("rtn3:%v\n", rtn3) // rtn3:[0 8]
	rtn4 := re.FindString(str)
	fmt.Printf("rtn4:%v\n", rtn4) // rtn4:axxbyyyc

	// 得到匹配的全部
	// index为取第几个匹配, 为-1时得全部的
	rtn5 := re.FindAllStringSubmatchIndex(str, index) // [[0 8 1 3 4 7] [13 20 14 16 17 19]]
	fmt.Printf("rtn5:%v\n", rtn5)
	rtn6 := re.FindAllStringSubmatch(str, index) // [[axxbyyyc xx yyy] [axxbyyc xx yy]]
	fmt.Printf("rtn6:%v\n", rtn6)
	rtn7 := re.FindAllStringIndex(str, index) // [[0 8] [13 20]]
	fmt.Printf("rtn7:%v\n", rtn7)
	rtn8 := re.FindAllString(str, index) // [axxbyyyc axxbyyc]
	fmt.Printf("rtn8:%v\n", rtn8)

}
