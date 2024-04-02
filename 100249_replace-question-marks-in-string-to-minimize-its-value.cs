/*
 * @lc app=leetcode.cn id=replace-question-marks-in-string-to-minimize-its-value lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100249] 替换字符串中的问号使分数最小
 *
 * https://leetcode.cn/problems/replace-question-marks-in-string-to-minimize-its-value/description/
 *
 * algorithms
 * Medium (29.62%)
 * Total Accepted:    1.5K
 * Total Submissions: 5.1K
 * Testcase Example:  '"???"'
 *
 * 给你一个字符串 s 。s[i] 要么是小写英文字母，要么是问号 '?' 。
 * 
 * 对于长度为 m 且 只 含有小写英文字母的字符串 t ，我们定义函数 cost(i) 为下标 i 之前（也就是范围 [0, i - 1] 中）出现过与
 * t[i] 相同 字符出现的次数。
 * 
 * 字符串 t 的 分数 为所有下标 i 的 cost(i) 之 和 。
 * 
 * 比方说，字符串 t = "aab" ：
 * 
 * 
 * cost(0) = 0
 * cost(1) = 1
 * cost(2) = 0
 * 所以，字符串 "aab" 的分数为 0 + 1 + 0 = 1 。
 * 
 * 
 * 你的任务是用小写英文字母 替换 s 中 所有 问号，使 s 的 分数最小 。
 * 
 * 请你返回替换所有问号 '?' 之后且分数最小的字符串。如果有多个字符串的 分数最小 ，那么返回字典序最小的一个。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "???" 
 * 
 * 输出： "abc" 
 * 
 * 解释：这个例子中，我们将 s 中的问号 '?' 替换得到 "abc" 。
 * 
 * 对于字符串 "abc" ，cost(0) = 0 ，cost(1) = 0 和 cost(2) = 0 。
 * 
 * "abc" 的分数为 0 。
 * 
 * 其他修改 s 得到分数 0 的字符串为 "cba" ，"abz" 和 "hey" 。
 * 
 * 这些字符串中，我们返回字典序最小的。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入： s = "a?a?"
 * 
 * 输出： "abac"
 * 
 * 解释：这个例子中，我们将 s 中的问号 '?' 替换得到 "abac" 。
 * 
 * 对于字符串 "abac" ，cost(0) = 0 ，cost(1) = 0 ，cost(2) = 1 和 cost(3) = 0 。
 * 
 * "abac" 的分数为 1 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 10^5
 * s[i] 要么是小写英文字母，要么是 '?' 。
 * 
 * 
 */
public class Solution
{
    public string MinimizeStringValue(string s)
    {
        const int Alpha = 26;
        long getPriority(long k, char c) => k * Alpha + (long)(c - 'a');
        var Q = new PriorityQueue<char, long>();
        for (var c = 'a'; c <= 'z'; c++)
        {
            var count = s.Count(ch => ch == c);
            Q.Enqueue(c, getPriority(count, c));
        }
        var sb = new StringBuilder();
        for (var target = s.Count(c => c == '?'); target > 0; target--)
        {
            Q.TryDequeue(out var c, out var count);
            sb.Append(c);
            Q.Enqueue(c, getPriority(count / Alpha + 1, c));
        }
        var candidates = sb.ToString()
            .OrderBy(c => c)
            .ToList();
        var ans = s.ToCharArray();
        for (var (i, j, n) = (0, 0, ans.Length); i < n; i++)
        {
            if (ans[i] == '?')
            {
                ans[i] = candidates[j++];
            }
        }
        return new string(ans);
    }
}
