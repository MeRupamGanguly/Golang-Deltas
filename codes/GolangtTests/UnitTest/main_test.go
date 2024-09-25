package main

import (
	"fmt"
	"testing"
	"time"

	"math/rand"
)

type inps struct {
	a []int
	t int
}
type exps struct {
	first  int
	second int
}

func TestComplexAlgo(t *testing.T) {
	a1 := []int{2, 4, 6, 8, 10, 12}
	t1 := 10
	a2 := []int{1, 2, 4, 6, 8, 10}
	t2 := 16
	a3 := []int{2, 3, 4, 6, 12}
	t3 := 15
	tests := []struct {
		input    inps
		expected exps
	}{
		{inps{a: a1, t: t1}, exps{first: 0, second: 3}},
		{inps{a: a2, t: t2}, exps{first: 3, second: 5}},
		{inps{a: a3, t: t3}, exps{first: 1, second: 4}},
	}

	for i := range tests {
		f, s := ComplexAlgo(tests[i].input.a, tests[i].input.t)
		if f != tests[i].expected.first && s != tests[i].expected.second {
			t.Errorf("ComplexAlgo(%d)=%d but got %d,%d", tests[i].input, tests[i].expected, f, s)
		}
	}
}
func BenchmarkComplexAlgo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.NewSource(time.Now().Unix())
		arraySize := 10
		randomArray := generateRandomArray(arraySize)
		targetSum := rand.Intn(20*2-2*1+1) + 2*1
		f, s := ComplexAlgo(randomArray, targetSum)
		fmt.Println(randomArray, targetSum, " -> ", f, s)
	}
}

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(20-1+1) + 1
	}
	return arr
}
