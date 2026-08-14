package main

import (
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

func TestRemoveNonalphanumeric(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "test only lowercase letters", input: "hello", want: "hello"},
		{name: "test only uppercase letters", input: "WORLD", want: "WORLD"},
		{name: "test only numbers", input: "1234567890", want: "1234567890"},
		{name: "test mixed alphanumeric", input: "GoLang123", want: "GoLang123"},
		{name: "test spaces and punctuation", input: "hello, world!!!", want: "helloworld"},
		{name: "test special characters and symbols", input: "a#b$c%1^2&3", want: "abc123"},
		{name: "test empty string", input: "", want: ""},
		{name: "test only non-alphanumeric", input: "!@#$%&*()_+", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeNonAlphanumeric(tt.input)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
