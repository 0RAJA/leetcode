package maxSlidingWindow

import (
	"reflect"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantRet []int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
				k:    3,
			},
			wantRet: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "test2",
			args: args{
				nums: []int{1},
				k:    1,
			},
			wantRet: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("maxSlidingWindow() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
