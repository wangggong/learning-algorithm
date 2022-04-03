/*
 * @lc app=leetcode.cn id=range-sum-query-2d-immutable lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [304] 二维区域和检索 - 矩阵不可变
 *
 * https://leetcode-cn.com/problems/range-sum-query-2d-immutable/description/
 *
 * algorithms
 * Medium (56.64%)
 * Total Accepted:    75.2K
 * Total Submissions: 132.8K
 * Testcase Example:  '["NumMatrix","sumRegion","sumRegion","sumRegion"]\n' +
  '[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]'
 *
 * 给定一个二维矩阵 matrix，以下类型的多个请求：
 *
 *
 * 计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，右下角 为 (row2, col2) 。
 *
 *
 * 实现 NumMatrix 类：
 *
 *
 * NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
 * int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1)
 * 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入:
 * ["NumMatrix","sumRegion","sumRegion","sumRegion"]
 *
 * [[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]
 * 输出:
 * [null, 8, 11, 12]
 *
 * 解释:
 * NumMatrix numMatrix = new
 * NumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]);
 * numMatrix.sumRegion(2, 1, 4, 3); // return 8 (红色矩形框的元素总和)
 * numMatrix.sumRegion(1, 1, 2, 2); // return 11 (绿色矩形框的元素总和)
 * numMatrix.sumRegion(1, 2, 2, 4); // return 12 (蓝色矩形框的元素总和)
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 200
 * -10^5 <= matrix[i][j] <= 10^5
 * 0 <= row1 <= row2 < m
 * 0 <= col1 <= col2 < n
 * 最多调用 10^4 次 sumRegion 方法
 *
 *
*/
type NumMatrix struct {
	Matrix    [][]int
	PrefixSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	// assert len(matrix) > 0 && len(matrix) <= 200;
	// assert len(matrix[0]) > 0 && len(matrix[0]) <= 200;
	n, m := len(matrix), len(matrix[0])
	prefixSum := make([][]int, 0, n)
	for _, row := range matrix {
		ps := make([]int, m+1)
		for i, c := range row {
			ps[i+1] = ps[i] + c
		}
		prefixSum = append(prefixSum, ps)
	}
	return NumMatrix{Matrix: matrix, PrefixSum: prefixSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// assert row1 >= 0 && row1 < row2 && row2 < len(this.Matrix);
	// assert col1 >= 0 && col1 < col2 && col2 < len(this.Matrix[0]);
	ans := 0
	for i := row1; i <= row2; i++ {
		// fmt.Printf("%v %v %v %v %v\n", this.PrefixSum, row1, col1, row2, col2)
		ans += this.PrefixSum[i][col2+1] - this.PrefixSum[i][col1]
	}
	return ans
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
