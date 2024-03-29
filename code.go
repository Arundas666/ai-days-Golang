package main

import (
	"fmt"
	"sort"
)

func MergeAndSortUnique(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 && len(slice2) == 0 {
		return nil, fmt.Errorf("both slices are empty")
	}

	merged := append(slice1, slice2...)
	uniqueMap := make(map[int]bool)
	var uniqueSlice []int

	for _, v := range merged {
		if _, ok := uniqueMap[v]; !ok {
			uniqueSlice = append(uniqueSlice, v)
			uniqueMap[v] = true
		}
	}

	sort.Ints(uniqueSlice)

	return uniqueSlice, nil
}