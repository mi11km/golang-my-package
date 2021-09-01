package chapter2

import (
	"reflect"
	"testing"
)

func TestNewSinglyLinkedList(t *testing.T) {
	tests := []struct {
		name string
		want *SinglyLinkedList
	}{
		{name: "success", want: &SinglyLinkedList{nil, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSinglyLinkedList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSinglyLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSinglyLinkedList_Insert_Delete_and_Get(t *testing.T) {
	type args struct {
		insertVals  []int
		deleteIndex int
		getIndex    int
	}
	type wants struct {
		insert []int
		delete []int
		get    int
	}
	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			name:  "success insert and delete",
			args:  args{insertVals: []int{1, 2, 3, 4}, deleteIndex: 2, getIndex: 0},
			wants: wants{insert: []int{1, 2, 3, 4}, delete: []int{1, 2, 4}, get: 1},
		},
		{
			name:  "insert 1 val, delete 1 val (from head), get from non-exist-value index",
			args:  args{insertVals: []int{1}, deleteIndex: 0, getIndex: 0},
			wants: wants{insert: []int{1}, delete: []int{}, get: -1},
		},
		{
			name:  "insert 3 val, delete 1 val (from head)",
			args:  args{insertVals: []int{1, 2, 3}, deleteIndex: 0, getIndex: 0},
			wants: wants{insert: []int{1, 2, 3}, delete: []int{2, 3}, get: 2},
		},
		{
			name:  "delete 1 val (from tail)",
			args:  args{insertVals: []int{1, 2, 3}, deleteIndex: 2, getIndex: 1},
			wants: wants{insert: []int{1, 2, 3}, delete: []int{1, 2}, get: 2},
		},
		{
			name:  "delete index out of range",
			args:  args{insertVals: []int{1, 2, 3}, deleteIndex: 5, getIndex: 1},
			wants: wants{insert: []int{1, 2, 3}, delete: []int{1, 2, 3}, get: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSinglyLinkedList()

			for _, val := range tt.args.insertVals {
				l.Insert(val)
			}
			if l.Len() != len(tt.wants.insert) {
				t.Errorf("want: %v, got: %v\n", len(tt.wants.insert), l.Len())
			}
			if got := l.Slice(); !reflect.DeepEqual(tt.wants.insert, got) {
				t.Errorf("want: %v, got: %v\n", tt.wants.insert, got)
			}

			l.Delete(tt.args.deleteIndex)
			if l.Len() != len(tt.wants.delete) {
				t.Errorf("want: %v, got: %v\n", len(tt.wants.delete), l.Len())
			}
			if got := l.Slice(); !reflect.DeepEqual(tt.wants.delete, got) {
				t.Errorf("want: %v, got: %v\n", tt.wants.delete, got)
			}

			if got := l.Get(tt.args.getIndex); got != tt.wants.get {
				t.Errorf("want: %v, got: %v\n", tt.wants.get, got)
			}
		})
	}
}

// 辿れるかのテスト
func TestSinglyLinkedList_Slice(t *testing.T) {
	type args struct {
		vals []int
	}
	type wants struct {
		head []int
	}
	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			name:  "follow in both directions",
			args:  args{vals: []int{1, 2, 3, 4, 5}},
			wants: wants{head: []int{1, 2, 3, 4, 5}},
		},
		{
			name:  "empty",
			args:  args{vals: []int{}},
			wants: wants{head: []int{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSinglyLinkedList()
			for _, val := range tt.args.vals {
				l.Insert(val)
			}

			if got := l.Slice(); !reflect.DeepEqual(got, tt.wants.head) {
				t.Errorf("want: %v, got: %v\n", tt.wants.head, got)
			}
		})
	}
}

func TestSinglyLinkedList_String(t *testing.T) {
	tests := []struct {
		name  string
		args  []int
		wants string
	}{
		{name: "success", args: []int{1, 2, 3, 4, 5}, wants: "[1 2 3 4 5]"},
		{name: "empty", args: []int{}, wants: "[]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewSinglyLinkedList()
			for _, val := range tt.args {
				l.Insert(val)
			}

			if got := l.String(); got != tt.wants {
				t.Errorf("want: %v, got: %v\n", tt.wants, got)
			}
		})
	}
}
