/*
 * @lc app=leetcode.cn id=sender-with-largest-word-count lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6084] 最多单词数的发件人
 *
 * https://leetcode-cn.com/problems/sender-with-largest-word-count/description/
 *
 * algorithms
 * Medium (56.28%)
 * Total Accepted:    3.4K
 * Total Submissions: 6K
 * Testcase Example:  '["Hello userTwooo","Hi userThree","Wonderful day Alice","Nice day userThree"]\n' +
  '["Alice","userTwo","userThree","Alice"]'
 *
 * 给你一个聊天记录，共包含 n 条信息。给你两个字符串数组 messages 和 senders ，其中 messages[i] 是 senders[i]
 * 发出的一条 信息 。
 *
 * 一条 信息 是若干用单个空格连接的 单词 ，信息开头和结尾不会有多余空格。发件人的 单词计数 是这个发件人总共发出的 单词数
 * 。注意，一个发件人可能会发出多于一条信息。
 *
 * 请你返回发出单词数 最多 的发件人名字。如果有多个发件人发出最多单词数，请你返回 字典序 最大的名字。
 *
 * 注意：
 *
 *
 * 字典序里，大写字母小于小写字母。
 * "Alice" 和 "alice" 是不同的名字。
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：messages = ["Hello userTwooo","Hi userThree","Wonderful day Alice","Nice
 * day userThree"], senders = ["Alice","userTwo","userThree","Alice"]
 * 输出："Alice"
 * 解释：Alice 总共发出了 2 + 3 = 5 个单词。
 * userTwo 发出了 2 个单词。
 * userThree 发出了 3 个单词。
 * 由于 Alice 发出单词数最多，所以我们返回 "Alice" 。
 *
 *
 * 示例 2：
 *
 * 输入：messages = ["How is leetcode for everyone","Leetcode is useful for
 * practice"], senders = ["Bob","Charlie"]
 * 输出："Charlie"
 * 解释：Bob 总共发出了 5 个单词。
 * Charlie 总共发出了 5 个单词。
 * 由于最多单词数打平，返回字典序最大的名字，也就是 Charlie 。
 *
 *
 *
 * 提示：
 *
 *
 * n == messages.length == senders.length
 * 1 <= n <= 10^4
 * 1 <= messages[i].length <= 100
 * 1 <= senders[i].length <= 10
 * messages[i] 包含大写字母、小写字母和 ' ' 。
 * messages[i] 中所有单词都由 单个空格 隔开。
 * messages[i] 不包含前导和后缀空格。
 * senders[i] 只包含大写英文字母和小写英文字母。
 *
 *
*/

import (
	"sort"
	"strings"
)

type cntInfo struct {
	Cnt  int
	Name string
}

func largestWordCount(M []string, S []string) string {
	cnt := make(map[string]int)
	n := len(M)
	for i := 0; i < n; i++ {
		cnt[S[i]] += len(strings.Split(M[i], " "))
	}
	var ans []cntInfo
	for n, c := range cnt {
		ans = append(ans, cntInfo{Cnt: c, Name: n})
	}
	sort.Slice(ans, func(p, q int) bool {
		return ans[p].Cnt > ans[q].Cnt || ans[p].Cnt == ans[q].Cnt && ans[p].Name > ans[q].Name
	})
	return ans[0].Name
}
