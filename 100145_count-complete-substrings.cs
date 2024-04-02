/*
 * @lc app=leetcode.cn id=count-complete-substrings lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100145] 统计完全子字符串
 *
 * https://leetcode.cn/problems/count-complete-substrings/description/
 *
 * algorithms
 * Medium (24.63%)
 * Total Accepted:    1.5K
 * Total Submissions: 6.1K
 * Testcase Example:  '"igigee"\n2'
 *
 * 给你一个字符串 word 和一个整数 k 。
 * 
 * 如果 word 的一个子字符串 s 满足以下条件，我们称它是 完全字符串：
 * 
 * 
 * s 中每个字符 恰好 出现 k 次。
 * 相邻字符在字母表中的顺序 至多 相差 2 。也就是说，s 中两个相邻字符 c1 和 c2 ，它们在字母表中的位置相差 至多 为 2 。
 * 
 * 
 * 请你返回 word 中 完全 子字符串的数目。
 * 
 * 子字符串 指的是一个字符串中一段连续 非空 的字符序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：word = "igigee", k = 2
 * 输出：3
 * 解释：完全子字符串需要满足每个字符恰好出现 2 次，且相邻字符相差至多为 2 ：igigee, igigee, igigee 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：word = "aaabbbccc", k = 3
 * 输出：6
 * 解释：完全子字符串需要满足每个字符恰好出现 3 次，且相邻字符相差至多为 2 ：aaabbbccc, aaabbbccc, aaabbbccc,
 * aaabbbccc, aaabbbccc, aaabbbccc 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= word.length <= 10^5
 * word 只包含小写英文字母。
 * 1 <= k <= word.length
 * 
 * 
 */
public class Solution
{
    private int Count(string s, int k)
    {
        const int Alpha = 26;
        int n = s.Length;
        var counts = new int[n + 1][];
        counts[0] = new int[Alpha];
        var ans = 0;
        for (var i = 1; i <= n; i++)
        {
            counts[i] = new int[Alpha];
            Array.Copy(counts[i - 1], counts[i], Alpha);
            counts[i][(int)(s[i - 1] - 'a')]++;
            ans += Enumerable.Range(1, Alpha)
                .Count(j =>
                {
                    if (j * k > i) { return false; }
                    for (var a = 0; a < Alpha; a++)
                    {
                        if (counts[i][a] == counts[i - j * k][a]) { continue; }
                        if (counts[i][a] - counts[i - j * k][a] != k) { return false; }
                    }
                    return true;
                });
        }
        return ans;
    }

    public int CountCompleteSubstrings(string word, int k)
    {
        var ans = 0;
        for (var (p, q, n) = (0, 0, word.Length); p < n; p = q)
        {
            for (q = p + 1; q < n && Math.Abs(word[q] - word[q - 1]) <= 2; q++) { }
            ans += Count(word[p..q], k);
        }
        return ans;
    }
}
