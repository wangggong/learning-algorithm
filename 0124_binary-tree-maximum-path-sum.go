/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
import (
	"fmt"
)
*/

const MIN = -1e3 - 10

type S struct {
	Left, Right int
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dfs(R *TreeNode, m map[*TreeNode]S) int {
	if R == nil {
		return 0
	}

	l := dfs(R.Left, m) + R.Val
	r := dfs(R.Right, m) + R.Val
	s := S{
		Left:  l,
		Right: r,
	}

	m[R] = s

	ans := max(max(l, r), R.Val)
	return ans
}

func maxPathSum(R *TreeNode) int {
	m := make(map[*TreeNode]S)
	ans := int(MIN)
	dfs(R, m)

	for k, v := range m {
		if k == nil {
			continue
		}

		// fmt.Printf("R.val = %v, left = %v, right = %v\n", k.Val, v.Left, v.Right)
		ans = max(ans, v.Left)
		ans = max(ans, v.Right)
		ans = max(ans, v.Left+v.Right-k.Val)
		ans = max(ans, k.Val)
	}

	// fmt.Println()
	return ans
}
