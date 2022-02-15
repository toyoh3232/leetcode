package p68

import (
	"reflect"
	"testing"
)

func Test_justify(t *testing.T) {
	type args struct {
		words    []string
		maxwidth int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"", args{[]string{"Science", "is", "what", "we"}, 20}, "Science  is  what we"},
		{"", args{[]string{"understand", "well"}, 20}, "understand      well"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := justify(tt.args.words, tt.args.maxwidth); got != tt.want {
				t.Errorf("justify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"", args{[]string{"justification."}, 16}, []string{"justification.  "}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}
