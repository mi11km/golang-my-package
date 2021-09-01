package chapter2

import (
	"reflect"
	"testing"
)

func TestSinglyLinkedList_Partition(t *testing.T) {
	type args struct {
		values []int
		x      int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"success", args{x: 5, values: []int{3, 5, 8, 5, 10, 2, 1}}, []int{1, 2, 3, 5, 8, 5, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSinglyLinkedList()
			for _, v := range tt.args.values {
				l.Insert(v)
			}
			l.Partition(tt.args.x)
			if got := l.Slice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v\n", got, tt.want)
			}
		})
	}
}

func TestSinglyLinkedList_Partition2(t *testing.T) {
	type args struct {
		values []int
		x      int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"success", args{x: 5, values: []int{3, 5, 8, 5, 10, 2, 1}}, []int{3, 2, 1, 5, 8, 5, 10}},
		{"nothing less than x", args{x: 0, values: []int{3, 5, 8, 5, 10, 2, 1}}, []int{3, 5, 8, 5, 10, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSinglyLinkedList()
			for _, v := range tt.args.values {
				l.Insert(v)
			}
			l.Partition2(tt.args.x)
			if got := l.Slice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v\n", got, tt.want)
			}
		})
	}
}

func TestSinglyLinkedList_SplitList(t *testing.T) {
	type args struct {
		values []int
		x      int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"success", args{x: 5, values: []int{3, 5, 8, 5, 10, 2, 1}}, []int{3, 2, 1, 5, 10, 5, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSinglyLinkedList()
			for _, v := range tt.args.values {
				l.Insert(v)
			}
			l.SplitList(tt.args.x)
			if got := l.Slice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v\n", got, tt.want)
			}
		})
	}
}

func TestSinglyLinkedList_Swap(t *testing.T) {
	type args struct {
		values []int
		i1, i2 int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"success", args{values: []int{3, 5, 8, 5, 10, 2, 1}, i1: 1, i2: 4}, []int{3, 10, 8, 5, 5, 2, 1}},
		{"index out of range", args{values: []int{3, 5, 8, 5, 10, 2, 1}, i1: 7, i2: 4}, []int{3, 5, 8, 5, 10, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSinglyLinkedList()
			for _, v := range tt.args.values {
				l.Insert(v)
			}
			l.Swap(tt.args.i1, tt.args.i2)
			if got := l.Slice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v\n", got, tt.want)
			}
		})
	}
}
