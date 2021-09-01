package chapter1

import "testing"

func TestIsPalindromeOfPermutation(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "success", args: args{str: "tact coa"}, want: true},
		{name: "contain Upper", args: args{str: "Tact Coa"}, want: true},
		{name: "all even", args: args{str: "tact ca"}, want: true},
		{name: "contain 4 counts", args: args{str: "tact ca aa"}, want: true},
		{name: "contain 3 counts with one char", args: args{str: "tact ca a"}, want: true},
		{name: "2 odds", args: args{str: "tact coa a"}, want: false},
		{name: "3 odds", args: args{str: "tact coa ai"}, want: false},
		{name: "ja", args: args{str: "よのなかねかお か おかねかなのよ"}, want: true},
		{name: "mark", args: args{str: "!#$%!#$%"}, want: true},
		{name: "one char", args: args{str: "a"}, want: true},
		{name: "empty", args: args{str: ""}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindromeOfPermutation(tt.args.str); got != tt.want {
				t.Errorf("IsPalindromeOfPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalindromeOfPermutationWithBitVector(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "success", args: args{str: "tact coa"}, want: true},
		{name: "contain Upper", args: args{str: "Tact Coa"}, want: true},
		{name: "all even", args: args{str: "tact ca"}, want: true},
		{name: "contain 4 counts", args: args{str: "tact ca aa"}, want: true},
		{name: "contain 3 counts with one char", args: args{str: "tact ca a"}, want: true},
		{name: "2 odds", args: args{str: "tact coa a"}, want: false},
		{name: "3 odds", args: args{str: "tact coa ai"}, want: false},
		{name: "one char", args: args{str: "a"}, want: true},
		{name: "empty", args: args{str: ""}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindromeOfPermutationWithBitVector(tt.args.str); got != tt.want {
				t.Errorf("IsPalindromeOfPermutationWithBitVector() = %v, want %v", got, tt.want)
			}
		})
	}
}
