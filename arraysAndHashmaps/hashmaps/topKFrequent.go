package main

import "math"

// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

// //

// // Example 1:

// Input: nums = [1,1,1,2,2,3], k = 2
// // Output: [1,2]

// // Example 2:

// Input: nums = [1], k = 1
// // Output: [1]

//
func topKFrequent(nums []int, k int) []int {
    mapNum := make(map[int]int)
    for _, num := range nums {
        mapNum[num]++
    }

    res := []int{}
    mostFreq := math.Inf(-1) 
    for i := 0; i < k; i++ {
        mostFreq = math.Inf(-1)
        var key int
        for num, freq := range mapNum {
            if float64(freq) > mostFreq {
                mostFreq = float64(freq)
                key = num
            }
        }
        res = append(res, key)
        delete(mapNum, key)
    }

    return res
}