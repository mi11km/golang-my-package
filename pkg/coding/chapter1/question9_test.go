package chapter1

import "testing"

func TestIsRotateString(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "success", args: args{s1: "waterbottle", s2: "erbottlewat"}, want: true},
		{name: "fail: wrong len", args: args{s1: "abcd", s2: "abcde"}, want: false},
		{name: "fail: not rotate", args: args{s1: "waterbottle", s2: "waterbattle"}, want: false},
		{name: "empty", args: args{s1: "", s2: ""}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRotateString(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("IsRotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
