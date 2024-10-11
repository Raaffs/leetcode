package main

func sortColors(nums []int)  {
    colors:=make(map[int]int)
    sortedColors:=[]int{}
    for _,num :=range nums{
        if num==0{
            colors[0]++
        }
        if num==1{
            colors[1]++
        }
        if num==2{
            colors[2]++
        }
    }
    for i:=0;i<colors[0];i++{
        sortedColors=append(sortedColors,0)
    }
    for i:=0;i<colors[1];i++{
        sortedColors=append(sortedColors,1)
    }
    for i:=0;i<colors[2];i++{
        sortedColors=append(sortedColors,2)
    }

}