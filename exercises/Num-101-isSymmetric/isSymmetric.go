package Num_101_IsSymmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

//func isSymmetric(root *TreeNode) bool {
//	u, v := root, root
//	var q []*TreeNode
//	q = append(q, u, v)
//	for len(q) > 0 {
//		u, v = q[0], q[1]
//		q = q[2:]
//		if u == nil && v == nil {
//			continue
//		}
//		if u == nil || v == nil {
//			return false
//		}
//		if u.Val != v.Val {
//			return false
//		}
//		q = append(q, u.Left, v.Right, u.Right, v.Left)
//	}
//	return true
//}
