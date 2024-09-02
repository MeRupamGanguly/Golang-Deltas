package main

func twoSum(arr []int, target int) (int, int) {
	for i := 0; i < len(arr); i++ { // Adjacent Sum check
		if arr[i]+arr[i+1] == target {
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

func main() {
	arr := []int{2, 3, 8, 6, 9, 11, 15, 13, 12}
	twoSum(arr, 14)
}
