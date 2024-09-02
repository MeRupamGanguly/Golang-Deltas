package main

import "fmt"

func moveZeroes(arr []int) []int {
	outArr := make([]int, len(arr))
	outIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			outArr[outIndex] = arr[i]
			outIndex++
		}
	}
	return outArr
}
func main() {
	arr := []int{2, 3, 8, 0, 9, 11, 0, 13, 0}
	fmt.Println(moveZeroes(arr))
}
