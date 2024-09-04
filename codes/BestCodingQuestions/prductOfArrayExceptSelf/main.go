package main

import "fmt"

func productOfArrayExceptSelf(a []int) []int {
	var outArr []int
	for i := 0; i < len(a); i++ {
		prod := 1
		for j := 0; j < len(a); j++ {
			if j == i {
				continue
			}
			prod = prod * a[j]
		}
		outArr = append(outArr, prod)
	}
	return outArr
}

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(productOfArrayExceptSelf(a))
}
