package main

import (
	"runtime"
	"sync"

	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	arrLen  = 1e8
	intSize = 1e6
)

var (
	waiter sync.WaitGroup
)

type queue chan ([]int)

func createQueue(len int) *queue {
	var q queue = make(chan ([]int), len)
	return &q
}

func (q *queue) isEmpty() bool {
	return len(*q) == 0
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

func partitionerThread(s *queue) {
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

func sort(arr []int, cores int) {
	if len(arr) <= 1 {
		return
	}
	qLen := len(arr) / 2

	q := createQueue(qLen)
	waiter.Add(1)
	*q <- arr

	for loop := 0; loop < cores; loop++ {
		go partitionerThread(q)
	}

	for {
		waiter.Wait()
		if q.isEmpty() {
			fmt.Println("sorted")

			var emptySlice []int = make([]int, 0)
			for loop := 0; loop < cores; loop++ {
				*q <- emptySlice
			}
			return
		}

	}

}

func main() {
	var arr []int
	var unsplit string
	fmt.Print("enter array (seperated by ,) :")
	reader := bufio.NewReader(os.Stdin)
	unsplit, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	unsplit = strings.Trim(unsplit, "\n")

	split := strings.Split(unsplit, ",")
	if len(unsplit) != 0 {
		for _, item := range split {
			element, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			arr = append(arr, element)
		}
	} else {
		arr = make([]int, 0)
	}

	startTm := time.Now()
	sort(arr, runtime.NumCPU())
	endTm := time.Now()

	tmTaken := endTm.Sub(startTm)
	fmt.Println(tmTaken)
	fmt.Println(arr)
}
