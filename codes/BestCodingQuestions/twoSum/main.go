package main

import "fmt"

func twoSum(arr []int, target int) (int, int) {
	for i := 1; i < len(arr); i++ { // Adjacent Sum check
		if arr[i-1]+arr[i] == target {
			return i, i + 1
		}
	}
	for r := len(arr) - 1; r > 0; r-- { // Check every sequences
		l := 0
		for l < r {
			if arr[l]+arr[r] == target {
				return l, r
			}
			l++
		}
	}
	return 0, 0
}
func twoSumOpt(arr []int, target int) (int, int) {
	hmap := make(map[int]int)
	for i := range arr {
		complement := target - arr[i]
		k, v := hmap[complement]
		if v {
			return i, k
		} else {
			hmap[arr[i]] = i
		}
	}
	return -1, -1
}

func main() {
	arr := []int{2, 3, 8, 6, 9, 11, 15, 13, 12}
	fmt.Println(twoSum(arr, 22))
}
