package p0503

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			name:    "[1,2,1]",
			args:    args{nums: []int{1, 2, 1}},
			wantRes: []int{2, -1, 2},
		},
		{
			name:    "[1,2,3,4,3]",
			args:    args{nums: []int{1, 2, 3, 4, 3}},
			wantRes: []int{2, 3, 4, -1, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := nextGreaterElements(tt.args.nums); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("nextGreaterElements() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
