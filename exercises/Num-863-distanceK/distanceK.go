package Num_863_DistanceK

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) (ans []int) {
	if root == nil || target == nil {
		return
	}
	var parents = map[*TreeNode]*TreeNode{}
	var dfs func(parent, cur *TreeNode)
	dfs = func(parent, cur *TreeNode) {
		if cur == nil {
			return
		}
		parents[cur] = parent
		dfs(cur, cur.Left)
		dfs(cur, cur.Right)
	}
	dfs(nil, root)

	var find func(cur, from *TreeNode, length int)
	find = func(cur, from *TreeNode, length int) {
		if cur == nil {
			return
		}
		if length == k {
			ans = append(ans, cur.Val)
		}
		if parents[cur] != from {
			find(parents[cur], cur, length+1)
		}
		if cur.Left != from {
			find(cur.Left, cur, length+1)
		}
		if cur.Right != from {
			find(cur.Right, cur, length+1)
		}
	}
	find(target, nil, 0)
	return
}
