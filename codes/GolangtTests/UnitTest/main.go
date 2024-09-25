package main

func ComplexAlgo(a []int, t int) (int, int) {

	m := make(map[int]int)

	for i := range a {
		m[a[i]] = i
	}
	for i := range a {
		snd := t - a[i]
		v, ok := m[snd]
		if ok {
			return i, v
		}
	}
	return -1, -1
}
