/*
 * Medium
 *
 * 给定一个字符串列表 strs，你可以将这些字符串连接成一个循环字符串，对于每个字符串，你可以选择是否翻转它。在所有可能的循环字符串中，你需要分割循环字符串（这将使循环字符串变成一个常规的字符串），然后找到字典序最大的字符串。
 *
 * 具体来说，要找到字典序最大的字符串，你需要经历两个阶段：
 *
 * 将所有字符串连接成一个循环字符串，你可以选择是否翻转某些字符串，并按照给定的顺序连接它们。
 * 在循环字符串的某个位置分割它，这将使循环字符串从分割点变成一个常规的字符串。
 * 你的工作是在所有可能的常规字符串中找到字典序最大的一个。
 *
 *
 *
 * 示例 1:
 *
 * 输入: strs = ["abc","xyz"]
 * 输出: "zyxcba"
 * 解释: 你可以得到循环字符串 "-abcxyz-", "-abczyx-", "-cbaxyz-", "-cbazyx-"，其中 '-' 代表循环状态。
 * 答案字符串来自第四个循环字符串， 你可以从中间字符 'a' 分割开然后得到 "zyxcba"。
 * 示例 2:
 *
 * 输入: strs = ["abc"]
 * 输出: "cba"
 *
 *
 * 提示:
 *
 * 1 <= strs.length <= 1000
 * 1 <= strs[i].length <= 1000
 * 1 <= sum(strs[i].length) <= 1000
 * strs[i] 只包含小写英文字母
 * 通过次数1,347
 * 提交次数3,747
 */

// 题目意思很难理解, 是说字符串拼接前可以选择是否要翻转这个字符串;
// 拼接后形成一个环, 可以选择把这个环在任何位置切开, 求字典序最大.
// 解法主要注意:
// - 如果不是 "切开位置" 的词的话, 可以先判断是否要翻转;
// - 如果这个词要被切开了, 就需要枚举翻转前与翻转后的所有切点, 取最大值;
// - 因此只需要枚举所有 "切开位置" 的词语, 按上述说法走一遍, 取全局最大就好.
func splitLoopedString(strs []string) string {
	n := len(strs)
	for i := range strs {
		if rev := getReverse(strs[i]); rev > strs[i] {
			strs[i] = rev
		}
	}
	ans := ""
	for _, s := range strs {
		ans = ans + s
	}
	m := len(ans)
	for i := 0; i < n; i++ {
		tot := ""
		for j := 1; j < n; j++ {
			tot = tot + strs[(i+j)%n]
		}
		cur := strs[i] + tot
		cur = cur + cur
		for j := 0; j < len(strs[i]); j++ {
			if cur[j:j+m] > ans {
				ans = cur[j : j+m]
			}
		}
		cur = getReverse(strs[i]) + tot
		cur = cur + cur
		for j := 0; j < len(strs[i]); j++ {
			if cur[j:j+m] > ans {
				ans = cur[j : j+m]
			}
		}
	}
	return ans
}

func getReverse(s string) string {
	arr := []byte(s)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return string(arr)
}
