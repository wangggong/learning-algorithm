/*
 * @lc app=leetcode.cn id=min-cost-to-connect-all-points lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1584] 连接所有点的最小费用
 *
 * https://leetcode-cn.com/problems/min-cost-to-connect-all-points/description/
 *
 * algorithms
 * Medium (66.16%)
 * Total Accepted:    35.5K
 * Total Submissions: 53.6K
 * Testcase Example:  '[[0,0],[2,2],[3,10],[5,2],[7,0]]'
 *
 * 给你一个points 数组，表示 2D 平面上的一些点，其中 points[i] = [xi, yi] 。
 *
 * 连接点 [xi, yi] 和点 [xj, yj] 的费用为它们之间的 曼哈顿距离 ：|xi - xj| + |yi - yj| ，其中 |val| 表示
 * val 的绝对值。
 *
 * 请你返回将所有点连接的最小总费用。只有任意两点之间 有且仅有 一条简单路径时，才认为所有点都已连接。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
 * 输出：20
 * 解释：
 *
 * 我们可以按照上图所示连接所有点得到最小总费用，总费用为 20 。
 * 注意到任意两个点之间只有唯一一条路径互相到达。
 *
 *
 * 示例 2：
 *
 *
 * 输入：points = [[3,12],[-2,5],[-4,1]]
 * 输出：18
 *
 *
 * 示例 3：
 *
 *
 * 输入：points = [[0,0],[1,1],[1,0],[-1,1]]
 * 输出：4
 *
 *
 * 示例 4：
 *
 *
 * 输入：points = [[-1000000,-1000000],[1000000,1000000]]
 * 输出：4000000
 *
 *
 * 示例 5：
 *
 *
 * 输入：points = [[0,0]]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= points.length <= 1000
 * -10^6 <= xi, yi <= 10^6
 * 所有点 (xi, yi) 两两不同。
 *
 *
 */
import "sort"

var fa []int

func query(x int) int {
	if x != fa[x] {
		fa[x] = query(fa[x])
	}
	return fa[x]
}

func merge(p, q int) {
	fa[query(p)] = query(q)
	return
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// Kruskal 模版题. 忘了!
func minCostConnectPoints(points [][]int) int {
	n := len(points)
	fa = make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var edges [][3]int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			u, v := points[i], points[j]
			d := abs(u[0]-v[0]) + abs(u[1]-v[1])
			edges = append(edges, [3]int{i, j, d})
		}
	}
	sort.Slice(edges, func(p, q int) bool { return edges[p][2] < edges[q][2] })
	ans := 0
	for _, e := range edges {
		u, v, d := e[0], e[1], e[2]
		if query(u) == query(v) {
			continue
		}
		merge(u, v)
		ans += d
	}
	return ans
}
