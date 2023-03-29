package hello

import "testing"

func TestGreeter(t *testing.T) {
	got := Greeter()
	if got != "hello" {
		t.Errorf("did not get greeted")
	}
}

func TestWorlder(t *testing.T) {
	got := Worlder()
	if got != "world" {
		t.Errorf("did not get world")
	}
}
