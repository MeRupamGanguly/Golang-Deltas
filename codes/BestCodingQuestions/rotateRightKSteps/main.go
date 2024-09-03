package main

import "fmt"

func rotateKStepsRight(a []int, k int) []int {
	part := a[:k]
	a = a[k:]
	a = append(a, part...)
	return a
}
func main() {
	fmt.Println(rotateKStepsRight([]int{2, 3, 8, 0, 9, 11, 0, 13, 0}, 8))
}
