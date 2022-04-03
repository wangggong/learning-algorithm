/*
 * @lc app=leetcode.cn id=reverse-words-in-a-string lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [151] 翻转字符串里的单词
 *
 * https://leetcode-cn.com/problems/reverse-words-in-a-string/description/
 *
 * algorithms
 * Medium (49.53%)
 * Total Accepted:    221.2K
 * Total Submissions: 444.6K
 * Testcase Example:  '"the sky is blue"'
 *
 * 给你一个字符串 s ，颠倒字符串中 单词 的顺序。
 *
 * 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
 *
 * 返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
 *
 * 注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "the sky is blue"
 * 输出："blue is sky the"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "  hello world  "
 * 输出："world hello"
 * 解释：颠倒后的字符串中不能存在前导空格和尾随空格。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "a good   example"
 * 输出："example good a"
 * 解释：如果两个单词间有多余的空格，颠倒后的字符串需要将单词间的空格减少到仅有一个。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^4
 * s 包含英文大小写字母、数字和空格 ' '
 * s 中 至少存在一个 单词
 *
 *
 *
 *
 *
 *
 *
 * 进阶：如果字符串在你使用的编程语言中是一种可变数据类型，请尝试使用 O(1) 额外空间复杂度的 原地 解法。
 *
 */

/*
 * func reverseWords(s string) string {
 * 	ws := strings.Split(s, " ")
 * 	var words []string
 * 	for _, w := range ws {
 * 		if len(w) > 0 {
 * 			words = append(words, w)
 * 		}
 * 	}
 * 	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
 * 		words[i], words[j] = words[j], words[i]
 * 	}
 * 	return strings.Join(words, " ")
 * }
 */

// NOTE:
// 1. 快慢指针, 去除头部空格
// 2. 去除单词间冗余空格
// 3. 去除尾部空格, 开始表演 (好像是编程珠玑里面的技巧)
// 4. 反转整个单词
// 5. 每个单词反转
func reverseWords(s string) string {
	ans := []byte(s)
	n := len(ans)
	fast, slow := 0, 0
	for fast < n && ans[fast] == ' ' {
		fast++
	}
	for fast < n {
		if fast > 1 && ans[fast] == ans[fast-1] && ans[fast] == ' ' {
			fast++
			continue
		}
		ans[slow] = ans[fast]
		fast, slow = fast+1, slow+1
	}
	// 尾部空格: 如果尾部存在空格, 上面的操作会把最后一个字符改为空格.
	if ans[slow-1] == ' ' {
		ans = ans[:slow-1]
	} else {
		ans = ans[:slow]
	}
	n = len(ans)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	p, q := 0, 0
	for q < n {
		for q < n && ans[q] != ' ' {
			q++
		}
		for i, j := p, q-1; i < j; i, j = i+1, j-1 {
			ans[i], ans[j] = ans[j], ans[i]
		}
		// 这里 q 别忘了前进一个.
		q++
		p = q
	}
	return string(ans)
}
