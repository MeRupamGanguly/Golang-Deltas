package main

import "fmt"

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
func main() {
	fmt.Println(anagram("listen", "silint"))
}
