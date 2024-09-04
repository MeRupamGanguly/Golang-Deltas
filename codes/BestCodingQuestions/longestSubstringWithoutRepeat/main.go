package main

import "fmt"

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

func main() {
	fmt.Println(longestSubstringWithoutRepeat("Hllo"))
}
