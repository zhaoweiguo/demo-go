package main

var allchars = "abcdefghijklmnopqrstuvwxyz"

type WordDict struct {
	children map[byte]*WordDict
	exist    bool
}

func newWordDict() *WordDict {
	return &WordDict{
		children: make(map[byte]*WordDict),
	}
}

func (dict *WordDict) Add(word string) {
	doAdd(dict, word)
}

func doAdd(index *WordDict, word string) {
	if word == "" {
		index.exist = true
		return
	}

	b := word[0]
	_, exist := index.children[b]
	if !exist {
		child := newWordDict()
		index.children[b] = child
	}
	doAdd(index.children[b], word[1:])
}

func (dict *WordDict) Hight() int {
	return doHight(dict)
}

func doHight(dict *WordDict) int {
	var hightest = 0
	for _, child := range dict.children {
		h := doHight(child) + 1
		if h > hightest {
			hightest = h
		}
	}
	return hightest
}

func (dict *WordDict) Search(word string) bool {
	return doSearch(dict, word)
	return false
}

func doSearch(index *WordDict, word string) bool {
	if index == nil {
		return false
	}
	if word == "" {
		return index.exist
	}
	b := word[0]
	if b == '.' {
		for i := 0; i < len(allchars); i++ {
			index1 := index
			index1 = index1.children[allchars[i]]
			rtn := doSearch(index1, word[1:])
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
