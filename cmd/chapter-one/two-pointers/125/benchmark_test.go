package main

import (
	"fmt"
	"strings"
	"testing"
)

var ResultSink bool

func BenchmarkPalindromes(b *testing.B) {

	scenarios := []struct {
		name  string
		input string
	}{
		{name: "test input string is max input size", input: strings.Repeat("a", 200000)},
		{name: "test input string is max input size all non-alphanumeric", input: strings.Repeat("!@#$%^&*()_+{}|[]:;'<>?,./`~-=\"", 6250)},
	}

	for _, tc := range scenarios {

		b.Run(fmt.Sprintf("isPalindrome/%s", tc.name), func(b *testing.B) {
			for b.Loop() {
				ResultSink = isPalindrome(tc.input)
			}
		})

		b.Run(fmt.Sprintf("isValidPalindrome/%s", tc.name), func(b *testing.B) {
			for b.Loop() {
				ResultSink = isValidPalindrome(tc.input)
			}
		})
	}
}
