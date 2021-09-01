package chapter1

import (
	"testing"
)

func TestIsDuplicated(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty", args: args{str: ""}, want: true},
		{name: "empty char", args: args{str: " "}, want: true},
		{name: "empty chars", args: args{str: "  "}, want: false},
		{name: "ja true", args: args{str: "こんにちは"}, want: true},
		{name: "ja false", args: args{str: "ぺんぺん"}, want: false},
		{name: "en false", args: args{str: "air"}, want: true},
		{name: "en false", args: args{str: "Hello"}, want: false},
		{name: "mark true", args: args{str: "-^¥[:/.,"}, want: true},
		{name: "mark false", args: args{str: "-^¥[:-/.,"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDuplicated(tt.args.str); got != tt.want {
				t.Errorf("IsDuplicated() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUniqueChars(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty", args: args{str: ""}, want: true},
		{name: "empty char", args: args{str: " "}, want: true},
		{name: "empty chars", args: args{str: "  "}, want: false},
		{name: "en false", args: args{str: "air"}, want: true},
		{name: "en false", args: args{str: "Hello"}, want: false},
		{name: "mark true", args: args{str: "!#$"}, want: true},
		{name: "mark false", args: args{str: "!#$%%"}, want: false},
		{name: "over 128 words", args: args{str: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
			"aaaaaaaaaaaaaaaa"}, want: false},
		{name: "ASCII words", args: args{
			str: "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOP" +
				"QRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~\u007F",
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUniqueChars(tt.args.str); got != tt.want {
				t.Errorf("IsUniqueChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
