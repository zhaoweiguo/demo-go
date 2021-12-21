package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestWordDict_Add(t *testing.T) {
	type fields struct {
		children map[byte]*WordDict
		exist    bool
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"1", fields{nil, false}, args{"abc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dict := newWordDict()
			dict.Add(tt.args.word)
		})
	}
}

func Test_doSearch1(t *testing.T) {
	index := newWordDict()
	got1 := index.Search("abc")
	log.Println(got1)

}
func Test_doSearch2(t *testing.T) {
	index := newWordDict()
	index.Add("abcd")
	got2 := index.Search("abc.")
	log.Println(got2)
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

func Test_replaceRegular(t *testing.T) {
	tests := []struct {
		name  string
		words []string
	}{
		{"1", []string{"abc"}},
		{"2", []string{"ab."}},
		{"3", []string{"a.c"}},
		{"4", []string{".bc"}},
		{"5", []string{"a.."}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := replaceRegular(tt.words)
			t.Log(got)
		})
	}
}
