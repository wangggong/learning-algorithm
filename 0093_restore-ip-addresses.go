import "errors"

/*
 * @lc app=leetcode.cn id=restore-ip-addresses lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [93] 复原 IP 地址
 *
 * https://leetcode-cn.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (55.25%)
 * Total Accepted:    191.1K
 * Total Submissions: 344.9K
 * Testcase Example:  '"25525511135"'
 *
 * 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
 *
 *
 * 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312"
 * 和 "192.168@1.1" 是 无效 IP 地址。
 *
 *
 * 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能
 * 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "25525511135"
 * 输出：["255.255.11.135","255.255.111.35"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "0000"
 * 输出：["0.0.0.0"]
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "101023"
 * 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 20
 * s 仅由数字组成
 *
 *
 */

var ans []string
var cur []int
var errInvalidStr = errors.New("invalid string")

func restoreIpAddresses(s string) []string {
	bytes := []byte(s)
	ans, cur = nil, nil
	backtrack(bytes, 4)
	return ans
}

func backtrack(arr []byte, k int) {
	if k == 0 {
		if len(arr) > 0 {
			return
		}
		var tmp []byte
		for i, c := range cur {
			tmp = append(tmp, itoa(c)...)
			if i+1 < len(cur) {
				tmp = append(tmp, '.')
			}
		}
		ans = append(ans, string(tmp))
		return
	}
	n := len(arr)
	if n == 0 {
		return
	}
	for i := 1; i <= n; i++ {
		ans, err := atoi(arr[:i])
		if err != nil || ans < 0 || ans >= (1<<8) {
			continue
		}
		cur = append(cur, ans)
		backtrack(arr[i:], k-1)
		cur = cur[:len(cur)-1]
	}
}

func atoi(arr []byte) (int, error) {
	n := len(arr)
	if n == 0 {
		return 0, errInvalidStr
	}
	if n > 1 && arr[0] == '0' {
		return 0, errInvalidStr
	}
	ans := 0
	for _, a := range arr {
		if a < '0' || a > '9' {
			return 0, errInvalidStr
		}
		ans = ans*10 + int(a-'0')
	}
	return ans, nil
}

func itoa(x int) []byte {
	if x == 0 {
		return []byte("0")
	}
	var ans []byte
	for x > 0 {
		ans = append(ans, byte(x%10)+'0')
		x /= 10
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}
