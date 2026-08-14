<h2> 125. Valid Palindrome</h2>

![Easy](https://img.shields.io/badge/Easy-46C6C2)

<h2> Description</h2>
A phrase is a <b>palindrome</b> if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string <code>s</code>, return <code>true</code> if it is a <b>palindrome</b>, or <code>false</code> otherwise.


<b>Example 1</b>:

<b>Input</b>: <i>s = "A man, a plan, a canal: Panama"</i>

<b>Output</b>: <i>true</i>

<b>Explanation</b>: <i>"amanaplanacanalpanama"</i> is a palindrome.
<hr/>

<b>Example 2</b>:

<b>Input</b>: <i>s = "race a car"</i>

<b>Output</b>: <i>false</i>

<b>Explanation</b>: <i>"raceacar" is not a palindrome.</i>
<hr/>

<b>Example 3</b>:

<b>Input</b>: <i>s = " "</i>

<b>Output</b>: <i>true</i>

<b>Explanation</b>: <i>s is an empty string "" after removing non-alphanumeric characters. Since an empty string reads the same forward and backward, it is a palindrome.</i>
<hr/>

<b>Constraints</b>:


- <code>1 <= len(s) <= 2 * 10<sup>5</sup></code>
- <code>s</code> consists only of printable ASCII characters.

