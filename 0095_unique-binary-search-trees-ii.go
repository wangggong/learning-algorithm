/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Pair struct {
	First, Second int
}

func generateTrees(n int) []*TreeNode {
	m := make(map[Pair][]*TreeNode)
	p := Pair{First: 1, Second: n}
	return dfs(p, m)
}

func dfs(p Pair, m map[Pair][]*TreeNode) []*TreeNode {
	if tns, ok := m[p]; ok {
		return tns
	}

	tns := make([]*TreeNode, 0)
	defer func() {
		m[p] = tns
	}()

	if p.First > p.Second {
		return tns
	}

	for i := p.First; i <= p.Second; i++ {
		left := dfs(Pair{First: p.First, Second: i - 1}, m)
		right := dfs(Pair{First: i + 1, Second: p.Second}, m)

		if len(left) == 0 {
			left = append(left, nil)
		}

		if len(right) == 0 {
			right = append(right, nil)
		}

		for _, l := range left {
			for _, r := range right {
				tns = append(tns, &TreeNode{Val: i, Left: l, Right: r})
			}
		}
	}

	return tns
}
