package chapter2

import (
	"reflect"
	"testing"
)

func TestNewDoublyLinkedList(t *testing.T) {
	tests := []struct {
		name string
		want *DoublyLinkedList
	}{
		{name: "success", want: &DoublyLinkedList{nil, nil, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoublyLinkedList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDoublyLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_Insert_Delete_and_Get(t *testing.T) {
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
			ll := NewDoublyLinkedList()

			for _, val := range tt.args.insertVals {
				ll.Insert(val)
			}
			if ll.Len() != len(tt.wants.insert) {
				t.Errorf("want: %v, got: %v\n", len(tt.wants.insert), ll.Len())
			}
			if got := ll.Slice(); !reflect.DeepEqual(tt.wants.insert, got) {
				t.Errorf("want: %v, got: %v\n", tt.wants.insert, got)
			}

			ll.Delete(tt.args.deleteIndex)
			if ll.Len() != len(tt.wants.delete) {
				t.Errorf("want: %v, got: %v\n", len(tt.wants.delete), ll.Len())
			}
			if got := ll.Slice(); !reflect.DeepEqual(tt.wants.delete, got) {
				t.Errorf("want: %v, got: %v\n", tt.wants.delete, got)
			}

			if got := ll.Get(tt.args.getIndex); got != tt.wants.get {
				t.Errorf("want: %v, got: %v\n", tt.wants.get, got)
			}
		})
	}
}

// 双方向に辿れるかのテスト
func TestDoublyLinkedList_Slice(t *testing.T) {
	type args struct {
		vals []int
	}
	type wants struct {
		head []int
		tail []int
	}
	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			name:  "follow in both directions",
			args:  args{vals: []int{1, 2, 3, 4, 5}},
			wants: wants{head: []int{1, 2, 3, 4, 5}, tail: []int{5, 4, 3, 2, 1}},
		},
		{
			name:  "empty",
			args:  args{vals: []int{}},
			wants: wants{head: []int{}, tail: []int{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := NewDoublyLinkedList()
			for _, val := range tt.args.vals {
				ll.Insert(val)
			}

			if got := ll.Slice(); !reflect.DeepEqual(got, tt.wants.head) {
				t.Errorf("want: %v, got: %v\n", tt.wants.head, got)
			}
			if got := ll.SliceFromTail(); !reflect.DeepEqual(got, tt.wants.tail) {
				t.Errorf("want: %v, got: %v\n", tt.wants.tail, got)
			}
		})
	}
}

func TestDoublyLinkedList_String(t *testing.T) {
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
			ll := NewDoublyLinkedList()
			for _, val := range tt.args {
				ll.Insert(val)
			}

			if got := ll.String(); got != tt.wants {
				t.Errorf("want: %v, got: %v\n", tt.wants, got)
			}
		})
	}
}
