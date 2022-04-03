/*
 * Hard
 *
 * 给你一个二维矩阵 matrix ，你需要处理下面两种类型的若干次查询：
 *
 * 更新：更新 matrix 中某个单元的值。
 * 求和：计算矩阵 matrix 中某一矩形区域元素的 和 ，该区域由 左上角 (row1, col1) 和 右下角 (row2, col2) 界定。
 * 实现 NumMatrix 类：
 *
 * NumMatrix(int[][] matrix) 用整数矩阵 matrix 初始化对象。
 * void update(int row, int col, int val) 更新 matrix[row][col] 的值到 val 。
 * int sumRegion(int row1, int col1, int row2, int col2) 返回矩阵 matrix 中指定矩形区域元素的 和 ，该区域由 左上角 (row1, col1) 和 右下角 (row2, col2) 界定。
 *
 * 示例 1：
 *
 *
 * 输入
 * ["NumMatrix", "sumRegion", "update", "sumRegion"]
 * [[[[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]], [2, 1, 4, 3], [3, 2, 2], [2, 1, 4, 3]]
 * 输出
 * [null, 8, null, 10]
 *
 * 解释
 * NumMatrix numMatrix = new NumMatrix([[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]);
 * numMatrix.sumRegion(2, 1, 4, 3); // 返回 8 (即, 左侧红色矩形的和)
 * numMatrix.update(3, 2, 2);       // 矩阵从左图变为右图
 * numMatrix.sumRegion(2, 1, 4, 3); // 返回 10 (即，右侧红色矩形的和)
 *
 * 提示：
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 200
 * -1e5 <= matrix[i][j] <= 1e5
 * 0 <= row < m
 * 0 <= col < n
 * -1e5 <= val <= 1e5
 * 0 <= row1 <= row2 < m
 * 0 <= col1 <= col2 < n
 * 最多调用1e4 次 sumRegion 和 update 方法
 */

type NumArray struct {
	Size int
	Nums []int
	Tree []int
}

func lowBit(x int) int { return x & -x }

func (na *NumArray) query(x int) int {
	// assert x >= 0 && x < na.Size
	ans := 0
	for i := x; i > 0; i -= lowBit(i) {
		ans += na.Tree[i]
	}
	return ans
}

func (na *NumArray) add(x, v int) {
	// assert x >= 0 && x < len(na.Tree)
	for i := x; i <= na.Size; i += lowBit(i) {
		na.Tree[i] += v
	}
}

type NumMatrix struct {
	Arrays []NumArray
}

func Constructor(matrix [][]int) NumMatrix {
	n, m := len(matrix), len(matrix[0])
	arrays := make([]NumArray, 0, n)
	for _, row := range matrix {
		tree := make([]int, m+1)
		na := NumArray{
			Size: m,
			Nums: row,
			Tree: tree,
		}
		for i := 0; i < m; i++ {
			na.add(i+1, row[i])
		}
		arrays = append(arrays, na)
	}
	return NumMatrix{Arrays: arrays}
}

func (nm *NumMatrix) Update(row int, col int, val int) {
	// assert row >= 0 && row <= len(nm.Arrays);
	// assert col >= 0 && col <= nm.Arrays[row].Size;
	na := nm.Arrays[row]
	na.add(col+1, val-na.Nums[col])
	na.Nums[col] = val
}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// assert row1 >= 0 && row1 < row2 && row2 <= len(nm.Arrays);
	// assert col1 >= 0 && col1 < col2 && col2 <= nm.Arrays[row].Size;
	ans := 0
	for i := row1; i <= row2; i++ {
		na := nm.Arrays[i]
		ans += na.query(col2+1) - na.query(col1)
	}
	return ans
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * obj.Update(row,col,val);
 * param_2 := obj.SumRegion(row1,col1,row2,col2);
 */
