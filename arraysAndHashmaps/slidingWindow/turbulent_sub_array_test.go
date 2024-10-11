package main

import (
    "testing"
)

func TestMaxTurbulentSubArray(t *testing.T) {
    // Test case 1: Typical turbulent array
    arr1 := []int{0,1,1,0,1,0,1,1,0,0}
    result1 := maxTurbulenceSize(arr1)
    expected1 := 5 // [4, 2, 10, 7, 8] is the longest turbulent subarray
    if result1 != expected1 {
        t.Errorf("Test case 1 failed: expected %d, got %d", expected1, result1)
    }

    // Test case 2: No turbulence (constant values)
    arr2 := []int{1,1,1,1}
    result2 := maxTurbulenceSize(arr2)
    expected2 := 1 // Since all elements are    equal, the max turbulent subarray is any single element
    if result2 != expected2 {
        t.Errorf("Test case 2 failed: expected %d, got %d", expected2, result2)
    }

    // // Test case 3: Fully turbulent array
    arr3 := []int{1, 2, 1, 2, 1, 2}
    result3 := maxTurbulenceSize(arr3)
    expected3 := 6 // The entire array is turbulent
    if result3 != expected3 {
        t.Errorf("Test case 3 failed: expected %d, got %d", expected3, result3)
    }

    // // Test case 4: Single element array
    arr4 := []int{5}
    result4 := maxTurbulenceSize(arr4)
    expected4 := 1 // The array has only one element
    if result4 != expected4 {
        t.Errorf("Test case 4 failed: expected %d, got %d", expected4, result4)
    }

    // // Test case 5: Edge case with only two elements
    arr5 := []int{100, 100}
    result5 := maxTurbulenceSize(arr5)
    expected5 := 1 // No turbulence in two equal elements
    if result5 != expected5 {
        t.Errorf("Test case 5 failed: expected %d, got %d", expected5, result5)
    }
}
