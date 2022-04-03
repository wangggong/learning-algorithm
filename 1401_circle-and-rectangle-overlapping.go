/*
 * @lc app=leetcode.cn id=circle-and-rectangle-overlapping lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1401] 圆和矩形是否有重叠
 *
 * https://leetcode-cn.com/problems/circle-and-rectangle-overlapping/description/
 *
 * algorithms
 * Medium (42.07%)
 * Total Accepted:    4.4K
 * Total Submissions: 10.4K
 * Testcase Example:  '1\n0\n0\n1\n-1\n3\n1'
 *
 * 给你一个以 (radius, x_center, y_center) 表示的圆和一个与坐标轴平行的矩形 (x1, y1, x2, y2)，其中 (x1,
 * y1) 是矩形左下角的坐标，(x2, y2) 是右上角的坐标。
 *
 * 如果圆和矩形有重叠的部分，请你返回 True ，否则返回 False 。
 *
 * 换句话说，请你检测是否 存在 点 (xi, yi) ，它既在圆上也在矩形上（两者都包括点落在边界上的情况）。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：radius = 1, x_center = 0, y_center = 0, x1 = 1, y1 = -1, x2 = 3, y2 = 1
 * 输出：true
 * 解释：圆和矩形有公共点 (1,0)
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：radius = 1, x_center = 0, y_center = 0, x1 = -1, y1 = 0, x2 = 0, y2 = 1
 * 输出：true
 *
 *
 * 示例 3：
 *
 *
 *
 * 输入：radius = 1, x_center = 1, y_center = 1, x1 = -3, y1 = -3, x2 = 3, y2 = 3
 * 输出：true
 *
 *
 * 示例 4：
 *
 * 输入：radius = 1, x_center = 1, y_center = 1, x1 = 1, y1 = -3, x2 = 2, y2 = -1
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= radius <= 2000
 * -10^4 <= x_center, y_center, x1, y1, x2, y2 <= 10^4
 * x1 < x2
 * y1 < y2
 *
 *
 */

/*
 * func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
 * 	// 圆心在矩形内
 * 	if x1 <= xCenter && xCenter <= x2 && y1 <= yCenter && yCenter <= y2 {
 * 		return true
 * 	}
 * 	// 圆心在矩形外, 找最近边
 * 	if xCenter < x1 && x1-xCenter > radius {
 * 		return false
 * 	}
 * 	if xCenter > x2 && xCenter-x2 > radius {
 * 		return false
 * 	}
 * 	if yCenter < y1 && y1-yCenter > radius {
 * 		return false
 * 	}
 * 	if yCenter > y2 && yCenter-y2 > radius {
 * 		return false
 * 	}
 * 
 * 	if y1 <= yCenter && yCenter <= y2 {
 * 		return true
 * 	}
 * 	if x1 <= xCenter && xCenter <= x2 {
 * 		return true
 * 	}
 * 	if xCenter < x1 {
 * 		if y1 > yCenter {
 * 			return radius*radius >= (y1-yCenter)*(y1-yCenter)+(x1-xCenter)*(x1-xCenter)
 * 		} else {
 * 			return radius*radius >= (y2-yCenter)*(y2-yCenter)+(x1-xCenter)*(x1-xCenter)
 * 		}
 * 	} else if xCenter > x2 {
 * 		if y1 > yCenter {
 * 			return radius*radius >= (y1-yCenter)*(y1-yCenter)+(x2-xCenter)*(x2-xCenter)
 * 		} else {
 * 			return radius*radius >= (y2-yCenter)*(y2-yCenter)+(x2-xCenter)*(x2-xCenter)
 * 		}
 * 	}
 * 	return false
 * }
 */

// 参考: https://leetcode-cn.com/problems/circle-and-rectangle-overlapping/solution/1401-cchao-100de-shu-xue-jie-fa-by-ffret-1ynn/
//
// 如何找最近的点?
func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	var dx, dy int
	if x1 > xCenter {
		dx = xCenter - x1
	} else if x2 < xCenter {
		dx = xCenter - x2
	}
	if y1 > yCenter {
		dy = yCenter - y1
	} else if y2 < yCenter {
		dy = yCenter - y2
	}
	return dx*dx+dy*dy <= radius*radius
}
