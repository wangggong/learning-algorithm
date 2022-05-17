/*
 * Medium
 *
 * 字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
 *
 *
 *
 * 示例 1:
 *
 * 输入:
 * first = "pale"
 * second = "ple"
 * 输出: True
 *
 *
 * 示例 2:
 *
 * 输入:
 * first = "pales"
 * second = "pal"
 * 输出: False
 * 通过次数51,836提交次数153,633
 *
 */

func oneEditAway(first string, second string) bool {
	n, m := len(first), len(second)
	if n-m > 1 || m-n > 1 {
		return false
	}
	modified := false
	for i, j := 0, 0; i < n && j < m; i, j = i+1, j+1 {
		if first[i] != second[j] {
			if modified {
				return false
			}
			modified = true
			if n > m {
				j--
			} else if n < m {
				i--
			} else {
			}
		}
	}
	return true
}