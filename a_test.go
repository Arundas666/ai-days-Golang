package main

import "testing"

func TestMergeAndSortUnique(t *testing.T) {
    slice1 := []int{1, 2, 3, 4}
    slice2 := []int{3, 4, 5, 6}

    expected := []int{1, 2, 3, 4, 5, 6}

    result, err := MergeAndSortUnique(slice1, slice2)

    if err != nil {
        t.Errorf("Expected error to be nil, but got %v", err)
    }

    if len(result) != len(expected) {
        t.Errorf("Expected result length to be %d, but got %d", len(expected), len(result))
    }

    for i := 0; i < len(result); i++ {
        if result[i] != expected[i] {
            t.Errorf("Expected element at index %d to be %d, but got %d", i, expected[i], result[i])
        }
    }
}