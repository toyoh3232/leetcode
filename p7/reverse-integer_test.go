package p7

import (
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{2147483647}, 0},
		{"2", args{-1200}, -21},
		{"3", args{-1563847412}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{"32"}, "23"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseString(tt.args.str); got != tt.want {
				t.Errorf("reverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_int32ToStr(t *testing.T) {
	type args struct {
		i int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{32}, "32"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := int32ToStr(tt.args.i); got != tt.want {
				t.Errorf("int32ToStr() = %v, want %v", got, tt.want)
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
		want int32
	}{
		// TODO: Add test cases.
		{"1", args{"32"}, 32},
		{"2", args{"-32"}, -32},
		{"2", args{"0000"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strToInt32(tt.args.str); got != tt.want {
				t.Errorf("strToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}
