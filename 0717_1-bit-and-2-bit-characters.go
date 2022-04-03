/*
 * @lc app=leetcode.cn id=1-bit-and-2-bit-characters lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [717] 1比特与2比特字符
 *
 * https://leetcode-cn.com/problems/1-bit-and-2-bit-characters/description/
 *
 * algorithms
 * Easy (51.33%)
 * Total Accepted:    33.8K
 * Total Submissions: 65.9K
 * Testcase Example:  '[1,0,0]'
 *
 * 有两种特殊字符：
 *
 *
 * 第一种字符可以用一个比特 0 来表示
 * 第二种字符可以用两个比特(10 或 11)来表示、
 *
 *
 * 给定一个以 0 结尾的二进制数组 bits ，如果最后一个字符必须是一位字符，则返回 true 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: bits = [1, 0, 0]
 * 输出: true
 * 解释: 唯一的编码方式是一个两比特字符和一个一比特字符。
 * 所以最后一个字符是一比特字符。
 *
 *
 * 示例 2:
 *
 *
 * 输入: bits = [1, 1, 1, 0]
 * 输出: false
 * 解释: 唯一的编码方式是两比特字符和两比特字符。
 * 所以最后一个字符不是一比特字符。
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= bits.length <= 1000
 * bits[i] == 0 or 1
 *
 *
 */
func isOneBitCharacter(bits []int) bool {
	/*
	n := len(bits)
	// assert n > 0 && n <= 1000;
	if n == 1 {
		return true
	}
	if bits[n-1] == 1 {
		return false
	}
	if n > 1 && bits[n-2] == 0 {
		return true
	}
	if valid(bits[:n-2]) {
		return false
	}
	return true
	*/
	n := len(bits)
	last := false
	for i := 0; i < n; i++ {
		if bits[i] == 0 {
			last = true
			continue
		}
		if bits[i] == 1 {
			i++
			last = false
		}
	}
	return last
}

func valid(bits []int) bool {
	n := len(bits)
	if n == 0 {
		return true
	}
	if n == 1 {
		return bits[0] == 0
	}
	if bits[n-1] == 0 && valid(bits[:n-1]) {
		return true
	}
	a, b := bits[n-2], bits[n-1]
	if a == 0 && b == 1 {
		return false
	}
	return valid(bits[:n-2])
}
