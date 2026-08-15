# Intuition

A string is a palindrome that, when read left to right or right to left, remains identical. In other words, if the string were to be reversed, it should read the same, ignoring spaces and punctuation.

A palindrome has the property that the first character should be the same as the last character, the second character should be the same as the second-to-last character, etc.

<b>Example</b>:

Forward:   d → a → t → a → t → a → d

Backward:  d ← a ← t ← a ← t ← a ← d

Forward:   r → a → c → e → c → a → r

Backward:  r ← a ← c ← e ← c ← a ← r

Palindromes with an odd length, such as racecar, are different since they have a middle character. In this case it is safe to ignore the middle character since there is no character that is symmetrical.


These properties make palindromes the ideal candidate for the two pointer technique.

# Implementation
What is a pointer? A pointer is a variable to points to an index or value in a data structure.

Let us consider two pointers, left and right. Start by initializing left to the index that marks the beginning of the string and initialize right to the index of the last character in the string. Our end goal is to compare the characters at each index, and if they are not the same, we know that the input string is not a palindrome. To do this we move the pointers toward each other in a loop. The loop terminates when the value of the left pointer is no longer less than the value of the right pointer.

## Processing non-alphanumeric characters
Based on the problem constraints we know that the input string will consist of ASCII printable characters.
Consider the following helper function:
```go
func isAlnumASCII(b byte) bool {
	if b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= '0' && b <= '9' {
		return true
	}
	return false
}
```
When iterating over a string in Go <code>s[index]</code> will return a byte. We pass this byte to the helper function and it will tell us if the character is alphanumeric. 

<b>Note</b>: This helper function will not work on Unicode characters. It is an exercise to the reader to determine how that might be handled.

Since non-alphanumeric characters do not affect whether a string is a palindrome we can skip them while iterating over the input string.
This can be accomplished by incrementing the left pointer and decrementing the right pointer respectively e.g.
```go
left++
right--
```
## Handling capitalization
Let's review a part of the requirements, "A string is a <b>palindrome if, after converting all uppercase letters into lowercase letters...</b>"

A simple approach would be to use <code>strings.ToLower</code> function on the input string. However, this has the tradeoff of making our function not pure, and would allocate additional memory. What if the interviewer said that we could not use the standard library function to convert the input string to lowercase? How would you respond?

We can take advantage of a property of ASCII encoding to solve this problem. In ASCII uppercase and lowercase letters differ by exactly one bit: bit 5(value 32).

<code>A = 01000001
a = 01100001
</code>

When you do:

<code>a|32</code> you are telling the computer convert this byte into binary and update the 5th bit. If it is a 0 change it to 1 else leave it alone.
This solution is fast and efficient since it is operating at the binary level.
