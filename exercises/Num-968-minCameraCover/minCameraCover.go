package Num_968_MinCameraCover

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//0
func minCameraCover(root *TreeNode) int {
	var dfs func(node *TreeNode) (a, b, c int)
	dfs = func(node *TreeNode) (a, b, c int) {
		if node == nil {
			return math.MaxInt32, 0, 0
		}
		la, lb, lc := dfs(node.Left)
		ra, rb, rc := dfs(node.Right)
		a = lc + rc + 1
		b = min(a, min(la+rb, lb+ra))
		c = min(a, lb+rb)
		return
	}
	_, ans, _ := dfs(root)
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/**
方法一：递归
思路与算法

本题以二叉树为背景，不难想到用递归的方式求解。本题的难度在于如何从左、右子树的状态，推导出父节点的状态。

为了表述方便，我们约定：如果某棵树的所有节点都被监控，则称该树被「覆盖」。

假设当前节点为 \textit{root}root，其左右孩子为 \textit{left}, \textit{right}left,right。如果要覆盖以 \textit{root}root 为根的树，有两种情况：

若在 \textit{root}root 处安放摄像头，则孩子 \textit{left}, \textit{right}left,right 一定也会被监控到。此时，只需要保证 \textit{left}left 的两棵子树被覆盖，同时保证 \textit{right}right 的两棵子树也被覆盖即可。
否则， 如果 \textit{root}root 处不安放摄像头，则除了覆盖 \textit{root}root 的两棵子树之外，孩子 \textit{left}, \textit{right}left,right 之一必须要安装摄像头，从而保证 \textit{root}root 会被监控到。
根据上面的讨论，能够分析出，对于每个节点 \textit{root}root ，需要维护三种类型的状态：

状态 aa：\textit{root}root 必须放置摄像头的情况下，覆盖整棵树需要的摄像头数目。
状态 bb：覆盖整棵树需要的摄像头数目，无论 \textit{root}root 是否放置摄像头。
状态 cc：覆盖两棵子树需要的摄像头数目，无论节点 \textit{root}root 本身是否被监控到。
根据它们的定义，一定有 a \geq b \geq ca≥b≥c。

对于节点 \textit{root}root 而言，设其左右孩子 \textit{left}, \textit{right}left,right 对应的状态变量分别为 (l_a,l_b,l_c)(l
a
​
 ,l
b
​
 ,l
c
​
 ) 以及 (r_a,r_b,r_c)(r
a
​
 ,r
b
​
 ,r
c
​
 )。根据一开始的讨论，我们已经得到了求解 a,ba,b 的过程：

a = l_c + r_c + 1a=l
c
​
 +r
c
​
 +1
b = \min(a, \min(l_a + r_b, r_a + l_b))b=min(a,min(l
a
​
 +r
b
​
 ,r
a
​
 +l
b
​
 ))
对于 cc 而言，要保证两棵子树被完全覆盖，要么 \textit{root}root 处放置一个摄像头，需要的摄像头数目为 aa；要么 \textit{root}root 处不放置摄像头，此时两棵子树分别保证自己被覆盖，需要的摄像头数目为 l_b + r_bl
b
​
 +r
b
​
 。

需要额外注意的是，对于 \textit{root}root 而言，如果其某个孩子为空，则不能通过在该孩子处放置摄像头的方式，监控到当前节点。因此，该孩子对应的变量 aa 应当返回一个大整数，用于标识不可能的情形。

最终，根节点的状态变量 bb 即为要求出的答案。

代码

C++JavaGolangJavaScriptPython3C

const inf = math.MaxInt32 / 2 // 或 math.MaxInt64 / 2

func minCameraCover(root *TreeNode) int {
    var dfs func(*TreeNode) (a, b, c int)
    dfs = func(node *TreeNode) (a, b, c int) {
        if node == nil {
            return inf, 0, 0
        }
        la, lb, lc := dfs(node.Left)
        ra, rb, rc := dfs(node.Right)
        a = lc + rc + 1
        b = min(a, min(la+rb, ra+lb))
        c = min(a, lb+rb)
        return
    }
    _, ans, _ := dfs(root)
    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
复杂度分析

时间复杂度：O(N)O(N)，其中 NN 为树中节点的数量。对于每个节点，我们在常数时间内计算出 a,b,ca,b,c 三个状态变量。
空间复杂度：O(N)O(N)。每次递归调用时，我们需要开辟常数大小的空间存储状态变量的取值，而递归栈的深度等于树的深度，即 O(N)O(N)。

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/binary-tree-cameras/solution/jian-kong-er-cha-shu-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
