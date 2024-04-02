/*
 * @lc app=leetcode.cn id=count-vowel-strings-in-ranges lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6347] 统计范围内的元音字符串数
 *
 * https://leetcode.cn/problems/count-vowel-strings-in-ranges/description/
 *
 * algorithms
 * Medium (48.61%)
 * Total Accepted:    3.3K
 * Total Submissions: 6.7K
 * Testcase Example:  '["aba","bcb","ece","aa","e"]\n[[0,2],[1,4],[1,1]]'
 *
 * 给你一个下标从 0 开始的字符串数组 words 以及一个二维整数数组 queries 。
 * 
 * 每个查询 queries[i] = [li, ri] 会要求我们统计在 words 中下标在 li 到 ri 范围内（包含
 * 这两个值）并且以元音开头和结尾的字符串的数目。
 * 
 * 返回一个整数数组，其中数组的第 i 个元素对应第 i 个查询的答案。
 * 
 * 注意：元音字母是 'a'、'e'、'i'、'o' 和 'u' 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：words = ["aba","bcb","ece","aa","e"], queries = [[0,2],[1,4],[1,1]]
 * 输出：[2,3,0]
 * 解释：以元音开头和结尾的字符串是 "aba"、"ece"、"aa" 和 "e" 。
 * 查询 [0,2] 结果为 2（字符串 "aba" 和 "ece"）。
 * 查询 [1,4] 结果为 3（字符串 "ece"、"aa"、"e"）。
 * 查询 [1,1] 结果为 0 。
 * 返回结果 [2,3,0] 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：words = ["a","e","i"], queries = [[0,2],[0,1],[2,2]]
 * 输出：[3,2,1]
 * 解释：每个字符串都满足这一条件，所以返回 [3,2,1] 。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= words.length <= 10^5
 * 1 <= words[i].length <= 40
 * words[i] 仅由小写英文字母组成
 * sum(words[i].length) <= 3 * 10^5
 * 1 <= queries.length <= 10^5
 * 0 <= queries[j][0] <= queries[j][1] < words.length
 * 
 * 
 */
public class Solution
{
    public int[] VowelStrings(string[] words, int[][] queries)
    {
        bool isVowel(char s) => s is 'a' or 'e' or 'i' or 'o' or 'u';
        var n = words.Length;
        var S = new int[n + 1];
        for (var i = 0; i < n; i++)
        {
            S[i + 1] = S[i] + (isVowel(words[i][0]) && isVowel(words[i][^1]) ? 1 : 0);
        }
        return queries.Select(q => S[q[1] + 1] - S[q[0]]).ToArray();
    }
}
