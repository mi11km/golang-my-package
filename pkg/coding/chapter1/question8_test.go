package chapter1

import (
	"reflect"
	"testing"
)

func TestToZero(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "no zero in row[0] and col[0]", args: args{[][]int{
			{1, 2, 3, 4},
			{2, 0, 2, 3},
			{4, 3, 0, 1},
			{3, 2, 1, 1},
		}}, want: [][]int{
			{1, 0, 0, 4},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{3, 0, 0, 1},
		}},
		{name: "zero in col[0]", args: args{[][]int{
			{1, 2, 3, 4},
			{0, 1, 2, 3},
			{4, 3, 2, 1},
			{3, 2, 1, 0},
		}}, want: [][]int{
			{0, 2, 3, 0},
			{0, 0, 0, 0},
			{0, 3, 2, 0},
			{0, 0, 0, 0},
		}},
		{name: "zero in row[0]", args: args{[][]int{
			{1, 0, 3, 4},
			{1, 1, 2, 3},
			{4, 3, 2, 1},
			{3, 2, 1, 0},
		}}, want: [][]int{
			{0, 0, 0, 0},
			{1, 0, 2, 0},
			{4, 0, 2, 0},
			{0, 0, 0, 0},
		}},
		{name: "matrix length 0*0", args: args{[][]int{}}, want: [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ToZero(tt.args.matrix); !reflect.DeepEqual(tt.args.matrix, tt.want) {
				t.Errorf("ToZero() = %v, want %v", tt.args.matrix, tt.want)
			}
		})
	}
}

func TestSetZeros(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "no zero in row[0] and col[0]", args: args{[][]int{
			{1, 2, 3, 4},
			{2, 0, 2, 3},
			{4, 3, 0, 1},
			{3, 2, 1, 1},
		}}, want: [][]int{
			{1, 0, 0, 4},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{3, 0, 0, 1},
		}},
		{name: "zero in col[0]", args: args{[][]int{
			{1, 2, 3, 4},
			{0, 1, 2, 3},
			{4, 3, 2, 1},
			{3, 2, 1, 0},
		}}, want: [][]int{
			{0, 2, 3, 0},
			{0, 0, 0, 0},
			{0, 3, 2, 0},
			{0, 0, 0, 0},
		}},
		{name: "zero in row[0]", args: args{[][]int{
			{1, 0, 3, 4},
			{1, 1, 2, 3},
			{4, 3, 2, 1},
			{3, 2, 1, 0},
		}}, want: [][]int{
			{0, 0, 0, 0},
			{1, 0, 2, 0},
			{4, 0, 2, 0},
			{0, 0, 0, 0},
		}},
		{name: "matrix length 0*0", args: args{[][]int{}}, want: [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if SetZeros(tt.args.matrix); !reflect.DeepEqual(tt.args.matrix, tt.want) {
				t.Errorf("SetZeros() = %v, want %v", tt.args.matrix, tt.want)
			}
		})
	}
}
