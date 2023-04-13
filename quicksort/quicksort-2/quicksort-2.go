package main

import (
	"flag"
	"fmt"
	"math/rand"

	prof "github.com/pkg/profile"

	_ "net/http/pprof"
)

//TODO : make it work on all comparable types , not jusr ints

func init() {
	initTstMd()
}

const (
	arrLen = 1e6
)

var (
	isTesting *bool
)

type stack [][]int

func createStack() *stack {
	var s stack
	return &s
}

func (s *stack) add(p []int) {
	*s = append(*s, p)
}

func (s *stack) pull() []int {
	i := len(*s) - 1
	x := (*s)[i]
	*s = (*s)[:i]
	return x
}

func (s *stack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	} else {
		return false
	}
}

func initTstMd() {
	isTesting = flag.Bool("t", false, "testing mode")
	flag.Parse()
}

// func runProfile() {
// 	log.Println(http.ListenAndServe("localhost:6060", nil))
// }

func runProfile(member string) interface{ Stop() } {
	if member == "CPU" {
		x := prof.Start(prof.CPUProfile)
		prof.ProfilePath("./prof")
		return x
	}
	return prof.Start()
}

func makeArr(length int) []int {
	arr := make([]int, length)
	for loop := 0; loop < length; loop++ {
		arr[loop] = rand.Intn(1e6)
	}
	return arr
}

func main() {
	//gnereate the array hesre into slice arr

	profr := runProfile("CPU")
	var arr []int
	if *isTesting == false {
		arr = makeArr(arrLen)
	} else {
		arr = []int{3, 7, 2, 6, 3, 6, 4}
	}

	//arr := []int{3, 7, 2, 6, 3, 6, 4}
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
	sort(arr)
	if *isTesting == true {
		fmt.Println(arr)
	}
	profr.Stop()
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

func sort(arr []int) {
	if len(arr) == 0 {
		return
	}

	// spliter := arr[len(arr)-1]
	// fmt.Println(spliter, arr)
	// partitionSeam := 0

	// for loop := 0; loop < len(arr); loop++ {
	// 	if arr[loop] < spliter {
	// 		arr[loop], arr[partitionSeam] = arr[partitionSeam], arr[loop]
	// 		partitionSeam++
	// 	}
	// }

	// arr[partitionSeam], arr[len(arr)-1] = arr[len(arr)-1], arr[partitionSeam]

	//partitionSeam := partition(arr)

	//this is only for testing, make it iterative

	// sort(arr[:partitionSeam])
	// sort(arr[partitionSeam+1:])

	s := createStack()
	s.add(arr)
	for {
		if s.isEmpty() {
			return
		}
		partition(s.pull(), s)
	}
}

func partition(arr []int, s *stack) {
	if len(arr) == 0 {
		return
	}

	spliter := arr[len(arr)-1]
	partitionSeam := 0

	for loop := 0; loop < len(arr); loop++ {
		if arr[loop] < spliter {
			arr[loop], arr[partitionSeam] = arr[partitionSeam], arr[loop]
			partitionSeam++
		}
	}

	arr[partitionSeam], arr[len(arr)-1] = arr[len(arr)-1], arr[partitionSeam]
	//return partitionSeam
	s.add(arr[:partitionSeam])
	s.add(arr[partitionSeam+1:])

}
