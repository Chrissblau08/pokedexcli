package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "   lots   of   spaces   ",
			expected: []string{"lots", "of", "spaces"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// 1. Länge vergleichen
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) returned %d words, expected %d. Got: %v",
				c.input, len(actual), len(c.expected), actual)
			continue
		}

		// 2. Inhalt vergleichen
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%q) word[%d] = %q, expected %q",
					c.input, i, actual[i], c.expected[i])
			}
		}
	}
}
