/*
 * Medium
 * 
 * 给你一个整数 n 和一个下标从 0 开始的字符串数组 words ，和一个下标从 0 开始的数组 groups ，两个数组长度都是 n 。
 * 
 * 两个长度相等字符串的 汉明距离 定义为对应位置字符 不同 的数目。
 * 
 * 你需要从下标 [0, 1, ..., n - 1] 中选出一个 最长子序列 ，将这个子序列记作长度为 k 的 [i0, i1, ..., ik - 1] ，它需要满足以下条件：
 * 
 * 相邻 下标对应的 groups 值 不同。即，对于所有满足 0 < j + 1 < k 的 j 都有 groups[ij] != groups[ij + 1] 。
 * 对于所有 0 < j + 1 < k 的下标 j ，都满足 words[ij] 和 words[ij + 1] 的长度 相等 ，且两个字符串之间的 汉明距离 为 1 。
 * 请你返回一个字符串数组，它是下标子序列 依次 对应 words 数组中的字符串连接形成的字符串数组。如果有多个答案，返回任意一个。
 * 
 * 子序列 指的是从原数组中删掉一些（也可能一个也不删掉）元素，剩余元素不改变相对位置得到的新的数组。
 * 
 * 注意：words 中的字符串长度可能 不相等 。
 * 
 *  
 * 
 * 示例 1：
 * 
 * 输入：n = 3, words = ["bab","dab","cab"], groups = [1,2,2]
 * 输出：["bab","cab"]
 * 解释：一个可行的子序列是 [0,2] 。
 * - groups[0] != groups[2]
 * - words[0].length == words[2].length 且它们之间的汉明距离为 1 。
 * 所以一个可行的答案是 [words[0],words[2]] = ["bab","cab"] 。
 * 另一个可行的子序列是 [0,1] 。
 * - groups[0] != groups[1]
 * - words[0].length = words[1].length 且它们之间的汉明距离为 1 。
 * 所以另一个可行的答案是 [words[0],words[1]] = ["bab","dab"] 。
 * 符合题意的最长子序列的长度为 2 。
 * 示例 2：
 * 
 * 输入：n = 4, words = ["a","b","c","d"], groups = [1,2,3,4]
 * 输出：["a","b","c","d"]
 * 解释：我们选择子序列 [0,1,2,3] 。
 * 它同时满足两个条件。
 * 所以答案为 [words[0],words[1],words[2],words[3]] = ["a","b","c","d"] 。
 * 它是所有下标子序列里最长且满足所有条件的。
 * 所以它是唯一的答案。
 *  
 * 
 * 提示：
 * 
 * 1 <= n == words.length == groups.length <= 1000
 * 1 <= words[i].length <= 10
 * 1 <= groups[i] <= n
 * words 中的字符串 互不相同 。
 * words[i] 只包含小写英文字母。
 * 
 * 通过次数 1.8K
 * 提交次数 5.4K
 * 通过率 33.9%
 * 
 */

public class Solution
{
    public IList<string> GetWordsInLongestSubsequence(int n, string[] words, int[] groups)
    {
        bool check(string s, string t) => s.Length == t.Length
            && s.Zip(t, (s, t) => (s, t))
                .Where(x => x.s != x.t)
                .Count() == 1;
        var ans = new List<string>();
        var dp = new int[n];
        var last = new int[n];
        Array.Fill(last, -1);
        for (var i = 0; i < n; i++)
        {
            dp[i] = 1;
            for (var j = i - 1; j >= 0; j--)
            {
                if (groups[i] != groups[j] && check(words[i], words[j]) && dp[i] < dp[j] + 1)
                {
                    dp[i] = dp[j] + 1;
                    last[i] = j;
                }
            }
        }
        var k = dp.Select((d, i) => (d, i))
            .OrderByDescending(x => x.d)
            .First().i;
        for (; k is not -1; k = last[k]) { ans.Add(words[k]); }
        ans.Reverse();
        return ans;
    }
}
