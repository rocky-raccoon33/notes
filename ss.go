package main

import "fmt"

func selection_sort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex], arr[j]
			}
		}
	}

	fmt.Printf("after selection_sort: %+v", arr)
}

func main() {
	selection_sort([]int{4, 3, 1, 6, 45, 0})
}
