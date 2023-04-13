package main

import (
	"fmt"
)

func main() {
	//gnereate the array hesre into slice arr

	arr := []int{}
	//q := 0
	//r := arr[len(arr)-1]
	//3, 7, 2, 6, 3, 6, 4
	//5, 6, 3, 6, 32, 2

	// for loop := 1; loop < len(arr)-1; loop++ {
	// 	if arr[loop] > r {
	// 		swap(arr, loop, q)
	// 		q++
	// 		fmt.Println(arr)
	// 	} else {
	// 		fmt.Println("no")
	// 	}
	// }

	// swap(arr, q, len(arr)-1)
	// fmt.Println(arr)
	partition3(arr)
	fmt.Println(arr)
	//fmt.Println("end", arr)
}

// func partition(arr []int, min int, max int, rep bool) {
// 	q := 0
// 	r := arr[max]
// 	fmt.Println(min, max, r)

// 	for loop := min + 1; loop <= max; loop++ {
// 		if arr[loop] > r {
// 			swap(arr, loop, q)
// 			fmt.Println(loop, q)
// 			q++
// 			fmt.Println(arr)
// 			// } else if arr[loop] == r {
// 			// 	swap(arr, loop, r-1)
// 			// 	loop--
// 		} else {
// 			fmt.Println("no")
// 		}
// 	}

// 	swap(arr, q+1, max)
// 	fmt.Println(arr, "swapped", q, max)
// 	// 	if rep == true {
// 	// 		fmt.Println("next", q)
// 	// 		partition(arr, q, max, false)
// 	// 		fmt.Println("next")
// 	// 		partition(arr, min, q, false)
// 	// 	}
// }

func partition2(arr []int) {
	spliter := arr[len(arr)/2]
	fmt.Println(spliter)
	leftHand := 0
	rightHand := len(arr) - 1
	for leftHand <= rightHand {
		for arr[leftHand] < spliter {
			leftHand++
		}

		for arr[rightHand] > spliter {
			rightHand--
		}

		if leftHand <= rightHand {
			arr[leftHand], arr[rightHand] = arr[rightHand], arr[leftHand]
			rightHand--
			leftHand++
		}
	}

}

func partition3(arr []int) {
	if len(arr) == 0 {
		return
	}

	spliter := arr[len(arr)-1]
	fmt.Println(spliter, arr)
	partitionSeam := 0

	for loop := 0; loop < len(arr); loop++ {
		if arr[loop] < spliter {
			arr[loop], arr[partitionSeam] = arr[partitionSeam], arr[loop]
			partitionSeam++
		}
	}

	arr[partitionSeam], arr[len(arr)-1] = arr[len(arr)-1], arr[partitionSeam]

	//this is only for testing, make it iterative

	partition3(arr[:partitionSeam])
	partition3(arr[partitionSeam+1:])
}
