package main

import "fmt"

func missingNumber(a []int) int {
	length := len(a) + 1
	actualSum := length * (length + 1) / 2
	sum := 0
	for i := range a {
		sum += a[i]
	}
	missingNum := actualSum - sum
	return missingNum
}
func main() {
	fmt.Println(missingNumber([]int{1, 3, 4, 5, 6, 7, 8, 9, 10, 11}))
}
