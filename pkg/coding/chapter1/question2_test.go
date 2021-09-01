package chapter1

import "testing"

func TestIsPermutation(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "success", args: args{str1: "abc", str2: "cba"}, want: true},
		{name: "success: a char", args: args{str1: "a", str2: "a"}, want: true},
		{name: "success: empty", args: args{str1: "", str2: ""}, want: true},
		{name: "success: empty char", args: args{str1: "  ", str2: "  "}, want: true},
		{name: "fail: include other char", args: args{str1: "abc", str2: "abd"}, want: false},
		{name: "fail: the number of a certain char is different", args: args{str1: "abcd", str2: "aabc"}, want: false},
		{name: "fail: length is different", args: args{str1: "123", str2: "1234"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPermutation(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("IsPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
