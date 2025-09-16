package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " Hello World ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   Go   Lang   ",
			expected: []string{"go", "lang"},
		},
		{
			input:    "PYTHON   rocks  HARD ",
			expected: []string{"python", "rocks", "hard"},
		},
		{
			input:    "   multiple   spaces   inside  ",
			expected: []string{"multiple", "spaces", "inside"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Len of actual: %v doesn't match len of expected: %v", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Failed test case due to words missmatching")
			}
		}
	}
}
