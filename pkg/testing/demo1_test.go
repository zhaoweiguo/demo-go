package main

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) == 3 {
		t.Log("1+2=3")
	}

	if Add(1, 1) == 3 {
		t.Error("1+1=3")
	}
}

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
