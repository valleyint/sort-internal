package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

func compASM() {
	//add vectorized instructions here
	TEXT("compareAndSwap", NOSPLIT, "func(arr [8]float32, psm [8]int)[8]bool")
	Doc("compares numbers for higher than and swaps them")

	//load the first array 
	arr := Mem{Base: Load(Param("arr"), GP64())}
	x1 := YMM()
	VMOVUPS(arr.Offset(0), x1)

	//loads the array which contains the partition element
	compArr := Mem{Base: Load(Param("psm").Base(), GP64())}
	x2 := YMM()
	VMOVUPS(compArr.Offset(0), x2)

	//assign the part of memory to put the set of arrays to 

}

func main() {
	TEXT("compASM", NOSPLIT, "func(arr []float32, psm []float32, int )[]bool")
	Doc("compares numbers for higher than and swaps them ASM function. is wrapped")
	arr := Mem{Base: Load(Param("arr").Base(), GP64())}
	x1 := YMM()
	VMOVUPS(arr.Offset(0), x1)

	compArr := Mem{Base: Load(Param("psm").Base(), GP64())}
	x2 := YMM()
	VMOVUPS(compArr.Offset(0), x2)

	VCMPPS(x1, x2)
	Generate()
}

func selectionSort (arr []float32) {
	
}
