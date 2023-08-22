package main

import "fmt"

func findAverage(a []int) float64 {
	count := 4
	sum := 0
	for i := 0; i < count; i++ {
		sum += a[i]
	}

	return float64(sum) / float64(count)
}

func main() {
	i := []int{5, 6, 7, 8}
	fmt.Println("average", findAverage(i))
}
