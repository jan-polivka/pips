package game

import (
	"reflect"
	"testing"
)

func spawnPip(board *int) *int {
	return board
}

func Test_spawnPip(t *testing.T) {
	type args struct {
		board *int
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spawnPip(tt.args.board); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spawnPip() = %v, want %v", got, tt.want)
			}
		})
	}
}
