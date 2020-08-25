package Num_133_cloneGraph

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	cg := func(node *Node) *Node { return nil }
	cg = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if _, ok := visited[node]; ok {
			return visited[node]
		}
		visited[node] = &Node{Val: node.Val}
		for _, v := range node.Neighbors {
			visited[node].Neighbors = append(visited[node].Neighbors, cg(v))
		}
		return visited[node]
	}
	return cg(node)
}
