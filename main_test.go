package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(4,6)
	want := 11

	if got != want {
		t.Errorf("Got %q , wanted %q",got,want)
	}
}