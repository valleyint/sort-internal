package main

import (
	"sleep"
)

type pipeE struct {
	arr []int
}

func mkPipe(size int) *pipeE {
	return &pipeE{arr: make([]int, size)}
}

func (p *pipeE) put(x int) {
	p.arr = append(p.arr, x)
}

func (p *pipeE) remove() int {
	x := p.arr[0]
	p.arr = p.arr[:1]
	return x
}

type partition struct {
	arr     []int
	x       int
	larger  pipeE
	smaller pipeE
}

func quicksort(arr []int) {
}

func doPartition(arr []int, x int) {
	p := partition{arr: arr, x: x}

	go p.scanRight()
	go p.scanLeft()
	time.sleep(100)
	//TODO: you need to get rid of ^ line , its FOR TESTING ONLY!!!
	
	for loop := 0 
}

func (p *partition) scanRight() {
	p.smaller = mkPipe()

	for loop := 0; loop < len(arr); loop++ {
		if p.arr[loop] < p.x {
			p.smaller.writer.Write(p.arr[loop])
		}
	}
}

func (p *partition) scanLeft() {
	p.larger = mkPipe()

	for loop := len(p.arr) - 1; loop > 0; loop-- {
		if p.arr[loop] > p.x {
			p.larger.writer.Write(p.arr[loop])
		}
	}
}

func main() {

}
