# Golang-Deltas

### Arrays
- Find the maximum/minimum element in an array.
```go
func findMaxMinOfArray(arr []int) (max, min int) {
	max = 0
	min = arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	return max, min
}
```
- Find the index of a specific element in an array.
```go
func findIndex(a []int, inp int) int {
	for i := range a {
		if a[i] == inp {
			return i
		}
	}
	return -1
}
```
- Reverse an array.
```go
func reverseArray(arr []int) []int {
	l := 0
	r := len(arr) - 1
	for l < r {
		temp := arr[l]
		arr[l] = arr[r]
		arr[r] = temp
		l++
		r--
	}
	return arr
}
```
- Check if an array is sorted.
```go
func arraySortCheck(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			return false
		}
	}
	return true
}
```
- Find the second largest element in an array.
```go
func secondLargest(a []int) int {
	max := 0
	secondMax := 0
	for i := range a {
		if a[i] > max {
			max = a[i]

		} else {
			if a[i] > secondMax {
				secondMax = a[i]
			}
		}
	}
	return secondMax
}
```
- Remove duplicates from an array.
```go
func removeDup(arr []int) []int {
	l := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[l] {
			l++
			arr[l] = arr[i]
		}
	}
	return arr[:l+1]
}
func removeDuplicates(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == arr[i+1] {
			temp := arr[i+1:]
			arr = arr[:i]
			arr = append(arr, temp...)
		}
	}
	return arr
}
```
- Find the average of elements in an array.
```go
func average(a []int) int {
	sum := 0
	for i := range a {
		sum += a[i]
	}
	return sum / len(a)
}
```
- Move all zeroes to the end of the array while maintaining the order of non-zero elements.
```go
func moveZeroes(arr []int) []int {
	outArr := make([]int, len(arr))
	outIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			outArr[outIndex] = arr[i]
			outIndex++
		}
	}
	return outArr
}
```
- Rotate an array to the right by k steps.
```go
func rotateKStepsRight(a []int, k int) []int {
	part := a[:k]
	a = a[k:]
	a = append(a, part...)
	return a
}
```
- Two Sum
```go
func twoSum(arr []int, target int) (int, int) {
	for i := 1; i < len(arr); i++ { // Adjacent Sum check
		if arr[i-1]+arr[i] == target {
			return i, i + 1
		}
	}
	for r := len(arr) - 1; r > 0; r-- { // Check every sequences
		l := 0
		for l < r {
			if arr[l]+arr[r] == target {
				return l, r
			}
			l++
		}
	}
	return 0, 0
}
func twoSumOpt(arr []int, target int) (int, int) {
	hmap := make(map[int]int)
	for i := range arr {
		complement := target - arr[i]
		k, v := hmap[complement]
		if v {
			return i, k
		} else {
			hmap[arr[i]] = i
		}
	}
	return -1, -1
}
```
- Product of Array Except Self
- Maximum Subarray
- Find the Missing Number
### Strings
- Reverse a string.
- Check if a string is a palindrome.
- Count the number of vowels in a string.
- Find the length of the longest substring without repeating characters.
- Check if two strings are anagrams.
- Convert a string to lowercase/uppercase.
- Remove all spaces from a string.
- Find the first non-repeated character in a string.
- Find all permutations of a string.
- Check if a string contains all unique characters.
- Longest Common Prefix
- Valid Parentheses
- Anagrams
### Linked Lists
- Print all elements in a linked list.
- Find the length of a linked list.
- Insert a node at the beginning of a linked list.
- Insert a node at the end of a linked list.
- Delete a node from a linked list.
- Find a node by its value.
- Reverse a linked list.
- Check if a linked list is cyclic.
- Find the middle element of a linked list.
- Merge two sorted linked lists.
- Detect a Cycle in a Linked List
- Remove Nth Node From End of List
### Stacks Queues
- Implement a stack using an array.
- Implement a stack using a linked list.
- Check for balanced parentheses in an expression.
- Sort a stack using another stack.
- Find the minimum element in a stack.
- Perform stack operations (push, pop, peek).
- Evaluate a postfix expression using a stack.
- Evaluate Reverse Polish Notation
- Implement a queue using an array.
- Implement a queue using a linked list.
- Implement a circular queue.
- Implement a priority queue.
### Sorting and Searching
- Implement Bubble Sort.
- Implement Selection Sort.
- Implement Insertion Sort.
- Find an element in a sorted array using Binary Search.
- Count the number of occurrences of an element in a sorted array.
- Perform Linear Search on an unsorted array.
- Sort an array using Merge Sort.
- Sort an array using Quick Sort.
### DP
- Subset Sum
- Climbing Stairs
- Longest Common Subsequence
- Coin Change
- Longest Common Prefix