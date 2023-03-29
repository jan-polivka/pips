package hello

import (
	"testing"
)

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

func Test_helloWorld(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"hello worl"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helloWorld()
		})
	}
}
