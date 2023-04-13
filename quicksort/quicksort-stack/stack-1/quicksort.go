package main

//TODO-2 : make it work on all comparable types , not jusr ints
//TODO-1 : paralllise it

import (
	"fmt"
	"sync"
	"time"

	aux "quicksort/auxilary"

	prof "github.com/pkg/profile"
)

const (
	arrLen = 1e8
)

var (
	//isTesting   *bool
	//isProfiling *bool
	opts   aux.Options
	waiter sync.WaitGroup
)

// type waiter struct {
// 	mut sync.Mutex
// 	x   int
// }

// func (w *waiter) add(inc int) {
// 	w.mut.Lock()
// 	w.x += inc
// 	w.mut.Unlock()
// }

// func (w *waiter) running() int {
// 	return w.x
// }

type stack chan ([]int)

func createStack() *stack {
	var s stack = make(chan ([]int), 10000000)
	return &s
}

func (s *stack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	} else {
		return false
	}
}

func runProfile(member string) interface{ Stop() } {
	if member == "CPU" {
		x := prof.Start(prof.CPUProfile)
		prof.ProfilePath("./prof")
		return x
	} else if member == "GO" {
		x := prof.Start(prof.GoroutineProfile)
		prof.ProfilePath("./prof")
		return x
	}
	return prof.Start()
}

func partition(s *stack) {
	var arr []int
	for {
		//arr = <-*s

		if len(arr) <= 1 {
			arr = <-*s
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

		a := arr[:partitionSeam]
		b := arr[partitionSeam+1:]

		if len(a) > 1 {
			*s <- a
			waiter.Add(1)
		}

		if len(b) > 1 {
			// waiter.Add(1)
			// *s <- b
			arr = b
			//waiter.Done()
			continue
		}
		waiter.Done()

		arr = <-*s
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
		go partition(s)
	}

	for {
		waiter.Wait()
		if s.isEmpty() {
			fmt.Println("sorted")
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
		profr = runProfile("CPU")
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
	//fmt.Println(arr)

	println("made	")

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
