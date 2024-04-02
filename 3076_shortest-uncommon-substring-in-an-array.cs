/*
 * @lc app=leetcode.cn id=shortest-uncommon-substring-in-an-array lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [3076] 数组中的最短非公共子字符串
 *
 * https://leetcode.cn/problems/shortest-uncommon-substring-in-an-array/description/
 *
 * algorithms
 * Medium (45.44%)
 * Total Accepted:    3.5K
 * Total Submissions: 7.6K
 * Testcase Example:  '["cab","ad","bad","c"]'
 *
 * 给你一个数组 arr ，数组中有 n 个 非空 字符串。
 * 
 * 请你求出一个长度为 n 的字符串 answer ，满足：
 * 
 * 
 * answer[i] 是 arr[i] 最短 的子字符串，且它不是 arr 中其他任何字符串的子字符串。如果有多个这样的子字符串存在，answer[i]
 * 应该是它们中字典序最小的一个。如果不存在这样的子字符串，answer[i] 为空字符串。
 * 
 * 
 * 请你返回数组 answer 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：arr = ["cab","ad","bad","c"]
 * 输出：["ab","","ba",""]
 * 解释：求解过程如下：
 * - 对于字符串 "cab" ，最短没有在其他字符串中出现过的子字符串是 "ca" 或者 "ab" ，我们选择字典序更小的子字符串，也就是 "ab" 。
 * - 对于字符串 "ad" ，不存在没有在其他字符串中出现过的子字符串。
 * - 对于字符串 "bad" ，最短没有在其他字符串中出现过的子字符串是 "ba" 。
 * - 对于字符串 "c" ，不存在没有在其他字符串中出现过的子字符串。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：arr = ["abc","bcd","abcd"]
 * 输出：["","","abcd"]
 * 解释：求解过程如下：
 * - 对于字符串 "abc" ，不存在没有在其他字符串中出现过的子字符串。
 * - 对于字符串 "bcd" ，不存在没有在其他字符串中出现过的子字符串。
 * - 对于字符串 "abcd" ，最短没有在其他字符串中出现过的子字符串是 "abcd" 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == arr.length
 * 2 <= n <= 100
 * 1 <= arr[i].length <= 20
 * arr[i] 只包含小写英文字母。
 * 
 * 
 */
public class Solution
{
    public string[] ShortestSubstrings(string[] arr)
    {
        var sets = arr.Select(s => Enumerable.Range(0, s.Length)
                .SelectMany(i => Enumerable.Range(i + 1, s.Length - i)
                    .Select(j => s[i..j]))
                .ToHashSet())
            .ToList();
        return sets.Select((S, i) => S.OrderBy(s => s.Length)
                .ThenBy(s => s)
                .FirstOrDefault(s => Enumerable.Range(0, arr.Length)
                    .All(j => i == j || !sets[j].Contains(s)), ""))
            .ToArray();
    }
}
