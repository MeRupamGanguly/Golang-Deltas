package main

import "fmt"

func occurrencesChar(s string, input string) int {
	arr := []rune(s)
	charMap := make(map[rune]int)
	for i := 0; i < len(arr); i++ {
		v, ok := charMap[arr[i]]
		if ok {
			charMap[arr[i]] = v + 1
		} else {
			charMap[arr[i]] = 1
		}
	}
	inputRune := []rune(input)
	return charMap[inputRune[0]]
}
func main() {
	fmt.Println(occurrencesChar("Hello World", "l"))
}
