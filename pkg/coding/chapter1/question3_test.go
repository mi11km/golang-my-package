package chapter1

import "testing"

func TestURLifyArray(t *testing.T) {
	type args struct {
		runes  []rune
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "success", args: args{runes: []rune("Mr John Smith "), length: 13}, want: "Mr%20John%20Smith"},
		{name: "success", args: args{runes: []rune(" Mr John Smith "), length: 14}, want: "%20Mr%20John%20Smith"},
		{name: "success", args: args{runes: []rune(" Mr John Smith "), length: 15}, want: "%20Mr%20John%20Smith%20"},
		{name: "success", args: args{runes: []rune("Hello"), length: 5}, want: "Hello"},
		{name: "success", args: args{runes: []rune("  "), length: 2}, want: "%20%20"},
		{name: "success", args: args{runes: []rune("  "), length: 1}, want: "%20"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := URLifyArray(tt.args.runes, tt.args.length); got != tt.want {
				t.Errorf("URLifyWithStringsPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLifyWithStringsPackage(t *testing.T) {
	type args struct {
		str    string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "success", args: args{str: "Mr John Smith ", length: 13}, want: "Mr%20John%20Smith"},
		{name: "success", args: args{str: " Mr John Smith ", length: 14}, want: "%20Mr%20John%20Smith"},
		{name: "success", args: args{str: " Mr John Smith ", length: 15}, want: "%20Mr%20John%20Smith%20"},
		{name: "success", args: args{str: "Hello", length: 5}, want: "Hello"},
		{name: "success", args: args{str: "  ", length: 2}, want: "%20%20"},
		{name: "success", args: args{str: "  ", length: 1}, want: "%20"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := URLifyWithStringsPackage(tt.args.str, tt.args.length); got != tt.want {
				t.Errorf("URLifyWithStringsPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLify(t *testing.T) {
	type args struct {
		str    string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "success", args: args{str: "Mr John Smith ", length: 13}, want: "Mr%20John%20Smith"},
		{name: "success", args: args{str: " Mr John Smith ", length: 14}, want: "%20Mr%20John%20Smith"},
		{name: "success", args: args{str: " Mr John Smith ", length: 15}, want: "%20Mr%20John%20Smith%20"},
		{name: "success", args: args{str: "Hello", length: 5}, want: "Hello"},
		{name: "success", args: args{str: "  ", length: 2}, want: "%20%20"},
		{name: "success", args: args{str: "  ", length: 1}, want: "%20"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := URLify(tt.args.str, tt.args.length); got != tt.want {
				t.Errorf("Urlify() = %v, want %v", got, tt.want)
			}
		})
	}
}
