/*
 * @lc app=leetcode.cn id=decremental-string-concatenation lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6898] 字符串连接删减字母
 *
 * https://leetcode.cn/problems/decremental-string-concatenation/description/
 *
 * algorithms
 * Medium (26.95%)
 * Total Accepted:    835
 * Total Submissions: 3.1K
 * Testcase Example:  '["aa","ab","bc"]'
 *
 * 给你一个下标从 0 开始的数组 words ，它包含 n 个字符串。
 * 
 * 定义 连接 操作 join(x, y) 表示将字符串 x 和 y 连在一起，得到 xy 。如果 x 的最后一个字符与 y
 * 的第一个字符相等，连接后两个字符中的一个会被 删除 。
 * 
 * 比方说 join("ab", "ba") = "aba" ， join("ab", "cde") = "abcde" 。
 * 
 * 你需要执行 n - 1 次 连接 操作。令 str0 = words[0] ，从 i = 1 直到 i = n - 1 ，对于第 i
 * 个操作，你可以执行以下操作之一：
 * 
 * 
 * 令 stri = join(stri - 1, words[i])
 * 令 stri = join(words[i], stri - 1)
 * 
 * 
 * 你的任务是使 strn - 1 的长度 最小 。
 * 
 * 请你返回一个整数，表示 strn - 1 的最小长度。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：words = ["aa","ab","bc"]
 * 输出：4
 * 解释：这个例子中，我们按以下顺序执行连接操作，得到 str2 的最小长度：
 * str0 = "aa"
 * str1 = join(str0, "ab") = "aab"
 * str2 = join(str1, "bc") = "aabc" 
 * str2 的最小长度为 4 。
 * 
 * 示例 2：
 * 
 * 
 * 输入：words = ["ab","b"]
 * 输出：2
 * 解释：这个例子中，str0 = "ab"，可以得到两个不同的 str1：
 * join(str0, "b") = "ab" 或者 join("b", str0) = "bab" 。
 * 第一个字符串 "ab" 的长度最短，所以答案为 2 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：words = ["aaa","c","aba"]
 * 输出：6
 * 解释：这个例子中，我们按以下顺序执行连接操作，得到 str2 的最小长度：
 * str0 = "aaa"
 * str1 = join(str0, "c") = "aaac"
 * str2 = join("aba", str1) = "abaaac"
 * str2 的最小长度为 6 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= words.length <= 1000
 * 1 <= words[i].length <= 50
 * words[i] 中只包含小写英文字母。
 * 
 * 
 */
public class Solution
{
    public int MinimizeConcatenatedLength(string[] words)
    {
        var chars = words
            .Select(w => w.ToCharArray())
            .ToArray();
        var n = chars.Length;
        var memo = new Dictionary<(int, char, char), int>();
        int dfs(int k, char start, char end)
        {
            var key = (k, start, end);
            if (!memo.ContainsKey(key))
            {
                memo[key] = k == n
                    ? 0
                    : Math.Max(
                        dfs(k + 1, start, chars[k].Last()) + (end == chars[k].First() ? 1 : 0),
                        dfs(k + 1, chars[k].First(), end) + (start == chars[k].Last() ? 1 : 0));
            }
            return memo[key];
        }
        return chars
            .Select(cs => cs.Length)
            .Sum() - dfs(1, chars[0].First(), chars[0].Last());
    }
}
