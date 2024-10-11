package main

import "math"

 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
  }
 

func largestValues(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    result:=[]int{}    
    
    nodes := []*TreeNode{root}
    
    for len(nodes) > 0 {
        greatestPerRow :=math.MinInt32
        size := len(nodes)
        for i := 0; i < size; i++ {
            currentNode := nodes[0]
            greatestPerRow=max(greatestPerRow,currentNode.Val)
            
            nodes = nodes[1:]
            if currentNode.Left != nil {
                nodes = append(nodes, currentNode.Left)        
            }

            if currentNode.Right != nil {
                nodes = append(nodes, currentNode.Right)
            }
        }
        result=append(result,greatestPerRow)
    }

    return result

}