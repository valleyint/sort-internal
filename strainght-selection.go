package main

import (
	"fmt"
	//	"math/rand"
)

func straightSort(arr []int) []int {
	for loop := 0; loop < len(arr); loop++ {
		num := arr[loop]
		idx := loop
		for i := loop + 1; i < len(arr); i++ {
			if arr[i] < num {
				idx = i
				num = arr[i]
			}
		}
		arr[idx] = arr[loop]
		arr[loop] = num
	}

	return arr
}

// func main() {
// 	arr := make([]int, 10)

// 	for loop := 0; loop < 10; loop++ {
// 		arr[loop] = rand.Intn(100)
// 	}

// 	fmt.Println("unsorted :", arr)
// 	fmt.Println("sorted :", straightSort(arr))
// }

func main() {
	fmt.Println(straightSort([]int{}))
}
