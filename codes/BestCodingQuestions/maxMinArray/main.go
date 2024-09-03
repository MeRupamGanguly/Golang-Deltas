package main

import "fmt"

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
func average(a []int) int {
	sum := 0
	for i := range a {
		sum += a[i]
	}
	return sum / len(a)
}
func main() {
	arr := []int{7, 3, 18, 6, 9, 11, 20, 13, 12}
	fmt.Println(average(arr))
}
