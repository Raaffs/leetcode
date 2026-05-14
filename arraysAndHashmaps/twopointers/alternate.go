package main

import "fmt"

func alternateCount(nums []int)int{
	res:=0
	j:=0
	for i:=range len(nums){
		if nums[i]!=nums[j]+(i-j)%2{
			if nums[i-1]==nums[i]-1{
				j=i-1
			}else{
				j=i
			}
		}
		res=max(res,i-j+1)
	}
	return res
}

func main()  {
	fmt.Println(alternateCount([]int{1,2,1,2,3,4,3}))
}