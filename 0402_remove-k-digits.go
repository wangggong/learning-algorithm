/*
 * @lc app=leetcode.cn id=remove-k-digits lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [402] 移掉 K 位数字
 *
 * https://leetcode-cn.com/problems/remove-k-digits/description/
 *
 * algorithms
 * Medium (32.46%)
 * Total Accepted:    100.7K
 * Total Submissions: 310.1K
 * Testcase Example:  '"1432219"\n3'
 *
 * 给你一个以字符串表示的非负整数 num 和一个整数 k ，移除这个数中的 k 位数字，使得剩下的数字最小。请你以字符串形式返回这个最小的数字。
 *
 *
 * 示例 1 ：
 *
 *
 * 输入：num = "1432219", k = 3
 * 输出："1219"
 * 解释：移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219 。
 *
 *
 * 示例 2 ：
 *
 *
 * 输入：num = "10200", k = 1
 * 输出："200"
 * 解释：移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
 *
 *
 * 示例 3 ：
 *
 *
 * 输入：num = "10", k = 2
 * 输出："0"
 * 解释：从原数字移除所有的数字，剩余为空就是 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * num 仅由若干位数字（0 - 9）组成
 * 除了 0 本身之外，num 不含任何前导零
 *
 *
 */
func removeKdigits(num string, k int) string {
	if k == 0 {
		return removePrefix0(num)
	}
	arr := []byte(num)
	n := len(arr)
	if k >= n {
		return "0"
	}
	// 傻逼! 你都想到贪心策略维护递减序列了还不想单调栈! 钻牛角尖了知不知道!
	var S []byte
	for _, a := range arr {
		for k > 0 && len(S) > 0 && S[len(S)-1] > a {
			S = S[:len(S)-1]
			k--
		}
		S = append(S, a)
	}
	return removePrefix0(string(S[:len(S)-k]))
}

func removePrefix0(num string) string {
	ind := 0
	arr := []byte(num)
	n := len(arr)
	for ; ind < n && num[ind] == '0'; ind++ {

	}
	if ind == n {
		return "0"
	}
	return string(arr[ind:])
}
