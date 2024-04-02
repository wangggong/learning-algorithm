/*
 * @lc app=leetcode.cn id=preimage-size-of-factorial-zeroes-function lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [793] 阶乘函数后 K 个零
 *
 * https://leetcode.cn/problems/preimage-size-of-factorial-zeroes-function/description/
 *
 * algorithms
 * Hard (46.84%)
 * Total Accepted:    19.2K
 * Total Submissions: 41K
 * Testcase Example:  '0'
 *
 *  f(x) 是 x! 末尾是 0 的数量。回想一下 x! = 1 * 2 * 3 * ... * x，且 0! = 1 。
 *
 *
 * 例如， f(3) = 0 ，因为 3! = 6 的末尾没有 0 ；而 f(11) = 2 ，因为 11!= 39916800 末端有 2 个 0 。
 *
 *
 * 给定 k，找出返回能满足 f(x) = k 的非负整数 x 的数量。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：k = 0
 * 输出：5
 * 解释：0!, 1!, 2!, 3!, 和 4! 均符合 k = 0 的条件。
 *
 *
 * 示例 2：
 *
 *
 * 输入：k = 5
 * 输出：0
 * 解释：没有匹配到这样的 x!，符合 k = 5 的条件。
 *
 * 示例 3:
 *
 *
 * 输入: k = 3
 * 输出: 5
 *
 *
 *
 *
 * 提示:
 *
 *
 * 0 <= k <= 10^9
 *
 *
 */

const maxn int = 1e10

func preimageSizeFZF(k int) int {
	return largerFZF(k+1) - largerFZF(k)
}

func largerFZF(k int) int {
	p, q := 0, maxn
	for p < q {
		mid := (p + q) >> 1
		if FZF(mid) >= k {
			q = mid
		} else {
			p = mid + 1
		}
	}
	return p
}

func FZF(k int) (ans int) {
	for ; k > 0; k /= 5 {
		ans += k / 5
	}
	return
}
