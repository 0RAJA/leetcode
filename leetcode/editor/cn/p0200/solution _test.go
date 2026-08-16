package p0200

import (
	"testing"
)

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name: "1",
			args: args{grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			},
			wantRes: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := numIslands(tt.args.grid); gotRes != tt.wantRes {
				t.Errorf("numIslands() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
