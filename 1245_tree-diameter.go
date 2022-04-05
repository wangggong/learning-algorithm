/*
 * Medium
 *
 * 给你这棵「无向树」，请你测算并返回它的「直径」：这棵树上最长简单路径的 边数。
 *
 * 我们用一个由所有「边」组成的数组 edges 来表示一棵无向树，其中 edges[i] = [u, v] 表示节点 u 和 v 之间的双向边。
 *
 * 树上的节点都已经用 {0, 1, ..., edges.length} 中的数做了标记，每个节点上的标记都是独一无二的。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：edges = [[0,1],[0,2]]
 * 输出：2
 * 解释：
 * 这棵树上最长的路径是 1 - 0 - 2，边数为 2。
 * 示例 2：
 *
 *
 *
 * 输入：edges = [[0,1],[1,2],[2,3],[1,4],[4,5]]
 * 输出：4
 * 解释：
 * 这棵树上最长的路径是 3 - 2 - 1 - 4 - 5，边数为 4。
 *
 *
 * 提示：
 *
 * 0 <= edges.length < 10^4
 * edges[i][0] != edges[i][1]
 * 0 <= edges[i][j] <= edges.length
 * edges 会形成一棵无向树
 * 通过次数4,169
 * 提交次数8,156
 */

// 参考:
// - https://leetcode-cn.com/problems/diameter-of-n-ary-tree/solution/c-python3-2ci-bfsji-yi-hua-shu-de-zhi-ji-3rsl/
// - https://oi-wiki.org/graph/tree-diameter/
//
// 二次 BFS:
// 1. 第一次 BFS, 随便找一节点作为起始点, 记录最后访问到的节点;
// 2. 第二次 BFS, 以上次访问到的节点为起始点, 记录总长度;
func treeDiameter(edges [][]int) int {
	G := make(map[int][]int)
	m := make(map[int]struct{})
	for _, e := range edges {
		u, v := e[0], e[1]
		G[u] = append(G[u], v)
		G[v] = append(G[v], u)
		m[u] = struct{}{}
		m[v] = struct{}{}
	}
	n := len(m)
	if n == 0 {
		return 0
	}
	visit := make([]bool, n)
	var Q []int
	Q = append(Q, 0)
	visit[0] = true
	curr := 0
	for len(Q) > 0 {
		curr = Q[0]
		Q = Q[1:]
		for _, node := range G[curr] {
			if visit[node] {
				continue
			}
			Q = append(Q, node)
			visit[node] = true
		}
	}
	for i := range visit {
		visit[i] = false
	}
	level := 0
	Q = append(Q, curr)
	visit[curr] = true
	for len(Q) > 0 {
		level++
		for size := len(Q); size > 0; size-- {
			curr = Q[0]
			Q = Q[1:]
			for _, node := range G[curr] {
				if visit[node] {
					continue
				}
				Q = append(Q, node)
				visit[node] = true
			}
		}
	}
	return level - 1
}
