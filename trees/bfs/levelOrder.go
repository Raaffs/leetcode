package main
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    var result [][]int
    nodes := []*TreeNode{root}
    var row []int
    for len(nodes) > 0 {
        size := len(nodes)

        for i := 0; i < size; i++ {
            currentNode := nodes[0]
            nodes = nodes[1:]
            row = append(row, currentNode.Val)

            if currentNode.Left != nil {
                nodes = append(nodes, currentNode.Left)
            }

            if currentNode.Right != nil {
                nodes = append(nodes, currentNode.Right)
            }
        }

        result = append(result, row)
        row=[]int{}
    }

    return result
}