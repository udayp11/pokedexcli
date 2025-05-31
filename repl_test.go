package main

import (
	"testing"
)

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
			input:    " who will win IPL 2024",
			expected: []string{"who", "will", "win", "ipl", "2025"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Length of slices dont match, actual len is '%v' whereas expected leng: '%v'", actual, c.expected)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words dont match, actual word is '%v' whereas expected word: '%v'", word, expectedWord)
			}
		}
	}
}
