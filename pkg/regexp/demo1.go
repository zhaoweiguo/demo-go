package main

import (
	"fmt"
	"regexp"
)

func main() {
	// matched
	m1, e1 := regexp.MatchString("^/", "/abc")
	fmt.Printf("ismatched:%t, err:%v\n", m1, e1)   // ismatched:true, err:<nil>
	// not matched
	m2, e2 := regexp.MatchString("^/", "ab/c")
	fmt.Printf("ismatched:%t, err:%v\n", m2, e2)   // 	ismatched:false, err:<nil>

	re, _ := regexp.Compile("a(x*)b|(y*)c")
	str := "axxbyyyccc   axxbyycc"
	matchonestr := "aaaaaaxxbdsafdsc"
	notmatchstr := "not match"
	index := -1

	// 得到从左数第一个匹配相关
	// 前两个数为匹配成功的下标
	// 后面的为()内的下标数
	rtn1 := re.FindStringSubmatchIndex(str)
	fmt.Printf("rtn1:%v\n", rtn1)
	notmatchrtn1 := re.FindStringSubmatchIndex(notmatchstr)
	fmt.Printf("not match rtn1:%v\n", notmatchrtn1)
	matchonertn1 := re.FindStringSubmatchIndex(matchonestr)
	fmt.Printf("match once rtn1:%v\n", matchonertn1)

	rtn2 := re.FindStringSubmatch(str)
	fmt.Printf("rtn2:%v\n", rtn2)         // rtn2:[axxbyyyc xx yyy]
	rtn3 := re.FindStringIndex(str)
	fmt.Printf("rtn3:%v\n", rtn3)         // rtn3:[0 8]
	rtn4 := re.FindString(str)
	fmt.Printf("rtn4:%v\n", rtn4)         // rtn4:axxbyyyc

	// 得到匹配的全部
	// index为取第几个匹配, 为-1时得全部的
	rtn5 := re.FindAllStringSubmatchIndex(str, index)
	fmt.Printf("rtn5:%v\n", rtn5)
	rtn6 := re.FindAllStringSubmatch(str, index)
	fmt.Printf("rtn6:%v\n", rtn6)
	rtn7 := re.FindAllStringIndex(str, index)
	fmt.Printf("rtn7:%v\n", rtn7)
	rtn8 := re.FindAllString(str, index)
	fmt.Printf("rtn8:%v\n", rtn8)


}
