/*
 * @lc app=leetcode.cn id=largest-number-after-digit-swaps-by-parity lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6037] 按奇偶性交换后的最大数字
 *
 * https://leetcode-cn.com/problems/largest-number-after-digit-swaps-by-parity/description/
 *
 * algorithms
 * Easy (60.62%)
 * Total Accepted:    6.5K
 * Total Submissions: 10.8K
 * Testcase Example:  '1234'
 *
 * 给你一个正整数 num 。你可以交换 num 中 奇偶性 相同的任意两位数字（即，都是奇数或者偶数）。
 *
 * 返回交换 任意 次之后 num 的 最大 可能值。
 *
 *
 *
 * 示例 1：
 *
 * 输入：num = 1234
 * 输出：3412
 * 解释：交换数字 3 和数字 1 ，结果得到 3214 。
 * 交换数字 2 和数字 4 ，结果得到 3412 。
 * 注意，可能存在其他交换序列，但是可以证明 3412 是最大可能值。
 * 注意，不能交换数字 4 和数字 1 ，因为它们奇偶性不同。
 *
 *
 * 示例 2：
 *
 * 输入：num = 65875
 * 输出：87655
 * 解释：交换数字 8 和数字 6 ，结果得到 85675 。
 * 交换数字 5 和数字 7 ，结果得到 87655 。
 * 注意，可能存在其他交换序列，但是可以证明 87655 是最大可能值。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= num <= 10^9
 *
 *
 */
func largestInteger(num int) int {
	if num == 0 {
		return 0
	}
	counter := make([]int, 10)
	var arr []int
	for num > 0 {
		arr = append(arr, num%10)
		counter[num%10]++
		num /= 10
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	ans := 0
	for _, a := range arr {
		for i := 9; i >= 0; i-- {
			if (i-a)%2 == 0 && counter[i] > 0 {
				ans = ans*10 + i
				counter[i]--
				break
			}
		}
	}
	return ans
}
