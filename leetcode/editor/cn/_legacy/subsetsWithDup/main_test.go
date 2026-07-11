package main

import (
	"reflect"
	"testing"
)

func Test_subsetsWithDup(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]int
	}{
		// 	测试用例:[1,2,2]
		//			[0]
		// 期望结果:[[],[1],[1,2],[1,2,2],[2],[2,2]]
		// [[],[0]]
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 2},
			},
			wantRes: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
		{
			name: "test2",
			args: args{
				nums: []int{0},
			},
			wantRes: [][]int{{}, {0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := subsetsWithDup(tt.args.nums); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("subsetsWithDup() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
