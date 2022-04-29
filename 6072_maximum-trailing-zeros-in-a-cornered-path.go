/*
 * Medium
 *
 * 给你一个二维整数数组 grid ，大小为 m x n，其中每个单元格都含一个正整数。
 *
 * 转角路径 定义为：包含至多一个弯的一组相邻单元。具体而言，路径应该完全 向水平方向 或者 向竖直方向 移动过弯（如果存在弯），而不能访问之前访问过的单元格。在过弯之后，路径应当完全朝 另一个 方向行进：如果之前是向水平方向，那么就应该变为向竖直方向；反之亦然。当然，同样不能访问之前已经访问过的单元格。
 *
 * 一条路径的 乘积 定义为：路径上所有值的乘积。
 *
 * 请你从 grid 中找出一条乘积中尾随零数目最多的转角路径，并返回该路径中尾随零的数目。
 *
 * 注意：
 *
 * 水平 移动是指向左或右移动。
 * 竖直 移动是指向上或下移动。
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：grid = [[23,17,15,3,20],[8,1,20,27,11],[9,4,6,2,21],[40,9,1,10,6],[22,7,4,5,3]]
 * 输出：3
 * 解释：左侧的图展示了一条有效的转角路径。
 * 其乘积为 15 * 20 * 6 * 1 * 10 = 18000 ，共计 3 个尾随零。
 * 可以证明在这条转角路径的乘积中尾随零数目最多。
 *
 * 中间的图不是一条有效的转角路径，因为它有不止一个弯。
 * 右侧的图也不是一条有效的转角路径，因为它需要重复访问已经访问过的单元格。
 * 示例 2：
 *
 *
 *
 * 输入：grid = [[4,3,2],[7,6,1],[8,8,8]]
 * 输出：0
 * 解释：网格如上图所示。
 * 不存在乘积含尾随零的转角路径。
 *
 *
 * 提示：
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 105
 * 1 <= m * n <= 105
 * 1 <= grid[i][j] <= 1000
 * 通过次数1,224
 * 提交次数5,750
 */

/*
 * // 参考: https://leetcode-cn.com/problems/maximum-trailing-zeros-in-a-cornered-path/solution/by-tsreaper-ukq5/
 * // 现场没写出来, 后面看着题解推出来的. 丢人啊!
 * func maxTrailingZeros(G [][]int) int {
 * 	n, m := len(G), len(G[0])
 * 	row := make([][][2]int, n+1)
 * 	for i := range row {
 * 		row[i] = make([][2]int, m+1)
 * 	}
 * 	col := make([][][2]int, n+1)
 * 	for i := range col {
 * 		col[i] = make([][2]int, m+1)
 * 	}
 * 	var ans int
 * 	for i := 1; i <= n; i++ {
 * 		for j := 1; j <= m; j++ {
 * 			tmp := [2]int{row[i][j-1][0], row[i][j-1][1]}
 * 			for v := G[i-1][j-1]; v%2 == 0; v /= 2 {
 * 				tmp[0]++
 * 			}
 * 			for v := G[i-1][j-1]; v%5 == 0; v /= 5 {
 * 				tmp[1]++
 * 			}
 * 			row[i][j] = tmp
 * 			tmp = [2]int{col[i-1][j][0], col[i-1][j][1]}
 * 			for v := G[i-1][j-1]; v%2 == 0; v /= 2 {
 * 				tmp[0]++
 * 			}
 * 			for v := G[i-1][j-1]; v%5 == 0; v /= 5 {
 * 				tmp[1]++
 * 			}
 * 			col[i][j] = tmp
 * 		}
 * 	}
 * 	for i := 1; i <= n; i++ {
 * 		for j := 1; j <= m; j++ {
 * 			tmp := [2]int{row[i][j][0] + col[i-1][j][0], row[i][j][1] + col[i-1][j][1]}
 * 			ans = max(ans, min(tmp[0], tmp[1]))
 * 			tmp = [2]int{row[i][j][0] + col[n][j][0] - col[i][j][0], row[i][j][1] + col[n][j][1] - col[i][j][1]}
 * 			ans = max(ans, min(tmp[0], tmp[1]))
 * 			tmp = [2]int{row[i][m][0] - row[i][j][0] + col[i][j][0], row[i][m][1] - row[i][j][1] + col[i][j][1]}
 * 			ans = max(ans, min(tmp[0], tmp[1]))
 * 			tmp = [2]int{row[i][m][0] - row[i][j][0] + col[n][j][0] - col[i-1][j][0], row[i][m][1] - row[i][j][1] + col[n][j][1] - col[i-1][j][1]}
 * 			ans = max(ans, min(tmp[0], tmp[1]))
 * 		}
 * 	}
 * 	return ans
 * }
 *
 * func min(x, y int) int {
 * 	if x < y {
 * 		return x
 * 	}
 * 	return y
 * }
 *
 * func max(x, y int) int {
 * 	if x > y {
 * 		return x
 * 	}
 * 	return y
 * }
 */

func maxTrailingZeros(G [][]int) int {
	n, m := len(G), len(G[0])
	R2 := make([][]int, n+1)
	D2 := make([][]int, n+1)
	R5 := make([][]int, n+1)
	D5 := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		R2[i] = make([]int, m+1)
		D2[i] = make([]int, m+1)
		R5[i] = make([]int, m+1)
		D5[i] = make([]int, m+1)
	}
	// 从这里面能吸取的经验教训就是遇到这种问题尽量保证坐标对齐.
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			cnt2, cnt5 := 0, 0
			for v := G[i-1][j-1]; v%2 == 0; v /= 2 {
				cnt2++
			}
			for v := G[i-1][j-1]; v%5 == 0; v /= 5 {
				cnt5++
			}
			R2[i][j] = R2[i][j-1] + cnt2
			D2[i][j] = D2[i-1][j] + cnt2
			R5[i][j] = R5[i][j-1] + cnt5
			D5[i][j] = D5[i-1][j] + cnt5
		}
	}
	ans := 0
	// 另外傻逼! 这里假定先横后竖的话有四种可能情况, 没枚举全能怪谁啊!
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			ans = max(ans, min(R2[i][j]+D2[i-1][j], R5[i][j]+D5[i-1][j]))
			ans = max(ans, min(R2[i][j]+D2[n][j]-D2[i][j], R5[i][j]+D5[n][j]-D5[i][j]))
			ans = max(ans, min(R2[i][m]-R2[i][j]+D2[i][j], R5[i][m]-R5[i][j]+D5[i][j]))
			ans = max(ans, min(R2[i][m]-R2[i][j]+D2[n][j]-D2[i-1][j], R5[i][m]-R5[i][j]+D5[n][j]-D5[i-1][j]))
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
