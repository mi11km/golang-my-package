package chapter1

import "testing"

func TestIsTransformedByOneStep(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "success: equal", args: args{str1: "", str2: ""}, want: true},
		{name: "success: 1 char lack", args: args{str1: "pale", str2: "ple"}, want: true},
		{name: "success: 1 char surplus", args: args{str1: "pale", str2: "pales"}, want: true},
		{name: "success: 1 char diff", args: args{str1: "pale", str2: "bale"}, want: true},
		{name: "fail: 2 char diff with equal len", args: args{str1: "pale", str2: "bake"}, want: false},
		{name: "fail: 2 more diff len", args: args{str1: "pale", str2: "pale man"}, want: false},
		{name: "fail: 2 char diff with 1 diff len", args: args{str1: "pale", str2: "place"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTransformedByOneStep(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("IsTransformedByOneStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
