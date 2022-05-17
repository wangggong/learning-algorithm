/*
 * @lc app=leetcode.cn id=number-of-provinces lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [547] 省份数量
 *
 * https://leetcode-cn.com/problems/number-of-provinces/description/
 *
 * algorithms
 * Medium (62.05%)
 * Total Accepted:    210.3K
 * Total Submissions: 338.9K
 * Testcase Example:  '[[1,1,0],[1,1,0],[0,0,1]]'
 *
 *
 *
 * 有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c
 * 间接相连。
 *
 * 省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。
 *
 * 给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而
 * isConnected[i][j] = 0 表示二者不直接相连。
 *
 * 返回矩阵中 省份 的数量。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：isConnected = [[1,0,0],[0,1,0],[0,0,1]]
 * 输出：3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * n == isConnected.length
 * n == isConnected[i].length
 * isConnected[i][j] 为 1 或 0
 * isConnected[i][i] == 1
 * isConnected[i][j] == isConnected[j][i]
 *
 *
 *
 *
 */
const maxn = 200

var fa [maxn + 10]int

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

func initUF() {
	for i := range fa {
		fa[i] = i
	}
	return
}

func findCircleNum(isConnected [][]int) int {
	initUF()
	n := len(isConnected)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] != 0 {
				merge(i, j)
			}
		}
	}
	m := make(map[int]struct{})
	for i := 0; i < n; i++ {
		k := query(i)
		if _, ok := m[k]; ok {
			continue
		}
		m[k] = struct{}{}
	}
	return len(m)
}
