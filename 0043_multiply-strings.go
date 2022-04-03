/*
 * @lc app=leetcode.cn id=multiply-strings lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [43] 字符串相乘
 *
 * https://leetcode-cn.com/problems/multiply-strings/description/
 *
 * algorithms
 * Medium (44.99%)
 * Total Accepted:    198.5K
 * Total Submissions: 441.2K
 * Testcase Example:  '"2"\n"3"'
 *
 * 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
 *
 * 注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: num1 = "2", num2 = "3"
 * 输出: "6"
 *
 * 示例 2:
 *
 *
 * 输入: num1 = "123", num2 = "456"
 * 输出: "56088"
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= num1.length, num2.length <= 200
 * num1 和 num2 只能由数字组成。
 * num1 和 num2 都不包含任何前导零，除了数字0本身。
 *
 *
 */

const MAXN = 200*2 + 5

var mul [MAXN]int

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	arr1, arr2 := []byte(num1), []byte(num2)
	reverse(arr1)
	reverse(arr2)
	n1, n2 := len(arr1), len(arr2)
	for i := range mul {
		mul[i] = 0
	}
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			v1, v2 := int(arr1[i]-'0'), int(arr2[j]-'0')
			p := v1 * v2
			mul[i+j] += p % 10
			mul[i+j+1] += p / 10
		}
	}
	// assert mul[MAXN-1] < 10;
	for i := 0; i+1 < MAXN; i++ {
		if mul[i] < 10 {
			continue
		}
		mul[i+1] += mul[i] / 10
		mul[i] %= 10
	}
	var ans []byte
	index := MAXN - 1
	for index >= 0 && mul[index] == 0 {
		index--
	}
	for _, m := range mul[:index+1] {
		ans = append(ans, byte(m)+'0')
	}
	reverse(ans)
	return string(ans)
}

func reverse(arr []byte) {
	for p, q := 0, len(arr)-1; p < q; p, q = p+1, q-1 {
		arr[p], arr[q] = arr[q], arr[p]
	}
}
