/*
 * @lc app=leetcode.cn id=escape-the-spreading-fire lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2258] 逃离火灾
 *
 * https://leetcode-cn.com/problems/escape-the-spreading-fire/description/
 *
 * algorithms
 * Hard (28.78%)
 * Total Accepted:    1.3K
 * Total Submissions: 4.4K
 * Testcase Example:  '[[0,2,0,0,0,0,0],[0,0,0,2,2,1,0],[0,2,0,0,1,2,0],[0,0,2,2,2,0,2],[0,0,0,0,0,0,0]]'
 *
 * 给你一个下标从 0 开始大小为 m x n 的二维整数数组 grid ，它表示一个网格图。每个格子为下面 3 个值之一：
 *
 *
 * 0 表示草地。
 * 1 表示着火的格子。
 * 2 表示一座墙，你跟火都不能通过这个格子。
 *
 *
 * 一开始你在最左上角的格子 (0, 0) ，你想要到达最右下角的安全屋格子 (m - 1, n - 1) 。每一分钟，你可以移动到 相邻
 * 的草地格子。每次你移动 之后 ，着火的格子会扩散到所有不是墙的 相邻 格子。
 *
 * 请你返回你在初始位置可以停留的 最多 分钟数，且停留完这段时间后你还能安全到达安全屋。如果无法实现，请你返回 -1 。如果不管你在初始位置停留多久，你
 * 总是 能到达安全屋，请你返回 10^9 。
 *
 * 注意，如果你到达安全屋后，火马上到了安全屋，这视为你能够安全到达安全屋。
 *
 * 如果两个格子有共同边，那么它们为 相邻 格子。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：grid =
 * [[0,2,0,0,0,0,0],[0,0,0,2,2,1,0],[0,2,0,0,1,2,0],[0,0,2,2,2,0,2],[0,0,0,0,0,0,0]]
 * 输出：3
 * 解释：上图展示了你在初始位置停留 3 分钟后的情形。
 * 你仍然可以安全到达安全屋。
 * 停留超过 3 分钟会让你无法安全到达安全屋。
 *
 * 示例 2：
 *
 *
 *
 * 输入：grid = [[0,0,0,0],[0,1,2,0],[0,2,0,0]]
 * 输出：-1
 * 解释：上图展示了你马上开始朝安全屋移动的情形。
 * 火会蔓延到你可以移动的所有格子，所以无法安全到达安全屋。
 * 所以返回 -1 。
 *
 *
 * 示例 3：
 *
 *
 *
 * 输入：grid = [[0,0,0],[2,2,0],[1,2,0]]
 * 输出：1000000000
 * 解释：上图展示了初始网格图。
 * 注意，由于火被墙围了起来，所以无论如何你都能安全到达安全屋。
 * 所以返回 10^9 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 2 <= m, n <= 300
 * 4 <= m * n <= 2 * 10^4
 * grid[i][j] 是 0 ，1 或者 2 。
 * grid[0][0] == grid[m - 1][n - 1] == 0
 *
 *
 */

const maxn int = 1e9

const (
	grass = iota
	fire
	wall
)

var directions = [][]int{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
}

var n, m int
var times [][]int

func maximumMinutes(G [][]int) int {
	n, m = len(G), len(G[0])
	times = make([][]int, n)
	for i := 0; i < n; i++ {
		times[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			times[i][j] = maxn
		}
	}
	spreadFire(G)
	p, q := 0, m*n+1
	for p < q {
		mid := (p + q) >> 1
		if canEscape(G, mid) {
			p = mid + 1
		} else {
			q = mid
		}
	}
	if p > m*n {
		return maxn
	}
	return p - 1
}

func spreadFire(G [][]int) {
	var Q [][]int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if G[i][j] != fire {
				continue
			}
			Q = append(Q, []int{i, j, 0})
			times[i][j] = 0
		}
	}
	for len(Q) > 0 {
		q := Q[0]
		Q = Q[1:]
		for _, d := range directions {
			ni, nj, nt := q[0]+d[0], q[1]+d[1], q[2]+1
			if ni < 0 || ni >= n || nj < 0 || nj >= m || G[ni][nj] != grass || times[ni][nj] <= nt {
				continue
			}
			Q = append(Q, []int{ni, nj, nt})
			times[ni][nj] = nt
		}
	}
	return
}

func canEscape(G [][]int, k int) bool {
	var Q [][]int
	Q = append(Q, []int{0, 0, k})
	visit := make([][]bool, n)
	for i := range visit {
		visit[i] = make([]bool, m)
	}
	visit[0][0] = true
	for len(Q) > 0 {
		q := Q[0]
		Q = Q[1:]
		if q[0] == n-1 && q[1] == m-1 && q[2] <= times[n-1][m-1] {
			return true
		}
		for _, d := range directions {
			ni, nj, nt := q[0]+d[0], q[1]+d[1], q[2]+1
			if ni == n-1 && nj == m-1 && nt <= times[ni][nj] {
				return true
			}
			if ni < 0 || ni >= n || nj < 0 || nj >= m || G[ni][nj] != grass || times[ni][nj] <= nt || visit[ni][nj] {
				continue
			}
			Q = append(Q, []int{ni, nj, nt})
			visit[ni][nj] = true
		}
	}
	return false
}
