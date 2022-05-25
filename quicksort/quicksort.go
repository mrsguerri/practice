package quicksort

import "fmt"

func Quicksort(array []int, start, end int) []int {
	//Steps for a quicksort:
	//1. Find a pivot (can be the first element, middle element, last element, or a random element)
	//2. Put all elements smaller than the pivot to its left, and all elements larger to its right
	pivot := array[(start+end)/2]

	fmt.Println(array, "start ", start, "end ", end, "pivot ", pivot)

	var i, j int = start, end
	for ; i <= j; i, j = i+1, j-1 {
		for array[i] < pivot {
			i++
		}

		for array[j] > pivot {
			j--
		}

		if i < j {
			array[i], array[j] = array[j], array[i]
			i++
			j--
		}
	}

	fmt.Println("start ", start, " end ", end, " i ", i, " j ", j)

	if start < j {
		fmt.Println("left")
		array = Quicksort(array, start, j)
	}
	if i < end {
		fmt.Println("right")
		array = Quicksort(array, i, end)
	}

	return array
}
