package developerbuysland

import (
	"testing"
)

func Test_developerBuysLand(t *testing.T) {
	type args struct {
		data [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				data: [][]int{{1, 2, 3}, {2, 1, 3}, {1, 2, 3}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := developerBuysLand(tt.args.data); got != tt.want {
				t.Errorf("developerBuysLand() = %v, want %v", got, tt.want)
			}
		})
	}
}
