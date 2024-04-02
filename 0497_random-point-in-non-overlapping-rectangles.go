/*
 * @lc app=leetcode.cn id=random-point-in-non-overlapping-rectangles lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [497] 非重叠矩形中的随机点
 *
 * https://leetcode.cn/problems/random-point-in-non-overlapping-rectangles/description/
 *
 * algorithms
 * Medium (40.90%)
 * Total Accepted:    8.6K
 * Total Submissions: 20.5K
 * Testcase Example:  '["Solution","pick","pick","pick","pick","pick"]\n' +
  '[[[[-2,-2,1,1],[2,2,4,6]]],[],[],[],[],[]]'
 *
 * 给定一个由非重叠的轴对齐矩形的数组 rects ，其中 rects[i] = [ai, bi, xi, yi] 表示 (ai, bi) 是第 i
 * 个矩形的左下角点，(xi, yi) 是第 i
 * 个矩形的右上角点。设计一个算法来随机挑选一个被某一矩形覆盖的整数点。矩形周长上的点也算做是被矩形覆盖。所有满足要求的点必须等概率被返回。
 *
 * 在给定的矩形覆盖的空间内的任何整数点都有可能被返回。
 *
 * 请注意 ，整数点是具有整数坐标的点。
 *
 * 实现 Solution 类:
 *
 *
 * Solution(int[][] rects) 用给定的矩形数组 rects 初始化对象。
 * int[] pick() 返回一个随机的整数点 [u, v] 在给定的矩形所覆盖的空间内。
 *
 *
 *
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
 * ["Solution", "pick", "pick", "pick", "pick", "pick"]
 * [[[[-2, -2, 1, 1], [2, 2, 4, 6]]], [], [], [], [], []]
 * 输出:
 * [null, [1, -2], [1, -1], [-1, -2], [-2, -2], [0, 0]]
 *
 * 解释：
 * Solution solution = new Solution([[-2, -2, 1, 1], [2, 2, 4, 6]]);
 * solution.pick(); // 返回 [1, -2]
 * solution.pick(); // 返回 [1, -1]
 * solution.pick(); // 返回 [-1, -2]
 * solution.pick(); // 返回 [-2, -2]
 * solution.pick(); // 返回 [0, 0]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= rects.length <= 100
 * rects[i].length == 4
 * -10^9 <= ai < xi <= 10^9
 * -10^9 <= bi < yi <= 10^9
 * xi - ai <= 2000
 * yi - bi <= 2000
 * 所有的矩形不重叠。
 * pick 最多被调用 10^4 次。
 *
 *
*/
import "math/rand"

type Solution struct {
	Rects [][]int
	Tot   int
}

// 关键点 #1: 面积的求法, 别忘了 `+1`.
func Area(rect []int) int { return (rect[2] - rect[0] + 1) * (rect[3] - rect[1] + 1) }

func Constructor(rects [][]int) Solution {
	tot := 0
	for _, r := range rects {
		tot += Area(r)
	}
	return Solution{
		Rects: rects,
		Tot:   tot,
	}
}

func (s *Solution) Pick() []int {
	k := rand.Intn(s.Tot)
	for i, cur, n := 0, 0, len(s.Rects); i < n; i++ {
		r := s.Rects[i]
		S := Area(r)
        // 关键点 #2: 取等条件.
		if cur+S <= k {
			cur += S
			continue
		}
		k -= cur
		return []int{r[0] + k/(r[3]-r[1]+1), r[1] + k%(r[3]-r[1]+1)}
	}
	// unreachable
	return nil
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */
