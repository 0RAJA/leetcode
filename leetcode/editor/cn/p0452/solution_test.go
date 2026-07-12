package p0452

import "testing"

func TestFindMinArrowShots(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			name:   "example1",
			points: [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
			want:   2,
		},
		{
			name:   "disjoint",
			points: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			want:   4,
		},
		{
			name:   "touching intervals overlap",
			points: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			want:   2,
		},
		{
			name:   "single balloon",
			points: [][]int{{1, 2}},
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinArrowShots(tt.points); got != tt.want {
				t.Fatalf("findMinArrowShots() = %d, want %d", got, tt.want)
			}
		})
	}
}
