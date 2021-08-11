package main

import (
	"container/heap"
	"log"
	"math/rand"
)

func main() {
	log.Println("Started")

	unorderedArray := fill()
	log.Printf("unordered array = %#v\n", unorderedArray)
	bubbleSorted := bubbleSort(unorderedArray)
	log.Printf("bubble sort array = %#v\n", bubbleSorted)

	heapSort(unorderedArray)
}

func fill() []int {
	array := []int{}
	for i := 0; i < 20; i++ {
		array = append(array, rand.Intn(100))
	}
	return array
}

func bubbleSort(array []int) []int {
	n := len(array) - 1
	for n != 1 {
		for i, v := range array[:n] {
			if v > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
			}
		}
		n--
	}
	return array
}

func heapSort(array []int) {
	h := &IntHeap{}
	heap.Init(h)
	for _, v := range array {
		h.Push(v)
	}
	h.Sort()
	log.Printf("heap sort array = %#v\n", h)
}
