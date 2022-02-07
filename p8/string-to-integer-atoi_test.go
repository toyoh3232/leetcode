package p8

import (
	"testing"
)

func Test_toIntString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"", args{"42"}, "42"},
		{"", args{"    s123abc"}, ""},
		{"", args{"-123abc"}, "-123"},
		{"", args{"-s123abc"}, ""},
		{"", args{"-sabc"}, ""},
		{"", args{"-+42"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toIntString(tt.args.s); got != tt.want {
				t.Errorf("toIntString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strToInt32(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{""}, 0},
		{"", args{"2147483647"}, 2147483647},
		{"", args{"2147483648"}, 2147483647},
		{"", args{"-2147483648"}, -2147483648},
		{"", args{"-2147483649"}, -2147483648},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strToInt32(tt.args.str); got != tt.want {
				t.Errorf("strToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}
