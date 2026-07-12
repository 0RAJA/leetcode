package p0768

import (
	"reflect"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			name:    "ababcbacadefegdehijhklij",
			args:    args{s: "ababcbacadefegdehijhklij"},
			wantRes: []int{9, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := partitionLabels(tt.args.s); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("partitionLabels() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
