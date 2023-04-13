package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := make([]int, 8)

	for loop := 0; loop < 8; loop++ {
		arr[loop] = rand.Intn(100)
	}

	fmt.Println("unsorted :", arr)
	DISort(arr)
	fmt.Println("sorted :", arr)
}

func DISort(arr []int) {
	// stage 1 , where every 4 elements paired and sorted
	for loop := 0; loop < len(arr)/2; loop++ {
		spIdx := (loop + 4)
		checkAndSwap(arr, loop, spIdx)
	}

	for loop := 0; loop < len(arr)/2; loop++ {
		spIdx := (loop + 2)
		checkAndSwap(arr, loop, spIdx)
	}

	for loop := 0; loop < len(arr); loop += 2 {
		spIdx := (loop + 1)
		checkAndSwap(arr, loop, spIdx)
	}
}

func checkAndSwap(arr []int, pos1 int, pos2 int) {
	if arr[pos1] < arr[pos2] {
		temp := arr[pos1]
		arr[pos1] = arr[pos2]
		arr[pos2] = temp
	}
}


