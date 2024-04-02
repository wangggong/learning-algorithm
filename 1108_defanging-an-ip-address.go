/*
 * @lc app=leetcode.cn id=defanging-an-ip-address lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1108] IP 地址无效化
 *
 * https://leetcode.cn/problems/defanging-an-ip-address/description/
 *
 * algorithms
 * Easy (83.82%)
 * Total Accepted:    81.9K
 * Total Submissions: 96.9K
 * Testcase Example:  '"1.1.1.1"'
 *
 * 给你一个有效的 IPv4 地址 address，返回这个 IP 地址的无效化版本。
 *
 * 所谓无效化 IP 地址，其实就是用 "[.]" 代替了每个 "."。
 *
 *
 *
 * 示例 1：
 *
 * 输入：address = "1.1.1.1"
 * 输出："1[.]1[.]1[.]1"
 *
 *
 * 示例 2：
 *
 * 输入：address = "255.100.50.0"
 * 输出："255[.]100[.]50[.]0"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 给出的 address 是一个有效的 IPv4 地址
 *
 *
 */
func defangIPaddr(address string) string {
	n := len(address)
	arr := make([]byte, n+6)
	for i, j := 0, 0; i < n; i++ {
		if address[i] != '.' {
			arr[j] = byte(address[i])
			j++
		} else {
			arr[j] = byte('[')
			j++
			arr[j] = byte(address[i])
			j++
			arr[j] = byte(']')
			j++
		}
	}
	return string(arr)
}
