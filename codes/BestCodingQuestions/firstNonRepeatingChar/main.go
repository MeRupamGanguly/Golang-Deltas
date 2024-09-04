package main

import "fmt"

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
func main() {
	fmt.Println(nonRepeatingFirstChar("najinjia"))
}
