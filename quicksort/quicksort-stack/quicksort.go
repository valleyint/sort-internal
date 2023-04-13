package main

//TODO-2 : make it work on all comparable types , not jusr ints

import (
	"fmt"
	"sync"
	"time"

	aux "quicksort/auxilary"
)

const (
	arrLen = 1e8
)

var (
	opts   aux.Options
	waiter sync.WaitGroup
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

func partitionerThread(s *stack) {
	var arr []int
	for {

		if len(arr) <= 1 {
			arr = <-*s
			if len(arr) == 0 {
				return
			}
		}

		a, b := partition(arr)

		if len(a) > 1 {
			*s <- a
			waiter.Add(1)
		}

		if len(b) > 1 {
			arr = b
			continue
		}
		waiter.Done()

		arr = <-*s
		if len(arr) == 0 {
			return
		}
	}

}

func sort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	s := createStack()
	waiter.Add(1)
	*s <- arr

	cores := 4

	for loop := 0; loop < cores; loop++ {
		go partitionerThread(s)
	}

	for {
		waiter.Wait()
		if s.isEmpty() {
			fmt.Println("sorted")

			var emptySlice []int = make([]int, 0)
			for loop := 0; loop < cores; loop++ {
				*s <- emptySlice
			}
			return
		}

	}

}

func init() {
	opts = aux.InitMds()
}

func main() {
	//gnereate the array here into slice arr

	var profr interface{ Stop() }
	if *opts.IsProfile {
		profr = aux.RunProfile("CPU")
	}

	var arr []int
	if *opts.IsTest == false {
		arr = aux.MakeArr(arrLen)
	} else {
		arr = []int{5, 6, 3, 6, 32, 2}
		/* 2 test cases : 3, 7, 2, 6, 3, 6, 4 or
		5, 6, 3, 6, 32, 2
		*/
	}

	println("made")

	startTm := time.Now()
	sort(arr)
	endTm := time.Now()

	tmTaken := endTm.Sub(startTm)
	fmt.Println(tmTaken)
	if *opts.IsTest {
		fmt.Println("arr", arr)
	}

	if *opts.IsProfile {
		profr.Stop()
	}

	fmt.Println(aux.IsSorted(arr))
}
