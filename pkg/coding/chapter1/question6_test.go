package chapter1

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCompressStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "empty", args: args{str: ""}, want: ""},
		{name: "success", args: args{str: "aabcccccaaa"}, want: "a2b1c5a3"},
		{name: "not duplicate", args: args{str: "abcdefg"}, want: "abcdefg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompressStr(tt.args.str); got != tt.want {
				t.Errorf("CompressStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompressStrWithStringsBuilder(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "empty", args: args{str: ""}, want: ""},
		{name: "success", args: args{str: "aabcccccaaa"}, want: "a2b1c5a3"},
		{name: "not duplicate", args: args{str: "abcdefg"}, want: "abcdefg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompressStrWithStringsBuilder(tt.args.str); got != tt.want {
				t.Errorf("CompressStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompressStrWithStringsBuilder2(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "empty", args: args{str: ""}, want: ""},
		{name: "success", args: args{str: "aabcccccaaa"}, want: "a2b1c5a3"},
		{name: "not duplicate", args: args{str: "abcdefg"}, want: "abcdefg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompressStrWithStringsBuilder2(tt.args.str); got != tt.want {
				t.Errorf("CompressStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

var content string

func TestMain(m *testing.M) {
	data, _ := ioutil.ReadFile("mock.txt")
	content = string(data)
	status := m.Run()
	os.Exit(status)
}

func BenchmarkCompressStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompressStr(content)
	}
}

func BenchmarkCompressStrWithStringsBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompressStrWithStringsBuilder(content)
	}
}

func BenchmarkCompressStrWithStringsBuilder2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompressStrWithStringsBuilder2(content)
	}
}
