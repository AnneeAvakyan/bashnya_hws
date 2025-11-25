package main

import (
	"testing"
)

func TestUniq(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		options  Options
		expected []string
	}{

		{
			name:     "Empty input",
			input:    []string{},
			options:  Options{},
			expected: []string{},
		},
		{
			name:     "Single line",
			input:    []string{"hello"},
			options:  Options{},
			expected: []string{"hello"},
		},
		{
			name:     "Duplicate lines",
			input:    []string{"hello", "hello", "world"},
			options:  Options{},
			expected: []string{"hello", "world"},
		},
		{
			name:     "Multiple duplicates",
			input:    []string{"a", "b", "a", "c", "b", "a"},
			options:  Options{},
			expected: []string{"a", "b", "c"},
		},

		{
			name:  "Count mode basic",
			input: []string{"a", "b", "a", "c", "b", "a"},
			options: Options{
				CountMode: true,
			},
			expected: []string{"3 a", "2 b", "1 c"},
		},
		{
			name:  "Count mode single line",
			input: []string{"hello"},
			options: Options{
				CountMode: true,
			},
			expected: []string{"1 hello"},
		},

		{
			name:  "Repeated mode basic",
			input: []string{"a", "b", "a", "c", "b", "a"},
			options: Options{
				RepeatedMode: true,
			},
			expected: []string{"a", "b"},
		},
		{
			name:  "Repeated mode no duplicates",
			input: []string{"a", "b", "c"},
			options: Options{
				RepeatedMode: true,
			},
			expected: []string{},
		},

		{
			name:  "Unique mode basic",
			input: []string{"a", "b", "a", "c", "b", "a"},
			options: Options{
				UniqueMode: true,
			},
			expected: []string{"c"},
		},
		{
			name:  "Unique mode all duplicates",
			input: []string{"a", "a", "a"},
			options: Options{
				UniqueMode: true,
			},
			expected: []string{},
		},

		{
			name:  "Ignore case basic",
			input: []string{"Hello", "hello", "WORLD", "world", "Test"},
			options: Options{
				IgnoreCase: true,
			},
			expected: []string{"Hello", "WORLD", "Test"},
		},
		{
			name:  "Ignore case with count",
			input: []string{"Hello", "hello", "WORLD", "world", "Test"},
			options: Options{
				IgnoreCase: true,
				CountMode:  true,
			},
			expected: []string{"2 Hello", "2 WORLD", "1 Test"},
		},

		{
			name:  "Num fields basic",
			input: []string{"We love music.", "I love music.", "They love music.", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
			options: Options{
				NumFields: 1,
			},
			expected: []string{"We love music.", "I love music of Kartik.", "Thanks."},
		},
		{
			name:  "Num fields multiple spaces",
			input: []string{"We  love  music.", "I  love  music.", "They  love  music.", "I  love  music  of  Kartik.", "We  love  music  of  Kartik.", "Thanks."},
			options: Options{
				NumFields: 2,
			},
			expected: []string{"We  love  music.", "I  love  music  of  Kartik.", "Thanks."},
		},
		{
			name:  "Num fields more than available",
			input: []string{"short", "long string here"},
			options: Options{
				NumFields: 3,
			},
			expected: []string{"short"},
		},

		{
			name:  "Num chars basic",
			input: []string{"defabc", "xyzabc", "ghiabcdef"},
			options: Options{
				NumChars: 3,
			},
			expected: []string{"defabc", "ghiabcdef"},
		},
		{
			name:  "Num chars more than string length",
			input: []string{"abc", "def"},
			options: Options{
				NumChars: 5,
			},
			expected: []string{"abc", "def"},
		},

		{
			name:  "Fields and chars combined",
			input: []string{"field1 field2 abcdef", "field1 field2 abcxyz"},
			options: Options{
				NumFields: 2,
				NumChars:  3,
			},
			expected: []string{"field1 field2 abcdef", "field1 field2 abcxyz"},
		},
		{
			name:  "Ignore case with repeated mode",
			input: []string{"Hello", "hello", "World", "test"},
			options: Options{
				IgnoreCase:   true,
				RepeatedMode: true,
			},
			expected: []string{"Hello"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Uniq(tt.input, tt.options)

			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d lines, got %d", len(tt.expected), len(result))
				return
			}

			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Line %d: expected '%s', got '%s'", i, tt.expected[i], result[i])
				}
			}
		})
	}
}
