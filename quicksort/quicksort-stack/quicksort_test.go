package main

import (
	aux "quicksort/auxilary"
	"testing"
)

func Test_sort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name : "test case 10 elements : [1,2,3,4,5,6,7,8,9]",
			args : [1,2,3,4,5,6,7,8,9]
		},
		{
			name : "empty slice",
			args : []
		},
		{
			name : "single element",
			args : [3]
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort(tt.args.arr)
		})
	}
}

func BenchmarkSort (b *testing.B) {
	aux.MakeArr()
}
