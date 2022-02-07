package structure

import (
	"reflect"
	"testing"
)

func TestNewList(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"1", args{[]int{}}, nil},
		{"2", args{[]int{1}}, &ListNode{1, nil}},
		{"3", args{[]int{1, 2}}, &ListNode{1, &ListNode{2, nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewList(tt.args.vals...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewList() = %v, want %v", got, tt.want)
			}
		})
	}
}
