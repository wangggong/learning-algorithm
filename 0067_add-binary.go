/*
 * @lc app=leetcode.cn id=add-binary lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [67] 二进制求和
 *
 * https://leetcode-cn.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (54.02%)
 * Total Accepted:    223.4K
 * Total Submissions: 413.7K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给你两个二进制字符串，返回它们的和（用二进制表示）。
 *
 * 输入为 非空 字符串且只包含数字 1 和 0。
 *
 *
 *
 * 示例 1:
 *
 * 输入: a = "11", b = "1"
 * 输出: "100"
 *
 * 示例 2:
 *
 * 输入: a = "1010", b = "1011"
 * 输出: "10101"
 *
 *
 *
 * 提示：
 *
 *
 * 每个字符串仅由字符 '0' 或 '1' 组成。
 * 1 <= a.length, b.length <= 10^4
 * 字符串如果不是 "0" ，就都不含前导零。
 *
 *
 */
func addBinary(a string, b string) string {
	ba, bb := []byte(a), []byte(b)
	reverse(ba)
	reverse(bb)
	na, nb := len(ba), len(bb)
	c := 0
	var ans []byte
	for i := 0; i < max(na, nb); i++ {
		p, q := get(ba, i), get(bb, i)
		v := p + q + c
		c = v / 2
		ans = append(ans, byte(v%2)+'0')
	}
	if c > 0 {
		ans = append(ans, byte(c%2)+'0')
	}
	reverse(ans)
	return string(ans)
}

// 1010
// 1011

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func get(arr []byte, index int) int {
	if index >= len(arr) {
		return 0
	}
	return int(arr[index] - '0')
}
