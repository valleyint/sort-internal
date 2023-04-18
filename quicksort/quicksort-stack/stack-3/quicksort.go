package main

//TODO-2 : make it work on all comparable types , not jusr ints
// TOSDO 1 : amek it be able to do others work afteerwords
//note eveything but rhe last 1000 elements sorts properly

import (
	"fmt"
	"sync"
	"time"

	aux "quicksort/auxilary"
)

const (
	arrLen = 1e4
)

var (
	opts    aux.Options
	gotWork sync.WaitGroup
	isFull  bool = false
	done    sync.WaitGroup
)

type stack chan ([]int)

func createStack() *stack {
	var s stack = make(chan ([]int), 10000000)
	return &s
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func partition(arr []int) ([]int, []int) {
	spliter := arr[len(arr)-1]
	partitionSeam := 0

	for loop := 0; loop < len(arr); loop++ {
		if arr[loop] < spliter {
			arr[loop], arr[partitionSeam] = arr[partitionSeam], arr[loop]
			partitionSeam++
		}
	}

	arr[partitionSeam], arr[len(arr)-1] = arr[len(arr)-1], arr[partitionSeam]

	a := arr[:partitionSeam]
	b := arr[partitionSeam+1:]
	return a, b
}

func sequentialSort(s *stack) {
	for {
		if s.isEmpty() {
			return
		}
		a, b := partition(<-*s)
		if len(a) >= 1 {
			*s <- a
		}

		if len(b) >= 1 {
			*s <- b
		}
	}
}

func partitionerThread(gs *stack) {
	arr := <-*gs
	gotWork.Done()
	done.Add(1)
	ls := createStack()

	//for {

	if len(arr) <= 1 {
		arr = <-*gs
		if len(arr) == 0 {
			return
		}
	}
	//idk this ^ required

	a, b := partition(arr)

	if len(a) > 1 {
		if isFull != true {
			*gs <- a
		} else {
			*ls <- a
			sequentialSort(ls)
		}

		//waiter.Add(1)
	}

	if len(b) > 1 {
		if isFull != true {
			*gs <- b
		} else {
			*ls <- b
			sequentialSort(ls)
		}

		//arr = b
		//continue
	}
	//waiter.Done()

	if len(*ls) != 0 {
		arr = <-*ls
	} else {
		done.Done()
		return
	}

	if len(arr) == 0 {
		done.Done()
		return
	}
	//}

}

func sort(arr []int, cores int) {
	if len(arr) <= 1 {
		return
	}

	s := createStack()
	//waiter.Add(1)
	*s <- arr

	for loop := 0; loop < cores; loop++ {
		gotWork.Add(1)
		go partitionerThread(s)
	}

	gotWork.Wait()
	isFull = true

	// for {
	// 	if s.isEmpty() {
	// 		fmt.Println("sorted")

	// 		var emptySlice []int = make([]int, 0)
	// 		for loop := 0; loop < cores; loop++ {
	// 			*s <- emptySlice
	// 		}
	// 		return
	// 	}

	// }
	done.Wait()

}

// func sort(arr []int, cores int) {
// 	if len(arr) <= 1 {
// 		return
// 	}

// 	s := createStack()
// 	waiter.Add(1)
// 	*s <- arr

// 	//add block signaling weather to request to sequrntial stack or to do so to global stack
// 	//filled := //bool suggesting

// 	for loop := 0; loop < cores; loop++ {
// 		go partitionerThread(s)
// 	}

// 	for {
// 		waiter.Wait()
// 		if s.isEmpty() {
// 			fmt.Println("sorted")

// 			var emptySlice []int = make([]int, 0)
// 			for loop := 0; loop < cores; loop++ {
// 				*s <- emptySlice
// 			}
// 			return
// 		}

// 	}

// }

func init() {
	opts = aux.InitMds()
}

func main() {
	//gnereate the array here into slice arr

	var arr []int
	if *opts.IsTest == false {
		arr = aux.MakeArr(arrLen)
	} else {
		arr = []int{5, 6, 3, 6, 32, 2, 45, 2, 5, 15, 3, 6, 32, 5, 2, 5, 2, 1, 5, 5}
		/* 2 test cases : 3, 7, 2, 6, 3, 6, 4 or
		5, 6, 3, 6, 32, 2
		*/
	}

	var profr interface{ Stop() }
	if *opts.IsProfile {
		profr = aux.RunProfile("BLOCK")
	}

	println("made")
	//go aux.Readout(arr)

	startTm := time.Now()
	sort(arr, 4)
	endTm := time.Now()

	tmTaken := endTm.Sub(startTm)
	fmt.Println(tmTaken)
	if *opts.IsTest {
		//fmt.Println("arr", arr)
	}

	if *opts.IsProfile {
		profr.Stop()
	}

	fmt.Println(aux.IsSorted(arr))
	fmt.Println(arr[:100])
}
