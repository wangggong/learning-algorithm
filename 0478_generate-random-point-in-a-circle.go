/*
 * @lc app=leetcode.cn id=generate-random-point-in-a-circle lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [478] 在圆内随机生成点
 *
 * https://leetcode.cn/problems/generate-random-point-in-a-circle/description/
 *
 * algorithms
 * Medium (44.80%)
 * Total Accepted:    11.7K
 * Total Submissions: 25.8K
 * Testcase Example:  '["Solution","randPoint","randPoint","randPoint"]\n[[1.0,0.0,0.0],[],[],[]]'
 *
 * 给定圆的半径和圆心的位置，实现函数 randPoint ，在圆中产生均匀随机点。
 *
 * 实现 Solution 类:
 *
 *
 * Solution(double radius, double x_center, double y_center) 用圆的半径 radius
 * 和圆心的位置 (x_center, y_center) 初始化对象
 * randPoint() 返回圆内的一个随机点。圆周上的一点被认为在圆内。答案作为数组返回 [x, y] 。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入:
 * ["Solution","randPoint","randPoint","randPoint"]
 * [[1.0, 0.0, 0.0], [], [], []]
 * 输出: [null, [-0.02493, -0.38077], [0.82314, 0.38945], [0.36572, 0.17248]]
 * 解释:
 * Solution solution = new Solution(1.0, 0.0, 0.0);
 * solution.randPoint ();//返回[-0.02493，-0.38077]
 * solution.randPoint ();//返回[0.82314,0.38945]
 * solution.randPoint ();//返回[0.36572,0.17248]
 *
 *
 *
 * 提示：
 *
 *
 * 0 < radius <= 10^8
 * -10^7 <= x_center, y_center <= 10^7
 * randPoint 最多被调用 3 * 10^4 次
 *
 *
 */
import "math/rand"

type Solution struct {
	X, Y, R float64
}

func Constructor(r float64, x float64, y float64) Solution {
	return Solution{
		X: x,
		Y: y,
		R: r,
	}
}

func (s *Solution) RandPoint() []float64 {
	for true {
		x, y := (s.X-s.R)+rand.Float64()*s.R*2, (s.Y-s.R)+rand.Float64()*s.R*2
		if (x-s.X)*(x-s.X)+(y-s.Y)*(y-s.Y) > s.R*s.R {
			continue
		}
		return []float64{x, y}
	}
	// unreachable
	return nil
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
