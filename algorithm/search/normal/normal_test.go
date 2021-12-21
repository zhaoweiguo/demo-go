package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Test_doSearch1(t *testing.T) {
	index := newWordDict()
	got1 := index.Search("abc")
	log.Println(got1)
}

func TestWordDict_Search(t *testing.T) {
	tests := []struct {
		name   string
		add    string
		search string
		got    bool
	}{
		{"1", "a", "a", true},
		{"2", "abc", "abc", true},
		{"3", "abc", "ab.", true},
		{"4", "abc", "a.c", true},
		{"5", "abc", "a..", true},
		{"6", "abc", ".bc", true},
		{"7", "abc", "ab..", false},
		{"8", "abc", ".b.", true},
		{"9", "abc", ".c.", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dict := newWordDict()
			dict.Add(tt.add)
			got := dict.Search(tt.search)
			assert.Equal(t, got, tt.got)
		})
	}
}

/*
这儿就是我说的高度是500
*/
func Test_doHight(t *testing.T) {
	tests := []struct {
		name   string
		add    string
		search string
		got    int
	}{
		{"1", "a", "a", 1},
		{"2", "abc", "abc", 3},
		{"3", "abcdefghjjjjjjjjjjjj", "ab.", 20},
		{"4", "abcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjj", "ab.", 100},
		{"5", "abcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjjabcdefghjjjjjjjjjjjj", "ab.", 500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dict := newWordDict()
			dict.Add(tt.add)
			got := dict.Hight()
			assert.Equal(t, got, tt.got)
		})
	}
}
