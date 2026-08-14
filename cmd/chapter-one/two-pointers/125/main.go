package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "A man, a plan, a canal: Panama"
	result := isPalindrome(input)
	fmt.Printf("is %v a palindrome? %t\n", input, result)
}

func isPalindrome(s string) bool {
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

func removeNonAlphanumeric(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, s)
}
