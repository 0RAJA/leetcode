package p0003

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "empty", s: "", want: 0},
		{name: "example1", s: "abcabcbb", want: 3},
		{name: "example2", s: "bbbbb", want: 1},
		{name: "example3", s: "pwwkew", want: 3},
		{name: "repeated left boundary", s: "abba", want: 2},
		{name: "space and symbols", s: "a b!a", want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Fatalf("lengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
