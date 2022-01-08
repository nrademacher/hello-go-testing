package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Niko")
	want := "Hello, Niko"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
