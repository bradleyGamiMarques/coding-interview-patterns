package main

import (
	"fmt"
	"strings"
	"testing"
)

type palindromeTestCase struct {
	name  string
	input string
	want  bool
}

type isAlumTestCase struct {
	name  string
	input byte
	want  bool
}

type removeNonAlphanumericTestCase struct {
	name  string
	input string
	want  string
}

func generatePalindromeTests() []palindromeTestCase {
	return []palindromeTestCase{
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
}
func TestIsPalindrome(t *testing.T) {
	tests := generatePalindromeTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.input)
			if got != tt.want {
				t.Errorf("isPalindrome(%q) = %t, want %t", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsValidPalindrome(t *testing.T) {
	tests := generatePalindromeTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.input)
			if got != tt.want {
				t.Errorf("isPalindrome(%q) = %t, want %t", tt.input, got, tt.want)
			}
		})
	}

}

func generateIsAlumASCIITests() []isAlumTestCase {
	tests := make([]isAlumTestCase, 95)
	idx := 0

	for b := byte(' '); b <= '/'; b++ {
		tests[idx] = isAlumTestCase{
			name:  fmt.Sprintf("non alphanumeric characters before digits:%s", string(b)),
			input: b,
			want:  false,
		}
		idx++
	}

	for b := byte('0'); b <= '9'; b++ {
		tests[idx] = isAlumTestCase{
			name:  fmt.Sprintf("valid digit:%s", string(b)),
			input: b,
			want:  true,
		}
		idx++
	}

	for b := byte(':'); b <= '@'; b++ {
		tests[idx] = isAlumTestCase{
			name:  fmt.Sprintf("non alphanumeric characters before uppercase letters:%s", string(b)),
			input: b,
			want:  false,
		}
		idx++
	}

	for b := byte('A'); b <= 'Z'; b++ {
		tests[idx] = isAlumTestCase{
			name:  fmt.Sprintf("valid uppercase letter:%s", string(b)),
			input: b,
			want:  true,
		}
		idx++
	}

	for b := byte('['); b <= '`'; b++ {
		tests[idx] = isAlumTestCase{
			name:  fmt.Sprintf("non alphanumeric characters before lowercase letters:%s", string(b)),
			input: b,
			want:  false,
		}
		idx++
	}

	for b := byte('a'); b <= 'z'; b++ {
		tests[idx] = isAlumTestCase{
			name:  fmt.Sprintf("valid lowercase letter:%s", string(b)),
			input: b,
			want:  true,
		}
		idx++
	}

	for b := byte('{'); b <= '~'; b++ {
		tests[idx] = isAlumTestCase{
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

func generateRemoveNonAlphanumericTests() []removeNonAlphanumericTestCase {
	return []removeNonAlphanumericTestCase{
		{name: "test only lowercase letters", input: "hello", want: "hello"},
		{name: "test only uppercase letters", input: "WORLD", want: "WORLD"},
		{name: "test only numbers", input: "1234567890", want: "1234567890"},
		{name: "test mixed alphanumeric", input: "GoLang123", want: "GoLang123"},
		{name: "test spaces and punctuation", input: "hello, world!!!", want: "helloworld"},
		{name: "test special characters and symbols", input: "a#b$c%1^2&3", want: "abc123"},
		{name: "test empty string", input: "", want: ""},
		{name: "test only non-alphanumeric", input: "!@#$%&*()_+", want: ""},
	}
}
func TestRemoveNonAlphanumeric(t *testing.T) {
	tests := generateRemoveNonAlphanumericTests()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeNonAlphanumeric(tt.input)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
