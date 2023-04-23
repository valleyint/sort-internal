package main

import (
	"fmt"
	aux "quicksort/auxilary"
	"runtime"
	"time"
)

func main() {
	num := runtime.NumCPU()

	arr := aux.MakeArr(1e8)
	results := make([]time.Duration, n)
	for loop := 1; loop <= num; loop++ {
		sort(arr, loop)
	}

	fmt.Println(results)
}
