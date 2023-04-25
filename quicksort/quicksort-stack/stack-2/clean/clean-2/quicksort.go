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
	var spliter int
	spliter = arr[len(arr)-1]
	//im asumming 500 , look into it \
	if len(arr) > 500 {
		spliter = findMedian(arr)
	}
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
	//TODO : too bigggg , make it the smallest you can
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

func SelectionSort(arr []int) {
	//use bitonic sort if vectorizing
	if len(arr) <= 1 {
		return
	}

	for loop := len(arr) - 1; loop > 0; loop-- {
		maxI := max(arr)
		//this is ok cuz arr slice pointer
		arr[maxI], arr[loop] = arr[loop], arr[maxI]
	}

}

//return largest number's index
func max(arr []int) int {
	//cannot pass a empty array , deal with it
	LargestIndex := 0
	largestValue := arr[0]
	for loop, currentValue := range arr {
		if currentValue > largestValue {
			LargestIndex = loop
		}
	}
	return LargestIndex
}

func findMedian(arr []int) int {
	//eventualy make it take a samples int argument and use that instad of 21
	//im just guessing 21 elements is enough , I should look into this
	//we find 21 = len(arr)/ x
	if len(arr) <= 21 {
		return arr[len(arr)/2]
	}
	//you should change this to not send it here and upto some threashold
	//just sort it without smapling
	sampleArr := make([]int, 21)
	gap := len(arr) / 21
	sampleLoop := 0
	for loop := 0; loop < len(arr); loop += gap {
		sampleArr[sampleLoop] = arr[loop]
		sampleLoop++
	}

	for loop := len(arr) - 1; loop > 18; loop-- {
		maxI := max(arr)
		//this is ok cuz arr slice pointer
		arr[maxI], arr[loop] = arr[loop], arr[maxI]
	}
	return arr[19]
}
