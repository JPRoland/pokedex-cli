package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  go   up  ",
			expected: []string{"go", "up"},
		},
		{
			input:    " ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected length %d, got %d", len(c.expected), len(actual))
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Expected %s, got %s", c.expected[i], actual[i])
			}
		}
	}
}
