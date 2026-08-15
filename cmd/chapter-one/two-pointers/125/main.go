package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "A man, a plan, a canal: Panama"
	result := isPalindrome(input)
	fmt.Printf("isPalindrome: is %s a palindrome? %t\n", input, result)

	result = isValidPalindrome(input)
	fmt.Printf("isValidPalindrome: is %s a palindrome? %t\n", input, result)
}

// isPalindrome returns a bool representing whether the input string s is a palindrome.
// It operates in O(n) time and with O(1) space.
// A palindrome is a phrase if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward.
// It assumes that the input string s consists only of printable ASCII characters.
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if !isAlnumASCII(s[left]) {
			left++
			continue
		}

		if !isAlnumASCII(s[right]) {
			right--
			continue
		}

		if s[left]|32 != s[right]|32 {
			return false
		}

		left++
		right--
	}
	return true
}

// isAlnumASCII returns a bool representing whether the input byte is a printable ASCII character.
// It does not handle extended ASCII values.
func isAlnumASCII(b byte) bool {
	if b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= '0' && b <= '9' {
		return true
	}
	return false
}

// isValidPalindrome returns a bool representing whether the input string s is a palindrome.
// It operates in O(n) time and with O(n) space.
// The extra space is due to the use of the removeNonAlphanumeric helper func.
// A palindrome is a phrase if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward.
// It assumes that the input string s consists only of printable ASCII characters.
func isValidPalindrome(s string) bool {
	sSanitized := removeNonAlphanumeric(s)
	left, right := 0, len(sSanitized)-1
	for left < right {
		if sSanitized[left]|32 != sSanitized[right]|32 {
			return false
		}
		left++
		right--
	}
	return true
}

// removeNonAlphanumeric returns a copy of the input string with non-alphanumeric characters removed.
// It operates in O(n) time and with O(n) space.
// It assumes that the input string s consists only of printable ASCII characters.
func removeNonAlphanumeric(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, s)
}
