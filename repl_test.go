package main

import (
	"reflect"
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
			input:    " GO  is  AWESOME ",
			expected: []string{"go", "is", "awesome"},
		},
		{
			input:    " singleWord ",
			expected: []string{"singleword"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("For input '%s', expected %v, got %v", c.input, c.expected, actual)
		}
	}
}
