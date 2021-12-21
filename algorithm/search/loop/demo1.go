package main

import (
	"crypto/sha1"
)

var allchars = "abcdefghijklmnopqrstuvwxyz"

type WordDict struct {
	children map[byte]*WordDict
	exist    bool
}

func new() WordDict {
	return WordDict{
		children: nil,
		exist:    false,
	}
}

func (dict *WordDict) Add(word string) {
	h := sha1.New()
	h.Write([]byte(word))
	bs := h.Sum(nil)
	index := dict
	for _, b := range bs {
		_, exist := index.children[b]
		if !exist {
			child := &WordDict{
				children: nil,
				exist:    false,
			}
			index.children[b] = child
		}
		index = index.children[b]
	}
	index.exist = true
}

func (dict *WordDict) Search(word string) bool {
	h := sha1.New()
	h.Write([]byte(word))
	bs := h.Sum(nil)

	index := dict
	for _, b := range bs {
		if b == '.' {
			for i := 0; i < len(allchars); i++ {
				index = dict.children[allchars[i]]
			}
		} else {
			_, exist := index.children[b]
			if !exist {
				return false
			} else {
				index = dict.children[b]
			}
		}

	}

	return false
}

func main() {
	wd := new()
	wd.Add("abc")
}
