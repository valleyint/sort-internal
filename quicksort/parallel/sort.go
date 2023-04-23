package main

type queue chan ([]int)

func createQueue(len int) *queue {
	var q queue = make(chan ([]int), len)
	return &q
}

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}

func partition(arr []int, spliter int) ([]int, []int) {
	TODO : rewrite the following such that it can be vectorized easily 
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
