package day8

func Add(a, b int) int {
	return a + b
}

func Max(num []int) int {
	max := num[0]
	for _, n := range num {
		if n > max {
			max = n
		}
	}
	return max
}
