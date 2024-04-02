/*
 * @lc app=leetcode.cn id=maximum-repeating-substring lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1668] 最大重复子字符串
 *
 * https://leetcode.cn/problems/maximum-repeating-substring/description/
 *
 * algorithms
 * Easy (44.40%)
 * Total Accepted:    32.9K
 * Total Submissions: 70.2K
 * Testcase Example:  '"ababc"\n"ab"'
 *
 * 给你一个字符串 sequence ，如果字符串 word 连续重复 k 次形成的字符串是 sequence 的一个子字符串，那么单词 word 的
 * 重复值为 k 。单词 word 的 最大重复值 是单词 word 在 sequence 中最大的重复值。如果 word 不是 sequence
 * 的子串，那么重复值 k 为 0 。
 * 
 * 给你一个字符串 sequence 和 word ，请你返回 最大重复值 k 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：sequence = "ababc", word = "ab"
 * 输出：2
 * 解释："abab" 是 "ababc" 的子字符串。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：sequence = "ababc", word = "ba"
 * 输出：1
 * 解释："ba" 是 "ababc" 的子字符串，但 "baba" 不是 "ababc" 的子字符串。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：sequence = "ababc", word = "ac"
 * 输出：0
 * 解释："ac" 不是 "ababc" 的子字符串。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= sequence.length <= 100
 * 1 <= word.length <= 100
 * sequence 和 word 都只包含小写英文字母。
 * 
 * 
 */
public class Solution
{
    private const long Prime = 13331;
    private long[] Hash;
    private long[] Mul;

    private long GetHash(int p, int q) => Hash[q] - Hash[p] * Mul[q - p];

    public int MaxRepeating(string sequence, string word)
    {
        int n = sequence.Count();
        int m = word.Count();
        Hash = new long[n + 1];
        Mul = new long[n + 1];
        Mul[0] = 1;
        for (int i = 0; i < n; i++)
        {
            Hash[i + 1] = Hash[i] * Prime + (long) (sequence[i] - 'a');
            Mul[i + 1] = Mul[i] * Prime;
        }
        long hash = 0;
        for (int i = 0; i < m; i++)
        {
            hash = hash * Prime + (long) (word[i] - 'a');
        }
        int ans = 0;
        for (int i = 0; i < n; i++)
        {
            for (int k = 1; i + k * m <= n; k++)
            {
                if (GetHash(i + (k - 1) * m, i + k * m) != hash || sequence.Substring(i + (k - 1) * m, m) != word)
                {
                    break;
                }
                ans = Math.Max(ans, k);
            }
        }
        return ans;
    }
}
