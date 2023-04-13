package auxilary

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	prof "github.com/pkg/profile"
)

type Options struct {
	IsTest    *bool
	IsProfile *bool
	Cores     *int
}

func InitMds() Options {
	c := runtime.NumCPU()

	isTesting := flag.Bool("t", false, "testing mode")
	isProfiling := flag.Bool("p", false, "profiling on/off")
	cores := flag.Int("c", c, "manual number of cores specification")
	flag.Parse()
	return Options{IsTest: isTesting, IsProfile: isProfiling, Cores: cores}
}

func RunProfile(member string) interface{ Stop() } {
	//TODO :make it work with multiple options at once
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

func MakeArr(length int, numSize ...int) []int {
	sz := 0
	if len(numSize) > 0 {
		sz = numSize[0]
	} else {
		sz = 1e6
	}
	arr := make([]int, length)
	for loop := 0; loop < length; loop++ {
		arr[loop] = rand.Intn(sz)
	}
	return arr
}

func Readout(arr []int) {
	for {
		fmt.Println(arr)
		time.Sleep(1 * time.Second)
	}
}

func IsSorted(arr []int) bool {
	for loop := 1; loop < len(arr); loop++ {
		if arr[loop] < arr[loop-1] {
			//fmt.Println(arr[loop-1], arr[loop])
			return false
		}

	}
	return true
}