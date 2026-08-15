package main

import (
	"fmt"
)

func main() {
	input := "A man, a plan, a canal: Panama"
	result := isPalindrome(input)
	fmt.Printf("is %v a palindrome? %t\n", input, result)
}

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

func isAlnumASCII(b byte) bool {
	if b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= '0' && b <= '9' {
		return true
	}
	return false
}
