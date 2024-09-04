package main

import "fmt"

func maxSubArray(a []int) []int {
	sumMap := make(map[int][]int)
	for i := range a {
		paticipants := []int{}
		sum := 0
		for j := i; j < len(a); j++ {
			sum += a[j]
			paticipants = append(paticipants, j)
		}
		sumMap[sum] = paticipants
	}
	return sumMap[0]
}
func main() {
	fmt.Println(maxSubArray([]int{2, 3, -8, 6, 9, 11, -15, 13, 12}))
}
