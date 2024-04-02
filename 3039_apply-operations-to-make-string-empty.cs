/*
 * @lc app=leetcode.cn id=apply-operations-to-make-string-empty lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [3039] 进行操作使字符串为空
 *
 * https://leetcode.cn/problems/apply-operations-to-make-string-empty/description/
 *
 * algorithms
 * Medium (58.48%)
 * Total Accepted:    1.5K
 * Total Submissions: 2.6K
 * Testcase Example:  '"aabcbbca"'
 *
 * 给你一个字符串 s 。
 * 
 * 请你进行以下操作直到 s 为 空 ：
 * 
 * 
 * 每次操作 依次 遍历 'a' 到 'z'，如果当前字符出现在 s 中，那么删除出现位置 最早 的该字符。
 * 
 * 
 * 请你返回进行 最后 一次操作 之前 的字符串 s 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "aabcbbca"
 * 输出："ba"
 * 解释：我们进行以下操作：
 * - 删除 s = "aabcbbca" 中加粗加斜字符，得到字符串 s = "abbca" 。
 * - 删除 s = "abbca" 中加粗加斜字符，得到字符串 s = "ba" 。
 * - 删除 s = "ba" 中加粗加斜字符，得到字符串 s = "" 。
 * 进行最后一次操作之前的字符串为 "ba" 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "abcd"
 * 输出："abcd"
 * 解释：我们进行以下操作：
 * - 删除 s = "abcd" 中加粗加斜字符，得到字符串 s = "" 。
 * 进行最后一次操作之前的字符串为 "abcd" 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 5 * 10^5
 * s 只包含小写英文字母。
 * 
 * 
 */
public class Solution
{
    public string LastNonEmptyString(string s)
    {
        var d = s
            .Select((c, i) => (c, i))
            .GroupBy(x => x.c)
            .ToDictionary(g => g.Key, g => (c: g.Count(), i: g.Select(x => x.i)
                .Last()));
        var max = d
            .Select(kv => kv.Value.c)
            .Max();
        return new string(d
            .Where(kv => kv.Value.c == max)
            .OrderBy(kv => kv.Value.i)
            .Select(kv => kv.Key)
            .ToArray());

    }
}
