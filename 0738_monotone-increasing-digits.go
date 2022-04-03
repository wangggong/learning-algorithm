/*
 * @lc app=leetcode.cn id=monotone-increasing-digits lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [738] 单调递增的数字
 *
 * https://leetcode-cn.com/problems/monotone-increasing-digits/description/
 *
 * algorithms
 * Medium (50.03%)
 * Total Accepted:    51.6K
 * Total Submissions: 103.1K
 * Testcase Example:  '10'
 *
 * 当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。
 *
 * 给定一个整数 n ，返回 小于或等于 n 的最大数字，且数字呈 单调递增 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: n = 10
 * 输出: 9
 *
 *
 * 示例 2:
 *
 *
 * 输入: n = 1234
 * 输出: 1234
 *
 *
 * 示例 3:
 *
 *
 * 输入: n = 332
 * 输出: 299
 *
 *
 *
 *
 * 提示:
 *
 *
 * 0 <= n <= 10^9
 *
 *
 */
func monotoneIncreasingDigits(n int) int {
	arr := itoa(n)
	m := len(arr)
	ind := -1
	// 找到第一个严格递减的位置, 比如 `43`, 这时很容易发现最大的递增值只能是 `39`.
	for i := 0; i+1 < m; i++ {
		if arr[i] > arr[i+1] {
			ind = i
			break
		}
	}
	// 如果发现没有严格递减的位置, `n` 本身直接满足了, 直接返回.
	if ind < 0 {
		return n
	}
	arr[ind] = arr[ind] - byte(1)
	for ind = ind + 1; ind < m; ind++ {
		arr[ind] = '9'
	}
	// 但替换了一波之后可能还不是单调递增的, 所以递归一波.
	return monotoneIncreasingDigits(atoi(arr))
}

func itoa(n int) []byte {
	if n == 0 {
		return []byte("0")
	}
	var ans []byte
	for ; n > 0; n /= 10 {
		ans = append(ans, byte(n%10)+'0')
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}

func atoi(arr []byte) int {
	ans := 0
	for _, a := range arr {
		ans = ans*10 + int(a-'0')
	}
	return ans
}
