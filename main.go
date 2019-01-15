package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Seed random
func getSeed() {
	timeNow := time.Now().UTC().UnixNano()
	seedInt := timeNow
	rand.Seed(seedInt)
}

// Get random integer
func getRandomInt(min, max int) int {
	getSeed()
	return min + rand.Intn(max-min)
}

// Swap positions
func swap(input []int32, a int, b int) []int32 {
	fmt.Println("Swapping:", input[a], " < - >", input[b])
	temp := input[a]
	input[a] = input[b]
	input[b] = temp
	return input
}

func partition(list []int32, first int, last int) ([]int32, int) {
	pivot := getRandomInt(first, last)
	fmt.Println("Pivot Value:", list[pivot])
	list = swap(list, pivot, last)
	for i := first; i < last; i++ {
		if list[i] <= list[last] {
			swap(list, i, first)
			first++
		}
	}
	list = swap(list, first, last)
	fmt.Println("Partitioned list:", list)
	return list, first
}

func quickSort(list []int32, first int, last int) []int32 {
	if first < last {
		fmt.Println("First < Last")
		fmt.Println("First Item:", list[first])
		fmt.Println("Last Item:", list[last])
		list, pivot := partition(list, first, last)
		fmt.Println("Pivot point:", pivot)
		fmt.Println("Pivot Value:", list[pivot])
		list = quickSort(list, first, pivot-1)
		list = quickSort(list, pivot+1, last)
	} else {
		fmt.Println("Array contains less than or only one element")
	}
	fmt.Println(list)
	return list
}

func sortList(list []int32) []int32 {
	fmt.Println("Sorting list")
	list = quickSort(list, 0, len(list)-1)
	return list
}

func main() {
	example := []int32{90, 6456, 20, 34, 65, -1, 54, -15, 1, -999, 55, 529, 0}
	fmt.Println(example)
	example = sortList(example)
	fmt.Println(example)
}
