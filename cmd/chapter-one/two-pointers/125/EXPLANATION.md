# Intuition

A string is a palindrome that when read left to right or right to left remains identical. In other words, if the string were to be reversed it should read the same, ignoring spaces and punctuation.

A palindrome has the property where first character should be the same as the last character, the second character would be the same as the second to last character, etc.

<b>Example</b>:

Forward:   d → a → t → a → t → a → d

Backward:  d ← a ← t ← a ← t ← a ← d

Forward:   r → a → c → e → c → a → r

Backward:  r ← a ← c ← e ← c ← a ← r

Palindromes with an odd length, such as racecar, are different since they have a middle character. In this case it is safe to ignore the middle character since there is no character that is symmetrical.


These properties make palindromes the ideal candidate for the two pointer technique.

# Implementation
What is a pointer? A pointer is a variable to points to an index or value in a data structure.

Let us consider two pointers left and right. Start by initializing left to the index that marks the beginningof the string and initialize right to the index of the last character in the string. 
