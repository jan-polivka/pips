package game

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func Test_spawnPip(t *testing.T) {
	type args struct {
		board *int
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		{"test", var board [6]int, board}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spawnPip(tt.args.board); got != tt.want {
				t.Errorf("spawnPip() = %v, want %v", got, tt.want)
			}
		})
	}
}
