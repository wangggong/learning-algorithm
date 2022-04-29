/*
 * @lc app=leetcode.cn id=largest-palindrome-product lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [479] 最大回文数乘积
 *
 * https://leetcode-cn.com/problems/largest-palindrome-product/description/
 *
 * algorithms
 * Hard (44.16%)
 * Total Accepted:    8.9K
 * Total Submissions: 15.7K
 * Testcase Example:  '2'
 *
 * 给定一个整数 n ，返回 可表示为两个 n 位整数乘积的 最大回文整数 。因为答案可能非常大，所以返回它对 1337 取余 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入：n = 2
 * 输出：987
 * 解释：99 x 91 = 9009, 9009 % 1337 = 987
 *
 *
 * 示例 2:
 *
 *
 * 输入： n = 1
 * 输出： 9
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= n <= 8
 *
 *
 */
const mod int = 1337

// 从大到小枚举回文数, 然后判定是否是两个 `n` 位数的积.
//
// 关键点在于如何枚举. 首先范围是在 `n` - `2n` 位数字, 其次仅需要枚举左半边, 自动生成右半边即可.
// 另外判定是否为两个 `n` 位数的积, 直接从 `n` 位数枚举到 `\sqrt{k}` 即可.
func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}
	for left := pow10(n) - 1; left >= pow10(n-1); left-- {
		k := genPalindrome(left)
		for j := pow10(n) - 1; j*j >= k; j-- {
			if k%j == 0 {
				return k % mod
			}
		}
	}
	return -1
}

func pow10(k int) int {
	ans := 1
	for k > 0 {
		ans *= 10
		k--
	}
	return ans
}

func genPalindrome(left int) int {
	// assert left > 0
	ans := left
	for ; left > 0; left /= 10 {
		ans = ans*10 + (left % 10)
	}
	return ans
}

