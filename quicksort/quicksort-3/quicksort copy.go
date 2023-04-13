package main

//TODO-2 : make it work on all comparable types , not jusr ints
//TODO-1 : paralllise it

import (
	"flag"
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
	isTesting *bool
	waiter    sync.WaitGroup
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
	var s stack = make(chan ([]int), 1000)
	return &s
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

// func makeArr(length int) []int {
// 	arr := make([]int, length)
// 	for loop := 0; loop < length; loop++ {
// 		arr[loop] = rand.Intn(1e6)
// 	}
// 	return arr
// }

func partition(arr []int) {
	//add s *stack back when needed

	defer waiter.Done()
	if len(arr) <= 1 {
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
	// *s <- arr[:partitionSeam]
	// *s <- arr[partitionSeam+1:]

	a := arr[:partitionSeam]
	b := arr[partitionSeam+1:]

	if len(a) > 1 {
		waiter.Add(1)
		go partition(arr[:partitionSeam])
	}

	if len(b) > 1 {
		waiter.Add(1)
		go partition(arr[partitionSeam+1:])
	}

}

func sort(arr []int) {
	if len(arr) == 0 {
		return
	}

	// s := createStack()
	// *s <- arr
	//for {
	//x := <-*s
	// if len(x) == 0 {
	// 	continue
	// }
	// 	go partition(x, s)
	// 	fmt.Println("pulled", x)
	// 	fmt.Println(arr)
	// }

	waiter.Add(1)
	go partition(arr)
	waiter.Wait()
}

func init() {
	initTstMd()
}

func main() {
	//gnereate the array here into slice arr

	//profr := runProfile("CPU")
	var arr []int
	if *isTesting == false {
		arr = aux.MakeArr(arrLen)
	} else {
		arr = []int{5, 6, 3, 6, 32, 2}
		/* 2 test cases : 3, 7, 2, 6, 3, 6, 4 or
		5, 6, 3, 6, 32, 2
		*/
	}

	fmt.Println("made	")
	startTm := time.Now()

	sort(arr)
	endTm := time.Now()

	tmTaken := endTm.Sub(startTm)
	fmt.Println(tmTaken)
	if *isTesting == true {
		fmt.Println("arr", arr)
	}
	//profr.Stop()
	//fmt.Println("end", arr)
}
