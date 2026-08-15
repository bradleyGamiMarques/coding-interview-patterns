package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "test example 1", input: "A man, a plan, a canal: Panama", want: true},
		{name: "test example 2", input: "race a car", want: false},
		{name: "test example 3", input: " ", want: true},
		{name: "test input string is all punctuation", input: "!@##@!", want: true},
		{name: "test input string is punctuation plus letter in middle", input: "!@#p#@!", want: true},
		{name: "test input string is punctuation plus leading letter", input: "p!@##@!", want: true},
		{name: "test input string is punctuation plus trailing leter", input: "!@##@!p", want: true},
		{name: "test input string is all numbers 0-9", input: "01234567899876543210", want: true},
		{name: "test input string is numbers plus punctuation", input: "!01234567899!876543210!", want: true},
		{name: "test input string is max input size", input: strings.Repeat("a", 200000), want: true},
		{name: "test input string is min input size", input: "a", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.input)
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})
	}
}

type testCase struct {
	name  string
	input byte
	want  bool
}

func generateIsAlumASCIITests() []testCase {
	tests := make([]testCase, 95)
	idx := 0

	for b := byte(' '); b <= '/'; b++ {
		tests[idx] = testCase{
			name:  fmt.Sprintf("non alphanumeric characters before digits:%s", string(b)),
			input: b,
			want:  false,
		}
		idx++
	}

	for b := byte('0'); b <= '9'; b++ {
		tests[idx] = testCase{
			name:  fmt.Sprintf("valid digit:%s", string(b)),
			input: b,
			want:  true,
		}
		idx++
	}

	for b := byte(':'); b <= '@'; b++ {
		tests[idx] = testCase{
			name:  fmt.Sprintf("non alphanumeric characters before uppercase letters:%s", string(b)),
			input: b,
			want:  false,
		}
		idx++
	}

	for b := byte('A'); b <= 'Z'; b++ {
		tests[idx] = testCase{
			name:  fmt.Sprintf("valid uppercase letter:%s", string(b)),
			input: b,
			want:  true,
		}
		idx++
	}

	for b := byte('['); b <= '`'; b++ {
		tests[idx] = testCase{
			name:  fmt.Sprintf("non alphanumeric characters before lowercase letters:%s", string(b)),
			input: b,
			want:  false,
		}
		idx++
	}

	for b := byte('a'); b <= 'z'; b++ {
		tests[idx] = testCase{
			name:  fmt.Sprintf("valid lowercase letter:%s", string(b)),
			input: b,
			want:  true,
		}
		idx++
	}

	for b := byte('{'); b <= '~'; b++ {
		tests[idx] = testCase{
			name:  fmt.Sprintf("non alphanumeric after lowercase letters:%s", string(b)),
			input: b,
			want:  false,
		}
		idx++
	}

	return tests
}
func TestIsAlumASCII(t *testing.T) {
	tests := generateIsAlumASCIITests()
	for i := range tests {
		tc := tests[i]
		t.Run(tc.name, func(t *testing.T) {
			got := isAlnumASCII(tc.input)
			if got != tc.want {
				t.Errorf("isAlnumASCII(%q) = %t, want %t", tc.input, got, tc.want)
			}
		})
	}
}
