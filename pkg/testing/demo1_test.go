package main

import "testing"

func TestAdd(t *testing.T) {
	if Add(1,2) == 3 {
		t.Log("1+2=3")
	}

	if Add(1,1) == 3 {
		t.Error("1+1=3")
	}
}



