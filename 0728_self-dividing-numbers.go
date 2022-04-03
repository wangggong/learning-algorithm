/*
 * @lc app=leetcode.cn id=self-dividing-numbers lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [728] 自除数
 *
 * https://leetcode-cn.com/problems/self-dividing-numbers/description/
 *
 * algorithms
 * Easy (74.80%)
 * Total Accepted:    51.3K
 * Total Submissions: 66K
 * Testcase Example:  '1\n22'
 *
 * 自除数 是指可以被它包含的每一位数整除的数。
 *
 *
 * 例如，128 是一个 自除数 ，因为 128 % 1 == 0，128 % 2 == 0，128 % 8 == 0。
 *
 *
 * 自除数 不允许包含 0 。
 *
 * 给定两个整数 left 和 right ，返回一个列表，列表的元素是范围 [left, right] 内所有的 自除数 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：left = 1, right = 22
 * 输出：[1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]
 *
 *
 * 示例 2:
 *
 *
 * 输入：left = 47, right = 85
 * 输出：[48,55,66,77]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= left <= right <= 10^4
 *
 *
 */
func selfDividingNumbers(left int, right int) []int {
	var ans []int
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			ans = append(ans, i)
		}
	}
	return ans
}

func isSelfDividing(k int) bool {
	if k != 0 {
		for c := k; c != 0; c /= 10 {
			if c%10 == 0 {
				return false
			}
			if k%(c%10) != 0 {
				return false
			}
		}
	}
	return true
}
