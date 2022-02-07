package p61

import (
	. "leetcode/structure"
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"1", args{NewList(1, 2, 3, 4, 5), 1}, NewList(5, 1, 2, 3, 4)},
		{"2", args{NewList(1, 2, 3, 4, 5), 2}, NewList(4, 5, 1, 2, 3)},
		{"3", args{NewList(0, 1, 2), 1}, NewList(2, 0, 1)},
		{"4", args{NewList(0, 1, 2), 2}, NewList(1, 2, 0)},
		{"5", args{NewList(0, 1, 2), 3}, NewList(0, 1, 2)},
		{"6", args{NewList(0, 1, 2), 4}, NewList(2, 0, 1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
