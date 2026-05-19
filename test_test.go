package main

import "testing"

func TestMessage(t *testing.T) {
	m := Message("Manuel")

	if m != "Hello Manuel !!!" {
		t.Fatal()
	}

}
