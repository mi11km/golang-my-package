package greetings

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name     string
		arg      string
		expected string
	}{
		{
			name:     "default",
			arg:      "Mike",
			expected: "Hi, Mike. Welcome!",
		},
		{
			name:     "empty string",
			arg:      "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := Hello(tt.arg); got != tt.expected || err == nil {
				t.Errorf("got: %s, want %s", got, tt.expected)
			}
		})
	}
}
