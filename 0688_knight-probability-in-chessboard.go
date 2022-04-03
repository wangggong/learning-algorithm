/*
 * @lc app=leetcode.cn id=knight-probability-in-chessboard lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [688] 骑士在棋盘上的概率
 *
 * https://leetcode-cn.com/problems/knight-probability-in-chessboard/description/
 *
 * algorithms
 * Medium (51.48%)
 * Total Accepted:    14.3K
 * Total Submissions: 25.7K
 * Testcase Example:  '3\n2\n0\n0'
 *
 * 在一个 n x n 的国际象棋棋盘上，一个骑士从单元格 (row, column) 开始，并尝试进行 k 次移动。行和列是 从 0 开始
 * 的，所以左上单元格是 (0,0) ，右下单元格是 (n - 1, n - 1) 。
 *
 * 象棋骑士有8种可能的走法，如下图所示。每次移动在基本方向上是两个单元格，然后在正交方向上是一个单元格。
 *
 *
 *
 * 每次骑士要移动时，它都会随机从8种可能的移动中选择一种(即使棋子会离开棋盘)，然后移动到那里。
 *
 * 骑士继续移动，直到它走了 k 步或离开了棋盘。
 *
 * 返回 骑士在棋盘停止移动后仍留在棋盘上的概率 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: n = 3, k = 2, row = 0, column = 0
 * 输出: 0.0625
 * 解释: 有两步(到(1,2)，(2,1))可以让骑士留在棋盘上。
 * 在每一个位置上，也有两种移动可以让骑士留在棋盘上。
 * 骑士留在棋盘上的总概率是0.0625。
 *
 *
 * 示例 2：
 *
 *
 * 输入: n = 1, k = 0, row = 0, column = 0
 * 输出: 1.00000
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= n <= 25
 * 0 <= k <= 100
 * 0 <= row, column <= n
 *
 *
 */

// 注意以下几点:
// 1. 不能无脑模拟, 数据量会溢出, OOM
// 2. 不能攒到最后再算, 总可能数会溢出, 变成负数了
// 坑啊...

const stepCount = 8

type Pair struct{ x, y int }

var directions = [stepCount][2]int{
	{1, 2}, {2, 1},
	{-1, 2}, {-2, 1},
	{1, -2}, {2, -1},
	{-1, -2}, {-2, -1},
}

var ms []map[Pair]float64

func knightProbability(n int, k int, row int, column int) float64 {
	tot := 1
	ms = nil
	for i := 0; i < k; i++ {
		tot *= stepCount
		ms = append(ms, make(map[Pair]float64))
	}
	ans := dfs(n, k, Pair{row, column})
	// return float64(ans) / float64(tot)
	return ans
}

/*
 * // 这个方案为啥不行? 因为溢出了 [Doge]
 * func dfs(n, k int, p Pair) int {
 * 	// assert valid(ps[i]);
 * 	// assert ms[i] != nil;
 * 	if k == 0 {
 * 		return 1
 * 	}
 * 	if v, ok := ms[k-1][p]; ok {
 * 		return v
 * 	}
 * 	var ans int
 * 	defer func() { ms[k-1][p] = ans }()
 * 	for _, d := range directions {
 * 		np := Pair{p.x + d[0], p.y + d[1]}
 * 		if !valid(n, np) {
 * 			continue
 * 		}
 * 		ans += dfs(n, k-1, np)
 * 	}
 * 	return ans
 * }
 */

func dfs(n, k int, p Pair) float64 {
	if k == 0 {
		return 1
	}
	if v, ok := ms[k-1][p]; ok {
		return v
	}
	var ans float64
	defer func() { ms[k-1][p] = ans }()
	for _, d := range directions {
		np := Pair{p.x + d[0], p.y + d[1]}
		if !valid(n, np) {
			continue
		}
		ans += dfs(n, k-1, np) / 8
	}
	return ans
}

func valid(n int, p Pair) bool {
	return 0 <= p.x && p.x < n && 0 <= p.y && p.y < n
}

