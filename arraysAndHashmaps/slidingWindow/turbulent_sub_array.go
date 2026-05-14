package main

func maxTurbulenceSize(arr []int) int {
    if len(arr) == 1 {
        return 1
    }
	if len(arr) == 2 {
		if arr[0] == arr[1] {
			return 1
		}
		return 2
	}


    maxLen := 1
	start := 0
    allEqual:=true
    for i := 1; i < len(arr); i++ {
        if arr[i-1]!=arr[i]{
            allEqual=false
        }
        if i == len(arr)-1 || !check(arr[i-1], arr[i], arr[i+1]) {
            maxLen = max(maxLen, i-start+1)
            start = i
        }
    }
    if allEqual{
        return 1
    }
    return maxLen
}

func check(a, b, c int) bool {
    return (a < b && b > c) || (a > b && b < c)
}


