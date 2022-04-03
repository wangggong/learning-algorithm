/*
 * Easy
 *
 * 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "We are happy."
 * 输出："We%20are%20happy."
 *
 *
 * 限制：
 *
 * 0 <= s 的长度 <= 10000
 *
 * 通过次数394,781
 * 提交次数519,540
 */

func replaceSpace(s string) string {
	bytes := []byte(s)
	var ans []byte
	spc := []byte("%20")
	for _, b := range bytes {
		if b == ' ' {
			ans = append(ans, spc...)
		} else {
			ans = append(ans, b)
		}
	}
	return string(ans)
}
