package chapter2

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want bool
	}{
		{name: "success return true", args: []int{1, 2, 3, 4, 3, 2, 1}, want: true},
		{name: "success return false", args: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: false},
		{name: "empty list", args: []int{}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewSinglyLinkedList()
			for _, v := range tt.args {
				list.Insert(v)
			}

			if got := IsPalindrome(list); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
