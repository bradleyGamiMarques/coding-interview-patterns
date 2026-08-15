<h2> 167. Two Sum II - Input Array Is Sorted</h2>

![Medium](https://img.shields.io/badge/Medium-FAC31D)

<h2> Description</h2>
Given a <b>1-indexed</b> array of integers <code>numbers</code> that is already <i>sorted in non-decreasing order</i>, find two numbers such that they add up to a specific <code>target</code> number.

Let these two numbers be <code>numbers[index<sub>1</sub>]</code> and <code>numbers[index<sub>2</sub>]</code> where <code>1 <= index<sub>1</sub> < index<sub>2</sub> <= len(numbers).</code>

Return the indicies of the two numbers <code>index<sub>1</sub></code> and <code>index<sub>2</sub></code>, <b>each incremented by one</b>, as an integer array <code>[index<sub>1</sub>, index<sub>2</sub>]</code> of length 2.

The tests are generated such that there is <b>exactly one solution</b>. You <b>may not</b> use the same element twice.

Your solution must use only constant extra space.


<b>Example 1</b>:

<b>Input</b>: <i>numbers = [2, 7, 11, 15], target = 9</i>

<b>Output</b>: <i>[1, 2]</i>

<b>Explanation</b>: <i>The sum of 2 and 7 is 9. Therefore, index<sub>1</sub> = 1, index<sub>2</sub> = 2.</i>

We return [1, 2].
<hr/>

<b>Example 2</b>:

<b>Input</b>: <i>numbers = [2, 3, 4], target = 6</i>

<b>Output</b>: <i>[1, 3]</i>

<b>Explanation</b>: <i>The sum of 2 and 4 is 6. Therefore, index<sub>1</sub> = 1, index<sub>2</sub> = 3.</i>

We return [1, 3].
<hr/>

<b>Example 3</b>:

<b>Input</b>: <i>numbers = [-1, 0], target = -1</i>

<b>Output</b>: <i>[1, 2]</i>

<b>Explanation</b>: <i>The sum of -1 and 0 is -1. Therefore, index<sub>1</sub> = 1, index<sub>2</sub> = 2.</i>

We return [1, 2].
<hr/>

<b>Constraints</b>:

- <code>2 <= numbers.length <= 3 * 10<sup>4</sup></code>
- <code>-1000 <= numbers[i] <= 1000</code>
- <code>numbers</code> is sorted in <b>non-decreasing order</b>.
- <code>-1000 <= target <= 1000</code>
- The tests are generated such that there is <b>exactly one solution</b>.
