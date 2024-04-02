/*
 * Medium
 *
 * 想象一下你是个城市基建规划者，地图上有 n 座城市，它们按以 1 到 n 的次序编号。
 *
 * 给你整数 n 和一个数组 conections，其中 connections[i] = [xi, yi, costi] 表示将城市 xi 和城市 yi 连接所要的costi（连接是双向的）。
 *
 * 返回连接所有城市的最低成本，每对城市之间至少有一条路径。如果无法连接所有 n 个城市，返回 -1
 *
 * 该 最小成本 应该是所用全部连接成本的总和。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：n = 3, conections = [[1,2,5],[1,3,6],[2,3,1]]
 * 输出：6
 * 解释：选出任意 2 条边都可以连接所有城市，我们从中选取成本最小的 2 条。
 * 示例 2：
 *
 *
 *
 * 输入：n = 4, conections = [[1,2,3],[3,4,4]]
 * 输出：-1
 * 解释：即使连通所有的边，也无法连接所有城市。
 *
 *
 * 提示：
 *
 * 1 <= n <= 1e4
 * 1 <= connections.length <= 1e4
 * connections[i].length == 3
 * 1 <= xi, yi <= n
 * xi != yi
 * 0 <= costi <= 105
 * 通过次数7,763
 * 提交次数14,147
 */

// cnm 就这么简单. 按边权排序搞个并查集整一波就能通...
//
// Kruskal 模版题.
import "sort"

type UnionFind struct {
	Fa   []int
	Size []int
}

func (uf UnionFind) find(index int) int {
	for uf.Fa[index] != index {
		index = uf.Fa[index]
	}
	return index
}

func (uf UnionFind) isConnect(p, q int) bool {
	fp, fq := uf.find(p), uf.find(q)
	return fp == fq
}

func (uf UnionFind) connect(p, q int) {
	fp, fq := uf.find(p), uf.find(q)
	if uf.Size[fp] > uf.Size[fq] {
		uf.Fa[fq] = fp
		uf.Size[fp] += fq
	} else {
		uf.Fa[fp] = fq
		uf.Size[fq] += fp
	}
}

func NewUnionFind(n int) *UnionFind {
	fa := make([]int, 0, n+1)
	size := make([]int, 0, n+1)
	for i := 0; i <= n; i++ {
		fa = append(fa, i)
		size = append(size, 1)
	}
	uf := UnionFind{
		Fa:   fa,
		Size: size,
	}
	return &uf
}

func minimumCost(n int, connections [][]int) int {
	sort.Slice(connections, func(p, q int) bool {
		if connections[p][2] != connections[q][2] {
			return connections[p][2] < connections[q][2]
		}
		return connections[p][0] < connections[q][0]
	})
	ans := 0
	uf := NewUnionFind(n)
	for _, c := range connections {
		p, q, cost := c[0], c[1], c[2]
		if uf.isConnect(p, q) {
			continue
		}
		uf.connect(p, q)
		ans += cost
	}
	fa := uf.find(1)
	for i := 1; i <= n; i++ {
		if fa != uf.find(i) {
			return -1
		}
	}
	return ans
}
