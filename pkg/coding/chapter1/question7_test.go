package chapter1

import (
	"reflect"
	"testing"
)

func TestRotateMatrix90Degree(t *testing.T) {
	type args struct {
		matrix [][]rune
	}
	tests := []struct {
		name string
		args args
		want [][]rune
	}{
		{name: "success", args: args{
			matrix: [][]rune{
				{0, 1, 2, 3, 4},
				{0, 1, 2, 3, 4},
				{0, 1, 2, 3, 4},
				{0, 1, 2, 3, 4},
				{0, 1, 2, 3, 4},
			},
		}, want: [][]rune{
			{0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1},
			{2, 2, 2, 2, 2},
			{3, 3, 3, 3, 3},
			{4, 4, 4, 4, 4},
		}},
		{name: "success", args: args{
			matrix: [][]rune{
				{0, 1},
				{0, 1},
			},
		}, want: [][]rune{
			{0, 0},
			{1, 1},
		}},
		{name: "success", args: args{
			matrix: [][]rune{},
		}, want: [][]rune{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rotate90Degree(tt.args.matrix)
			if !reflect.DeepEqual(tt.args.matrix, tt.want) {
				t.Errorf("Rotate90Degree() 1 times = %v, want %v", tt.args.matrix, tt.want)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	type args struct {
		matrix [][]rune
	}
	tests := []struct {
		name string
		args args
		want [][]rune
	}{
		{name: "success", args: args{
			matrix: [][]rune{
				{0, 1, 2, 3, 4},
				{0, 1, 2, 3, 4},
				{0, 1, 2, 3, 4},
				{0, 1, 2, 3, 4},
				{0, 1, 2, 3, 4},
			},
		}, want: [][]rune{
			{0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1},
			{2, 2, 2, 2, 2},
			{3, 3, 3, 3, 3},
			{4, 4, 4, 4, 4},
		}},
		{name: "success", args: args{
			matrix: [][]rune{
				{0, 1},
				{0, 1},
			},
		}, want: [][]rune{
			{0, 0},
			{1, 1},
		}},
		{name: "success", args: args{
			matrix: [][]rune{},
		}, want: [][]rune{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Rotate(tt.args.matrix); !reflect.DeepEqual(tt.args.matrix, tt.want) {
				t.Errorf("Rotate() 1 times = %v, want %v", tt.args.matrix, tt.want)
			}
		})
	}
}
