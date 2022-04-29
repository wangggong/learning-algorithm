/*
 * @lc app=leetcode.cn id=add-to-array-form-of-integer lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [989] 数组形式的整数加法
 *
 * https://leetcode-cn.com/problems/add-to-array-form-of-integer/description/
 *
 * algorithms
 * Easy (46.91%)
 * Total Accepted:    59.6K
 * Total Submissions: 127K
 * Testcase Example:  '[1,2,0,0]\n34'
 *
 * 整数的 数组形式  num 是按照从左到右的顺序表示其数字的数组。
 *
 *
 * 例如，对于 num = 1321 ，数组形式是 [1,3,2,1] 。
 *
 *
 * 给定 num ，整数的 数组形式 ，和整数 k ，返回 整数 num + k 的 数组形式 。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：num = [1,2,0,0], k = 34
 * 输出：[1,2,3,4]
 * 解释：1200 + 34 = 1234
 *
 *
 * 示例 2：
 *
 *
 * 输入：num = [2,7,4], k = 181
 * 输出：[4,5,5]
 * 解释：274 + 181 = 455
 *
 *
 * 示例 3：
 *
 *
 * 输入：num = [2,1,5], k = 806
 * 输出：[1,0,2,1]
 * 解释：215 + 806 = 1021
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= num.length <= 10^4
 * 0 <= num[i] <= 9
 * num 不包含任何前导零，除了零本身
 * 1 <= k <= 10^4
 *
 *
 */
func addToArrayForm(num []int, k int) []int {
	reverse(num)
	C := 0
	ind := 0
	for ; k > 0; k, ind = k/10, ind+1 {
		if ind < len(num) {
			sum := num[ind] + (k % 10) + C
			num[ind] = sum % 10
			C = sum / 10
		} else {
			sum := (k % 10) + C
			num = append(num, sum%10)
			C = sum / 10
		}
	}
	for ; C > 0; ind++ {
		if ind < len(num) {
			sum := num[ind] + C
			num[ind] = sum % 10
			C = sum / 10
		} else {
			num = append(num, C%10)
			C /= 10
		}
	}
	reverse(num)
	return num
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
