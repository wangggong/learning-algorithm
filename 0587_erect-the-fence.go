/*
 * @lc app=leetcode.cn id=erect-the-fence lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [587] 安装栅栏
 *
 * https://leetcode-cn.com/problems/erect-the-fence/description/
 *
 * algorithms
 * Hard (39.88%)
 * Total Accepted:    11.8K
 * Total Submissions: 19.6K
 * Testcase Example:  '[[1,1],[2,2],[2,0],[2,4],[3,3],[4,2]]'
 *
 * 在一个二维的花园中，有一些用 (x, y)
 * 坐标表示的树。由于安装费用十分昂贵，你的任务是先用最短的绳子围起所有的树。只有当所有的树都被绳子包围时，花园才能围好栅栏。你需要找到正好位于栅栏边界上的树的坐标。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [[1,1],[2,2],[2,0],[2,4],[3,3],[4,2]]
 * 输出: [[1,1],[2,0],[4,2],[3,3],[2,4]]
 * 解释:
 *
 *
 *
 * 示例 2:
 *
 * 输入: [[1,2],[2,2],[4,2]]
 * 输出: [[1,2],[2,2],[4,2]]
 * 解释:
 *
 * 即使树都在一条直线上，你也需要先用绳子包围它们。
 *
 *
 *
 *
 * 注意:
 *
 *
 * 所有的树应当被围在一起。你不能剪断绳子来包围树或者把树分成一组以上。
 * 输入的整数在 0 到 100 之间。
 * 花园至少有一棵树。
 * 所有树的坐标都是不同的。
 * 输入的点没有顺序。输出顺序也没有要求。
 *
 */

// 凸包的 Graham 算法
//
// 参考:
// - https://blog.csdn.net/viafcccy/article/details/87483567
// - https://leetcode-cn.com/problems/erect-the-fence/solution/an-zhuang-zha-lan-by-leetcode-solution-75s3/
import "math"

func outerTrees(trees [][]int) [][]int {
	n := len(trees)
	if n <= 3 {
		return trees
	}
	// 找左下角的点.
	minYTree := trees[0]
	for _, t := range trees {
		if (t[1] < minYTree[1]) || (t[1] == minYTree[1] && t[0] < minYTree[0]) {
			minYTree = t
		}
	}
	// 按照极角排序, 极角相同的按照远近排序.
	sort.Slice(trees, func(p, q int) bool {
		if c := cross(minYTree, trees[p], trees[q]); c > 0 {
			return true
		} else if c == 0 && dist(minYTree, trees[p]) < dist(minYTree, trees[q]) {
			return true
		}
		return false
	})

	// NOTE: 这里需要排一个序, 考虑的是在凸壳的最后一条边上如果有多个点, 那么需要把顺序反一反, 从最远的点考虑起, 避免丢点.
	r := n - 1
	for ; r >= 0 && cross(minYTree, trees[n-1], trees[r]) == 0; r-- {
	}
	for i, j := r+1, n-1; i < j; i, j = i+1, j-1 {
		trees[i], trees[j] = trees[j], trees[i]
	}

	// 往栈里面塞: 如果栈顶点 `p, q` 到当前点 `r` 是逆时针 / 共线的, 就入栈; 否则就把栈顶点 `q` 出栈.
	var S [][]int
	S = append(S, trees[0], trees[1])
	for i := 2; i < n; i++ {
		for len(S) > 2 && clockwise(S[len(S)-2], S[len(S)-1], trees[i]) {
			S = S[:len(S)-1]
		}
		S = append(S, trees[i])
	}
	return S
}

func getCos(arr, base []int) float64 {
	// assert len(arr) == 2
	// assert len(base) == 2
	if base[0] == arr[0] && base[1] == arr[1] {
		return math.MinInt32
	}
	return float64(arr[0]-base[0]) / math.Sqrt(float64(dist(arr, base)))
}

func dist(p0, p1 []int) int           { return (p0[0]-p1[0])*(p0[0]-p1[0]) + (p0[1]-p1[1])*(p0[1]-p1[1]) }
func cross(p0, p1, p2 []int) int      { return (p1[0]-p0[0])*(p2[1]-p0[1]) - (p1[1]-p0[1])*(p2[0]-p0[0]) }
func clockwise(p0, p1, p2 []int) bool { return cross(p0, p1, p2) < 0 }
