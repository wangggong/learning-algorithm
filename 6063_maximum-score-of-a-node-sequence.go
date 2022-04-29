/*
 * Hard
 *
 * 给你一个 n 个节点的 无向图 ，节点编号为 0 到 n - 1 。
 *
 * 给你一个下标从 0 开始的整数数组 scores ，其中 scores[i] 是第 i 个节点的分数。同时给你一个二维整数数组 edges ，其中 edges[i] = [ai, bi] ，表示节点 ai 和 bi 之间有一条 无向 边。
 *
 * 一个合法的节点序列如果满足以下条件，我们称它是 合法的 ：
 *
 * 序列中每 相邻 节点之间有边相连。
 * 序列中没有节点出现超过一次。
 * 节点序列的分数定义为序列中节点分数之 和 。
 *
 * 请你返回一个长度为 4 的合法节点序列的最大分数。如果不存在这样的序列，请你返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：scores = [5,2,9,8,4], edges = [[0,1],[1,2],[2,3],[0,2],[1,3],[2,4]]
 * 输出：24
 * 解释：上图为输入的图，节点序列为 [0,1,2,3] 。
 * 节点序列的分数为 5 + 2 + 9 + 8 = 24 。
 * 观察可知，没有其他节点序列得分和超过 24 。
 * 注意节点序列 [3,1,2,0] 和 [1,,2,3] 也是合法的，且分数为 24 。
 * 序列 [0,3,2,4] 不是合法的，因为没有边连接节点 0 和 3 。
 * 示例 2：
 *
 *
 *
 * 输入：scores = [9,20,6,4,11,12], edges = [[0,3],[5,3],[2,4],[1,3]]
 * 输出：-1
 * 解释：上图为输入的图。
 * 没有长度为 4 的合法序列，所以我们返回 -1 。
 *
 *
 * 提示：
 *
 * n == scores.length
 * 4 <= n <= 5 * 104
 * 1 <= scores[i] <= 108
 * 0 <= edges.length <= 5 * 104
 * edges[i].length == 2
 * 0 <= ai, bi <= n - 1
 * ai != bi
 * 不会有重边。
 * 通过次数566
 * 提交次数3,935
 */

type Info [2]int

func maximumScore(scores []int, edges [][]int) int {
	n := len(scores)
	G := make([][3]Info, n)
	for _, e := range edges {
		s, t := e[0], e[1]
		update(scores, G, s, t)
		update(scores, G, t, s)
	}
	ans := -1
	for _, e := range edges {
		s, t := e[0], e[1]
		cur := query(scores, G, s, t)
		ans = max(ans, cur)
	}
	return ans
}

func query(scores []int, G [][3]Info, s, t int) int {
	ans := scores[s] + scores[t]
	cur := -1
	for _, is := range G[s] {
		for _, it := range G[t] {
			if is[0] == t || it[0] == s {
				continue
			}
			if is[1] == 0 || it[1] == 0 {
				continue
			}
			if is[0] == it[0] {
				continue
			}
			cur = max(cur, is[1]+it[1])
		}
	}
	if cur < 0 {
		return -1
	}
	return ans + cur
}

func update(scores []int, G [][3]Info, s, t int) {
	info := Info([2]int{t, scores[t]})
	if scores[t] > G[s][0][1] {
		G[s][2] = G[s][1]
		G[s][1] = G[s][0]
		G[s][0] = info
	} else if scores[t] > G[s][1][1] {
		G[s][2] = G[s][1]
		G[s][1] = info
	} else if scores[t] > G[s][2][1] {
		G[s][2] = info
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
