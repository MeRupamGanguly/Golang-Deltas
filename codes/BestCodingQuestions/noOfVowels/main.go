package main

import (
	"fmt"
	"strings"
)

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
func main() {
	fmt.Println(noOfVowels("Haellr"))
}
