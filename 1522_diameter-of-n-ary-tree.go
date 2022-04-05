/*
 * Medium
 *
 * 给定一棵 N 叉树的根节点 root ，计算这棵树的直径长度。
 *
 * N 叉树的直径指的是树中任意两个节点间路径中 最长 路径的长度。这条路径可能经过根节点，也可能不经过根节点。
 *
 * （N 叉树的输入序列以层序遍历的形式给出，每组子节点用 null 分隔）
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：root = [1,null,3,2,4,null,5,6]
 * 输出：3
 * 解释：直径如图中红线所示。
 * 示例 2：
 *
 *
 *
 * 输入：root = [1,null,2,null,3,4,null,5,null,6]
 * 输出：4
 * 示例 3：
 *
 *
 *
 * 输入: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
 * 输出: 7
 *
 *
 * 提示：
 *
 * N 叉树的深度小于或等于 1000 。
 * 节点的总个数在 [0, 10^4] 间。
 * 通过次数1,333
 * 提交次数1,873
 */

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

var G map[*Node][]*Node

func diameter(root *Node) int {
	if root == nil {
		return 0
	}
	G = make(map[*Node][]*Node)
	buildGraph(root)
	visit := make(map[*Node]bool)
	for k := range G {
		visit[k] = false
	}
	var Q []*Node
	curr := root
	Q = append(Q, curr)
	visit[curr] = true
	for len(Q) > 0 {
		curr = Q[0]
		Q = Q[1:]
		for _, n := range G[curr] {
			if visit[n] {
				continue
			}
			Q = append(Q, n)
			visit[n] = true
		}
	}
	for k := range visit {
		visit[k] = false
	}
	level := 0
	Q = append(Q, curr)
	visit[curr] = true
	for len(Q) > 0 {
		level++
		for size := len(Q); size > 0; size-- {
			curr = Q[0]
			Q = Q[1:]
			for _, n := range G[curr] {
				if visit[n] {
					continue
				}
				Q = append(Q, n)
				visit[n] = true
			}
		}
	}
	return level - 1
}

func buildGraph(root *Node) {
	if root == nil {
		return
	}
	for _, c := range root.Children {
		if c == nil {
			continue
		}
		G[root] = append(G[root], c)
		G[c] = append(G[c], root)
		buildGraph(c)
	}
}

