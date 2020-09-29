package Num_117_Connect

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//func connect(root *Node) *Node {
//	nodes := make([][]*Node,0)
//	dfs := func(n *Node,level int) {}
//	dfs = func(n *Node, level int) {
//		if n == nil{
//			return
//		}
//		if level >= len(nodes){
//			nodes = append(nodes,[]*Node{n})
//		}else {
//			nodes[level] = append(nodes[level],n)
//		}
//		level++
//		dfs(n.Left,level)
//		dfs(n.Right,level)
//	}
//	dfs(root,0)
//	for i:=0;i<len(nodes);i++{
//		for j:=0;j<len(nodes[i])-1;j++{
//			nodes[i][j].Next = nodes[i][j+1]
//		}
//	}
//	return root
//}

//func connect(root *Node) *Node {
//	if root == nil {
//		return root
//	}
//	queue := []*Node{root}
//	for len(queue) > 0 {
//		l := len(queue)
//		for i := 0; i < l; i++ {
//			cur := queue[i]
//			if i != l-1{
//				cur.Next = queue[i+1]
//			}
//			if cur.Left != nil {
//				queue = append(queue, queue[i].Left)
//			}
//			if cur.Right != nil {
//				queue = append(queue, queue[i].Right)
//			}
//		}
//		queue = queue[l:]
//	}
//	return root
//}

func connect(root *Node) *Node {
	s := root
	var ns, ls *Node
	helper := func(n *Node) {
		if n == nil {
			return
		}
		if ns == nil {
			ns = n
		}
		if ls != nil {
			ls.Next = n
		}
		ls = n
	}

	for s != nil {
		for ; s != nil; s = s.Next {
			helper(s.Left)
			helper(s.Right)
		}
		s, ns, ls = ns, nil, nil
	}
	return root
}
