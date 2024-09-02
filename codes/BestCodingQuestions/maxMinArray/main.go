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
func main() {
	arr := []int{7, 3, 8, 6, 9, 11, 0, 13, 12}
	fmt.Println(findMaxMinOfArray(arr))
}
