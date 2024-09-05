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
```go
func productOfArrayExceptSelf(a []int) []int {
	var outArr []int
	for i := 0; i < len(a); i++ {
		prod := 1
		for j := 0; j < len(a); j++ {
			if j == i {
				continue
			}
			prod = prod * a[j]
		}
		outArr = append(outArr, prod)
	}
	return outArr
}
```
- Find the Missing Number
```go
func missingNumber(a []int) int {
	length := len(a) + 1
	actualSum := length * (length + 1) / 2
	sum := 0
	for i := range a {
		sum += a[i]
	}
	missingNum := actualSum - sum
	return missingNum
}
```
- Maximum Subarray

### Strings
- Reverse a string.
```go
func reverseString(s string) string {
	arr := []rune(s)
	l := 0
	r := len(arr) - 1
	for l < r {
		t := arr[l]
		arr[l] = arr[r]
		arr[r] = t
		l++
		r--
	}
	return string(arr)
}
```
- Check if a string is a palindrome.
```go
func palindrome(s string) bool {
	arr := []rune(s)
	l := 0
	r := len(arr) - 1
	for l < r {
		if arr[l] != arr[r] {
			return false
		}
		l++
		r--
	}
	return true
}
```
- Count the number of vowels in a string.
```go
func noOfVowels(s string) int {
	arr := []rune(strings.ToLower(s))
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == rune('a') || arr[i] == rune('e') || arr[i] == rune('i') || arr[i] == rune('o') || arr[i] == rune('u') {
			count++
		}
	}
	return count
}
```
- Find the length of the longest substring without repeating characters.
```go
func longestSubstringWithoutRepeat(s string) int {
	arr := []rune(s)
	uni := make(map[rune]int)
	count := 0
	for i := 0; i < len(arr); i++ {
		_, ok := uni[arr[i]]
		if ok {
			return count
		}
		uni[arr[i]] = 0
		count++
	}
	return count
}
```
- Check if two strings are anagrams.
```go
func anagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	arr1 := []rune(s1)
	arr2 := []rune(s2)
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)

	for i := 0; i < len(arr1); i++ {
		v, ok := m1[arr1[i]]
		if ok {

			m1[arr1[i]] = v + 1
		} else {
			m1[arr1[i]] = 0
		}
		v, ok = m2[arr2[i]]
		if ok {

			m2[arr2[i]] = v + 1
		} else {
			m2[arr2[i]] = 0
		}
	}
	for k, v1 := range m1 {
		v2, ok := m2[k]
		if ok && v1 != v2 {
			return false
		}
	}
	return true
}
```
- Find the first non-repeated character in a string.
```go
func nonRepeatingFirstChar(s string) string {
	arr := []rune(s)
	m := make(map[rune]int)
	for i := 0; i < len(arr); i++ {
		v, ok := m[arr[i]]
		if ok {
			m[arr[i]] = v + 1
		} else {
			m[arr[i]] = 0
		}
	}
	for _, v := range s {
		if m[v] == 0 {
			return string(v)
		}
	}
	return ""
}
```
- Find all permutations of a string.
- Check if a string contains all unique characters.
```go
func checkUnique(s string) bool {
	arr := []rune(s)
	m := make(map[rune]int)
	for i := range arr {
		v, ok := m[arr[i]]
		if ok {
			m[arr[i]] = v + 1
			return false
		} else {
			m[arr[i]] = 0
		}
	}
	return true
}
```
- Longest Common Prefix

### Linked Lists
- Print all elements in a linked list.
```go
type Node struct {
	data  int
	right *Node
}

type list struct {
	Nodes *Node
}

func NewList() *list {
	return &list{}
}
func (l *list) Add(d int) {
	node := Node{
		data: d,
	}
	if l.Nodes == nil {
		l.Nodes = &node
		return
	}
	c := l.Nodes
	for c.right != nil {
		c = c.right
	}
	c.right = &node
}
func (l *list) Remove(n int) {
	if l.Nodes == nil {
		return
	}
	c := l.Nodes
	for i := 0; i < n-1; i++ {
		if c.right != nil {
			c = c.right
		}
	}
	if c.right.right != nil {
		c.right = c.right.right
	}
}
func (l *list) print() {
	c := l.Nodes
	for c.right != nil {
		fmt.Println(c.data)
		c = c.right
	}
	fmt.Println(c.data)
	fmt.Println("------------------------")
}
func (l *list) NthAdd(d int, n int) {
	node := Node{
		data: d,
	}
	if l.Nodes == nil {
		l.Nodes = &node
	}
	if n == 0 {
		node.right = l.Nodes
		l.Nodes = &node
	}

	c := l.Nodes
	for i := 0; i < n-1; i++ {
		if c.right != nil {
			c = c.right
		}
	}
	node.right = c.right
	c.right = &node
}
func main() {
	l := NewList()
	l.Add(2)
	l.Add(4)
	l.Add(6)
	l.Add(9)
	l.Add(12)
	l.Add(16)
	l.print()
	l.NthAdd(17, 8)
	l.print()
	l.NthAdd(14, 3)
	l.print()
	l.Remove(3)
	l.print()
	l.Remove(2)
	l.print()
}
```
- Find a node by its value.
- Reverse a linked list.
- Check if a linked list is cyclic.
- Find the middle element of a linked list.
- Merge two sorted linked lists.
- Detect a Cycle in a Linked List
- Remove Nth Node From End of List
### Stacks Queues
- Implement a stack using an array.
```go
type stack struct {
	d []string
}

func NewStack(d []string) *stack {
	return &stack{
		d: d,
	}
}
func (s *stack) push(data string) {
	s.d = append(s.d, data)
}
func (s *stack) pop() string {
	if len(s.d) < 1 {
		return ""
	}
	t := s.d[len(s.d)-1]
	s.d = s.d[:len(s.d)-1]
	return t
}
```
- Valid Parentheses
```go
func validParenthesesCheck(arr []string) bool {
	o := NewStack([]string{})
	m := make(map[string]string)
	m["{"] = "}"
	m["["] = "]"
	m["("] = ")"
	m["}"] = "{"
	m["]"] = "["
	m[")"] = "("
	for i := 0; i < len(arr); i++ {
		char := arr[i]
		if len(o.d) > 0 {
			oelem := o.pop()
			if m[char] != oelem {
				o.push(oelem)
				o.push(char)
			}
		} else {
			o.push(char)
		}
	}
	if len(o.d) == 0 {
		return true
	} else {
		return false
	}
}
```
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