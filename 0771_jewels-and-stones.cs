/*
 * @lc app=leetcode.cn id=jewels-and-stones lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [771] 宝石与石头
 *
 * https://leetcode.cn/problems/jewels-and-stones/description/
 *
 * algorithms
 * Easy (85.12%)
 * Total Accepted:    191.2K
 * Total Submissions: 223.9K
 * Testcase Example:  '"aA"\n"aAAbbbb"'
 *
 *  给你一个字符串 jewels 代表石头中宝石的类型，另有一个字符串 stones 代表你拥有的石头。 stones
 * 中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。
 * 
 * 字母区分大小写，因此 "a" 和 "A" 是不同类型的石头。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：jewels = "aA", stones = "aAAbbbb"
 * 输出：3
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：jewels = "z", stones = "ZZ"
 * 输出：0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= jewels.length, stones.length <= 50
 * jewels 和 stones 仅由英文字母组成
 * jewels 中的所有字符都是 唯一的
 * 
 * 
 */
public class Solution
{
    public int NumJewelsInStones(string jewels, string stones) =>
        stones.Count(s => jewels.Contains(s));
}
