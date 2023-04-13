package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := make([]int, 10)

	for loop := 0; loop < 10; loop++ {
		arr[loop] = rand.Intn(100)
	}

	fmt.Println("unsorted :", arr)
	fmt.Println("sorted :", insertionSort(arr))
}

func insertionSort(arr []int) []int {
	for loop := 1; loop < len(arr); loop++ {
		num := arr[loop]
		for loop1 := loop - 1; loop1 > 0; loop1-- {
			if num > arr[loop1] {
				arr[loop-loop1] = num
				break
			}
		}
	}

	return arr
}
