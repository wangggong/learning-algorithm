/*
 * @lc app=leetcode.cn id=add-strings lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [415] 字符串相加
 *
 * https://leetcode-cn.com/problems/add-strings/description/
 *
 * algorithms
 * Easy (54.39%)
 * Total Accepted:    175.6K
 * Total Submissions: 322.8K
 * Testcase Example:  '"11"\n"123"'
 *
 * 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。
 *
 * 你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：num1 = "11", num2 = "123"
 * 输出："134"
 *
 *
 * 示例 2：
 *
 *
 * 输入：num1 = "456", num2 = "77"
 * 输出："533"
 *
 *
 * 示例 3：
 *
 *
 * 输入：num1 = "0", num2 = "0"
 * 输出："0"
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= num1.length, num2.length <= 10^4
 * num1 和num2 都只包含数字 0-9
 * num1 和num2 都不包含任何前导零
 *
 *
 */

const MAXN = 1e4 + 5

var sum [MAXN]int

func addStrings(num1 string, num2 string) string {
	if num1 == "0" {
		return num2
	}
	if num2 == "0" {
		return num1
	}
	for i := range sum {
		sum[i] = 0
	}
	arr1, arr2 := []byte(num1), []byte(num2)
	reverse(arr1)
	reverse(arr2)
	n1, n2 := len(arr1), len(arr2)
	for i := 0; i < max(n1, n2); i++ {
		sum[i] = get(arr1, i) + get(arr2, i)
	}
	// assert sum[MAXN-1] < 10;
	for i := 0; i+1 < MAXN; i++ {
		if sum[i] < 10 {
			continue
		}
		sum[i], sum[i+1] = sum[i]%10, sum[i+1]+sum[i]/10
	}
	index := int(MAXN - 1)
	for index >= 0 && sum[index] == 0 {
		index--
	}
	var ans []byte
	for _, s := range sum[:index+1] {
		ans = append(ans, byte(s+'0'))
	}
	reverse(ans)
	return string(ans)
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func get(arr []byte, index int) int {
	if index >= len(arr) {
		return 0
	}
	return int(arr[index] - '0')
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
