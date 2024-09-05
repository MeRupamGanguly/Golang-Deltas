package main

import "fmt"

func halfAsecHalfDesc(arr []int) []int {
	sortedArray := sort(arr)
	l := len(sortedArray) / 2
	r := len(sortedArray) - 1
	for l < r {
		temp := arr[l]
		arr[l] = arr[r]
		arr[r] = temp
		l++
		r--
	}
	return arr
}

func sort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[j] < a[i] {
				t := a[i]
				a[i] = a[j]
				a[j] = t
			}
		}
	}
	return a
}

func main() {
	fmt.Println(halfAsecHalfDesc([]int{8, 7, 1, 6, 5, 9}))
	fmt.Println(halfAsecHalfDesc([]int{4, 2, 8, 6, 15, 5, 20}))
}
