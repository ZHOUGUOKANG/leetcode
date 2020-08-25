package Num_95_generateTrees

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var cache map[string][]*TreeNode

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	cache = make(map[string][]*TreeNode)
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	cacheKey := fmt.Sprintf("%d-%d", start, end)
	if _, ok := cache[cacheKey]; ok {
		return cache[cacheKey]
	}
	var allTrees []*TreeNode
	for i := start; i <= end; i++ {
		left := helper(start, i-1)
		right := helper(i+1, end)
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i}
				root.Left = l
				root.Right = r
				allTrees = append(allTrees, root)
			}
		}
	}
	cache[cacheKey] = allTrees
	return allTrees
}
