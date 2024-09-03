package main

import "fmt"

func mergeSortedArray(a1, a2 []int) []int {
	var outArr []int
	l := 0
	r := 0
	for l < len(a1) && r < len(a2) {
		if a1[l] == a2[r] {
			outArr = append(outArr, a1[l], a2[r])
			l++
			r++
		} else {
			if a1[l] < a2[r] {
				outArr = append(outArr, a1[l])
				l++
			} else {
				if a1[l] > a2[r] {
					outArr = append(outArr, a2[r])
					r++
				}
			}
		}
	}
	outArr = append(outArr, a1[l:]...)
	outArr = append(outArr, a2[r:]...)
	return outArr
}
func main() {
	a1 := []int{1, 2, 4, 6}
	a2 := []int{7, 8, 15, 17, 19}
	fmt.Println(mergeSortedArray(a1, a2))
}
