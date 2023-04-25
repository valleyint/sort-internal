package main

import (
	aux "quicksort/auxilary"
	"sync"
	"time"

	"fmt"
)

const (
	arrLen  = 1e2
	intSize = 1e3
)

var maxCalls int = 0

func main() {
	arr := aux.MakeArr(arrLen, intSize)

	//var profr interface{ Stop() }

	println("made")

	startTm := time.Now()
	sort(arr, 2)
	//SelectionSort(arr)
	endTm := time.Now()

	tmTaken := endTm.Sub(startTm)
	fmt.Println(tmTaken)
	fmt.Println(arr[:60])

	fmt.Println(aux.IsSorted(arr))
	// arr := []int{1, 2, 34, 5, 67, 9, 9, 0, 3, 7, 8, 4}
	// SelectionSort(arr)
	// fmt.Println(arr)
}

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
	// if len(arr) > 500 {
	// 	spliter = findMedian(arr)
	// }
	partitionSeam := 0

	for loop := 0; loop < len(arr); loop++ {
		if arr[loop] < spliter {
			arr[loop], arr[partitionSeam] = arr[partitionSeam], arr[loop]
			partitionSeam++
		}
	}
	//if arr[partitionSeam] < spliter {
	arr[partitionSeam], arr[len(arr)-1] = arr[len(arr)-1], arr[partitionSeam]
	//}

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

		if len(arr) < 20 {
			SelectionSort(arr)
			continue
		}

		a, b := partition(arr)

		// if len(a) > len(b) {
		// 	if len(a) > 1 {
		// 		*s <- a
		// 		waiter.Add(1)
		// 	}

		// 	if len(b) > 1 {
		// 		arr = b
		// 		continue
		// 	}
		// 	waiter.Done()
		// } else {
		// 	if len(b) > 1 {
		// 		*s <- b
		// 		waiter.Add(1)
		// 	}

		// 	if len(a) > 1 {
		// 		arr = a
		// 		continue
		// 	}
		// 	waiter.Done()
		// }
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

func SelectionSort(arr []int) {
	//use bitonic sort if vectorizing
	if len(arr) <= 1 {
		return
	}

	for loop := len(arr) - 1; loop > 0; loop-- {
		_ = max(arr)
		//this is ok cuz arr slice pointer
		//arr[maxI], arr[loop] = arr[loop], arr[maxI]
		arr = arr[:len(arr)-1]
		//fmt.Println(arr)
		//fmt.Println(arr, maxI, loop)
	}
	//fmt.Println(arr)

}

//return largest number's index
func max(arr []int) int {
	maxCalls += 1
	//cannot pass a empty array , deal with it
	LargestIndex := 0
	largestValue := arr[0]
	//for loop, currentValue := range arr {
	for loop := 0; loop < len(arr); loop++ {
		currentValue := arr[loop]
		if currentValue > largestValue {
			//fmt.Println(arr[loop])
			LargestIndex = loop
			largestValue = arr[loop]
			//Println(largestValue, arr[loop])
		}
	}
	//fmt.Println(arr)
	arr[LargestIndex], arr[len(arr)-1] = arr[len(arr)-1], arr[LargestIndex]
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
	for loop := 0; sampleLoop < 21 && loop < len(arr); loop += gap {
		sampleArr[sampleLoop] = arr[loop]
		sampleLoop++
	}

	for loop := len(arr) - 1; len(arr) == 0; loop-- {
		//fmt.Println(sampleArr)
		maxI := max(sampleArr)
		//this is ok cuz arr slice pointer
		//fmt.Println(arr[maxI], arr[loop], maxI)
		arr[maxI], arr[loop] = arr[loop], arr[maxI]
		sampleArr = sampleArr[:len(sampleArr)-1]
	}
	return arr[19]
}
