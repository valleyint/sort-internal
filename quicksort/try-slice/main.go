package main

import "fmt"

func main() {
	arr := [20]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 22, 33, 44, 55, 66, 77, 88, 99}
	var arrslice []int = arr[0:]
	slice1 := arrslice[:11]
	slice1 = append(slice1, -1)
	fmt.Println(arr)

}

func runThrough(arr [][]int) {
}
