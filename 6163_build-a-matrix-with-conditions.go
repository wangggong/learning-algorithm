/*
 * @lc app=leetcode.cn id=build-a-matrix-with-conditions lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6163] 给定条件下构造矩阵
 *
 * https://leetcode.cn/problems/build-a-matrix-with-conditions/description/
 *
 * algorithms
 * Hard (51.30%)
 * Total Accepted:    2.9K
 * Total Submissions: 5.6K
 * Testcase Example:  '3\n[[1,2],[3,2]]\n[[2,1],[3,2]]'
 *
 * 给你一个 正 整数 k ，同时给你：
 *
 *
 * 一个大小为 n 的二维整数数组 rowConditions ，其中 rowConditions[i] = [abovei, belowi] 和
 * 一个大小为 m 的二维整数数组 colConditions ，其中 colConditions[i] = [lefti, righti] 。
 *
 *
 * 两个数组里的整数都是 1 到 k 之间的数字。
 *
 * 你需要构造一个 k x k 的矩阵，1 到 k 每个数字需要 恰好出现一次 。剩余的数字都是 0 。
 *
 * 矩阵还需要满足以下条件：
 *
 *
 * 对于所有 0 到 n - 1 之间的下标 i ，数字 abovei 所在的 行 必须在数字 belowi 所在行的上面。
 * 对于所有 0 到 m - 1 之间的下标 i ，数字 lefti 所在的 列 必须在数字 righti 所在列的左边。
 *
 *
 * 返回满足上述要求的 任意 矩阵。如果不存在答案，返回一个空的矩阵。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：k = 3, rowConditions = [[1,2],[3,2]], colConditions = [[2,1],[3,2]]
 * 输出：[[3,0,0],[0,0,1],[0,2,0]]
 * 解释：上图为一个符合所有条件的矩阵。
 * 行要求如下：
 * - 数字 1 在第 1 行，数字 2 在第 2 行，1 在 2 的上面。
 * - 数字 3 在第 0 行，数字 2 在第 2 行，3 在 2 的上面。
 * 列要求如下：
 * - 数字 2 在第 1 列，数字 1 在第 2 列，2 在 1 的左边。
 * - 数字 3 在第 0 列，数字 2 在第 1 列，3 在 2 的左边。
 * 注意，可能有多种正确的答案。
 *
 *
 * 示例 2：
 *
 * 输入：k = 3, rowConditions = [[1,2],[2,3],[3,1],[2,3]], colConditions = [[2,1]]
 * 输出：[]
 * 解释：由前两个条件可以得到 3 在 1 的下面，但第三个条件是 3 在 1 的上面。
 * 没有符合条件的矩阵存在，所以我们返回空矩阵。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= k <= 400
 * 1 <= rowConditions.length, colConditions.length <= 10^4
 * rowConditions[i].length == colConditions[i].length == 2
 * 1 <= abovei, belowi, lefti, righti <= k
 * abovei != belowi
 * lefti != righti
 *
 *
 */
func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	rows, ok := topo(k, rowConditions)
	// fmt.Println(rows)
	if !ok {
		return nil
	}
	cols, ok := topo(k, colConditions)
	// fmt.Println(cols)
	if !ok {
		return nil
	}
	ans := make([][]int, k)
	for i := range ans {
		ans[i] = make([]int, k)
	}
	rr := make([]int, k+1)
	for i, r := range rows {
		rr[r] = i
	}
	for i, c := range cols {
		ans[rr[c]][i] = c
	}
	return ans
}

func topo(n int, cond [][]int) ([]int, bool) {
	G := make([][]int, n+1)
	inorder := make([]int, n+1)
	for _, c := range cond {
		G[c[0]] = append(G[c[0]], c[1])
		inorder[c[1]]++
	}
	var Q []int
	for i := 1; i <= n; i++ {
		if inorder[i] == 0 {
			Q = append(Q, i)
		}
	}
	var ans []int
	for len(Q) > 0 {
		q := Q[0]
		Q = Q[1:]
		ans = append(ans, q)
		for _, k := range G[q] {
			inorder[k]--
			if inorder[k] == 0 {
				Q = append(Q, k)
			}
		}
	}
	return ans, len(ans) == n
}
