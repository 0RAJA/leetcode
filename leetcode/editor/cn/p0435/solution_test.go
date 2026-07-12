package p0435

import (
	"testing"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name:    "[[-52,31],[-73,-26],[82,97],[-65,-11],[-62,-49],[95,99],[58,95],[-31,49],[66,98],[-63,2],[30,47],[-40,-26]]",
			args:    args{intervals: [][]int{{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26}}},
			wantRes: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := eraseOverlapIntervals(tt.args.intervals); gotRes != tt.wantRes {
				t.Errorf("eraseOverlapIntervals() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
