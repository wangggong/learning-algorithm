/*
 * @lc app=leetcode.cn id=construct-the-longest-new-string lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6895] 构造最长的新字符串
 *
 * https://leetcode.cn/problems/construct-the-longest-new-string/description/
 *
 * algorithms
 * Medium (52.64%)
 * Total Accepted:    1.6K
 * Total Submissions: 3K
 * Testcase Example:  '2\n5\n1'
 *
 * 给你三个整数 x ，y 和 z 。
 * 
 * 这三个整数表示你有 x 个 "AA" 字符串，y 个 "BB" 字符串，和 z 个 "AB"
 * 字符串。你需要选择这些字符串中的部分字符串（可以全部选择也可以一个都不选择），将它们按顺序连接得到一个新的字符串。新字符串不能包含子字符串 "AAA"
 * 或者 "BBB" 。
 * 
 * 请你返回新字符串的最大可能长度。
 * 
 * 子字符串 是一个字符串中一段连续 非空 的字符序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：x = 2, y = 5, z = 1
 * 输出：12
 * 解释： 我们可以按顺序连接 "BB" ，"AA" ，"BB" ，"AA" ，"BB" 和 "AB" ，得到新字符串 "BBAABBAABBAB" 。
 * 字符串长度为 12 ，无法得到一个更长的符合题目要求的字符串。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：x = 3, y = 2, z = 2
 * 输出：14
 * 解释：我们可以按顺序连接 "AB" ，"AB" ，"AA" ，"BB" ，"AA" ，"BB" 和 "AA" ，得到新字符串
 * "ABABAABBAABBAA" 。
 * 字符串长度为 14 ，无法得到一个更长的符合题目要求的字符串。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= x, y, z <= 50
 * 
 * 
 */
public class Solution
{
    public int LongestString(int x, int y, int z)
    {
        var memo = new Dictionary<(int, int, int, int), int>();
        int dfs(int k, int x, int y, int z)
        {
            var key = (k, x, y, z);
            if (!memo.ContainsKey(key))
            {
                memo[key] = k switch
                {
                    1 => y == 0
                        ? 0
                        : (2 + Math.Max(dfs(0, x, y - 1, z), dfs(2, x, y - 1, z))),
                    2 => z == 0
                        ? 0
                        : (2 + Math.Max(dfs(0, x, y, z - 1), dfs(2, x, y, z - 1))),
                    _ => x == 0
                        ? 0
                        : (2 + dfs(1, x - 1, y, z)),
                };
            }
            return memo[key];
        }
        return Enumerable
            .Range(0, 3)
            .Select(i => dfs(i, x, y, z))
            .Max();
    }
}
