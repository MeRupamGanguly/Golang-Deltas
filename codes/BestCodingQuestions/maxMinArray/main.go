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
	max := a[0]
	secondMax := a[0]
	for i := range a {
		if a[i] > max {
			secondMax = max
			max = a[i]
		} else {
			if a[i] > secondMax && a[i] != max {
				secondMax = a[i]
			}
		}
	}
	return secondMax
}
func secondMinimum(a []int) int {
	min := a[0]
	secondMin := a[0]
	for i := range a {
		if a[i] < min {
			secondMin = min
			min = a[i]
		} else {
			if a[i] < secondMin && a[i] != min {
				secondMin = a[i]
			}
		}
	}
	return secondMin
}
func average(a []int) int {
	sum := 0
	for i := range a {
		sum += a[i]
	}
	return sum / len(a)
}
func main() {
	arr := []int{2, 5, 8, 1, 3, 6, 7, 0}
	fmt.Println(findMaxMinOfArray(arr))
	fmt.Println(secondMinimum(arr))
	fmt.Println(secondLargest(arr))
	fmt.Println(average(arr))
}
