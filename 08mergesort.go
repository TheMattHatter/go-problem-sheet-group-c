// Author: Matthew Shiel
// Date: 01-10-2017
// Go problem 08: Merge List and Sort
// Adapted from https://golang.org/pkg/builtin/#append and https://gobyexample.com/sorting



package main

import (
	"fmt"
	"sort"
)

func mergeSort(list1 []int, list2 []int) []int{
	mergedSlice := []int{} 
	mergedSlice = append(list1, list2...) // Conactenate the lists
	sort.Ints(mergedSlice) // Sort the ints

	return mergedSlice
}

func main() {
	// Lists to sort
	list1 := []int{3, 5, 8, 20}
	list2 := []int{33, 1, 9, 12, 17}

	sortedList := mergeSort(list1, list2) // Sort lists into one list
	fmt.Println(sortedList) // Return sorted list
}

