package intervalsum

import (
	"reflect"
	"testing"
)

func Test_intervalSum(t *testing.T) {
	type args struct {
		nums      []int
		intervals [][2]int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			name: "test1",
			args: args{
				nums:      []int{1, 2, 3, 4, 5},
				intervals: [][2]int{{0, 1}, {1, 3}},
			},
			wantResult: []int{3, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := intervalSum(tt.args.nums, tt.args.intervals); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("intervalSum() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
