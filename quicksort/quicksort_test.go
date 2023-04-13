package main

import (
	"fmt"
	"testing"
)

func qsBenchMark(t *testing.B) {
	arr := makeArr(100)
	t.ResetTimer()
	sort(arr)
	fmt.Println(arr)
}
