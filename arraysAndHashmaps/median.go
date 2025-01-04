package main

// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

// The overall run time complexity should be O(log (m+n)).
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    length := len(nums1) + len(nums2)
    res := []int{}  
    j, k := 0, 0    
    
    for i := 0; i < length; i++ {
        if j < len(nums1) && k < len(nums2) {
            if nums1[j] < nums2[k] {
                res = append(res, nums1[j])
                j++
            } else {
                res = append(res, nums2[k])  
                k++
            }
        } else if j < len(nums1) {
            res = append(res, nums1[j])  
            j++
        } else if k < len(nums2) {
            res = append(res, nums2[k])
            k++
        }
    }

    if length%2 == 0 {
        mid1 := length / 2
        mid2 := mid1 - 1
        return float64(res[mid1]+res[mid2]) / 2.0
    }

    return float64(res[length/2])
}
