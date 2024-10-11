package main
func maxSubArray(nums []int) int {
    curr:=nums[0]
    maxNum:=nums[0]
    for i:=1;i<len(nums);i++{
        if nums[i]>nums[i]+curr{
            curr=nums[i]
        }else{
            curr+=nums[i]
        }
        if curr>maxNum{
            maxNum=curr
        }
        if curr<0{
            curr=0
        }
    }
    return maxNum
}