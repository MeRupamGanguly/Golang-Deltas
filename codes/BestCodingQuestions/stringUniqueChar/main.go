package main

import "fmt"

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
func main() {
	fmt.Println(checkUnique("helo"))
}
