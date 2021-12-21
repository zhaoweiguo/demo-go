package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"log"
)

var allchars = "abcdefghijklmnopqrstuvwxyz"

type WordDict struct {
	children map[byte]*WordDict
	exist    bool
}

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func newWordDict() *WordDict {
	return &WordDict{
		children: make(map[byte]*WordDict),
	}
}

func (dict *WordDict) Add(word string) {
	h := sha1.New()
	h.Write([]byte(word))
	bs := fmt.Sprintf("%x", h.Sum(nil))
	log.Println(bs)
	doAdd(dict, bs)
	log.Println(dict)

}

func doAdd(index *WordDict, word string) {
	if word == "" {
		index.exist = true
		return
	}
	//log.Println(index)

	b := word[0]
	//log.Println(b)
	_, exist := index.children[b]
	if !exist {
		child := newWordDict()
		index.children[b] = child
		//log.Println(index)
	}
	doAdd(index.children[b], word[1:])
	//log.Println(index)
}

func (dict *WordDict) Search(word string) bool {
	newWords := replaceRegular([]string{word})
	log.Println(newWords)
	for _, w := range newWords {
		h := sha1.New()
		h.Write([]byte(w))
		bs := fmt.Sprintf("%x", h.Sum(nil))

		match := doSearch(dict, bs)
		if match {
			return true
		}
	}
	return false
}

// 把.转化成26个字母对应的数组
func replaceRegular(words []string) []string {
	newWords := []string{}
	flag := false
	log.Println(words)
	for _, word := range words {
		bt := new(bytes.Buffer)
		if flag { // 如果没有.，只需要执行一次
			break
		}
		flag = true
		for i := 0; i < len(word); i++ {
			b := word[i]
			log.Println(b)
			if b == '.' {
				for j := 0; j < len(allchars); j++ {
					bt2 := new(bytes.Buffer)
					bt2.WriteString(bt.String())
					bt2.WriteByte(allchars[j])
					bt2.WriteString(word[i+1:])
					newWords = append(newWords, bt2.String())
				}
				log.Println(newWords)
				flag = false
				break
			} else {
				bt.WriteByte(b)
			}
		}
	}
	if flag {
		return words
	}
	return replaceRegular(newWords)
}

func doSearch(index *WordDict, word string) bool {
	log.Println(index, word)
	if word == "" {
		return index.exist
	}
	b := word[0]
	log.Println(b)
	if b == '.' {
		for i := 0; i < len(allchars); i++ {
			index = index.children[allchars[i]]
			rtn := doSearch(index, word[1:])
			if rtn == true {
				return true
			}
		}
		return false
	} else {
		_, exist := index.children[b]
		if !exist {
			return false
		} else {
			return doSearch(index.children[b], word[1:])
		}
	}
}

func main() {
	wd := newWordDict()
	wd.Add("abc")
}
