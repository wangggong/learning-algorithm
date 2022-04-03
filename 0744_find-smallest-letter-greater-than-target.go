/*
 * @lc app=leetcode.cn id=find-smallest-letter-greater-than-target lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [744] 寻找比目标字母大的最小字母
 *
 * https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target/description/
 *
 * algorithms
 * Easy (45.78%)
 * Total Accepted:    72.4K
 * Total Submissions: 147.8K
 * Testcase Example:  '["c","f","j"]\n"a"'
 *
 * 给你一个排序后的字符列表 letters ，列表中只包含小写英文字母。另给出一个目标字母
 * target，请你寻找在这一有序列表里比目标字母大的最小字母。
 *
 * 在比较时，字母是依序循环出现的。举个例子：
 *
 *
 * 如果目标字母 target = 'z' 并且字符列表为 letters = ['a', 'b']，则答案返回 'a'
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: letters = ["c", "f", "j"]，target = "a"
 * 输出: "c"
 *
 *
 * 示例 2:
 *
 *
 * 输入: letters = ["c","f","j"], target = "c"
 * 输出: "f"
 *
 *
 * 示例 3:
 *
 *
 * 输入: letters = ["c","f","j"], target = "d"
 * 输出: "f"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= letters.length <= 10^4
 * letters[i] 是一个小写字母
 * letters 按非递减顺序排序
 * letters 最少包含两个不同的字母
 * target 是一个小写字母
 *
 *
 */

func nextGreatestLetter(letters []byte, target byte) byte {
	arr := make([]bool, 26)
	for _, l := range letters {
		arr[int(l-'a')] = true
	}
	for i := byte(1); i <= 26; i++ {
		b := byte((target-'a'+i)%26) + 'a'
		if arr[int(b-'a')] {
			return b
		}
	}
	return byte(0)
}
