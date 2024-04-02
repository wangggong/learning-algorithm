/*
 * @lc app=leetcode.cn id=substring-xor-queries lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6356] 子字符串异或查询
 *
 * https://leetcode.cn/problems/substring-xor-queries/description/
 *
 * algorithms
 * Medium (30.60%)
 * Total Accepted:    2.6K
 * Total Submissions: 8.6K
 * Testcase Example:  '"101101"\n[[0,5],[1,2]]'
 *
 * 给你一个 二进制字符串 s 和一个整数数组 queries ，其中 queries[i] = [firsti, secondi] 。
 * 
 * 对于第 i 个查询，找到 s 的 最短子字符串 ，它对应的 十进制值 val 与 firsti 按位异或 得到 secondi ，换言之，val ^
 * firsti == secondi 。
 * 
 * 第 i 个查询的答案是子字符串 [lefti, righti] 的两个端点（下标从 0 开始），如果不存在这样的子字符串，则答案为 [-1, -1]
 * 。如果有多个答案，请你选择 lefti 最小的一个。
 * 
 * 请你返回一个数组 ans ，其中 ans[i] = [lefti, righti] 是第 i 个查询的答案。
 * 
 * 子字符串 是一个字符串中一段连续非空的字符序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "101101", queries = [[0,5],[1,2]]
 * 输出：[[0,2],[2,3]]
 * 解释：第一个查询，端点为 [0,2] 的子字符串为 "101" ，对应十进制数字 5 ，且 5 ^ 0 = 5 ，所以第一个查询的答案为
 * [0,2]。第二个查询中，端点为 [2,3] 的子字符串为 "11" ，对应十进制数字 3 ，且 3 ^ 1 = 2 。所以第二个查询的答案为
 * [2,3] 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "0101", queries = [[12,8]]
 * 输出：[[-1,-1]]
 * 解释：这个例子中，没有符合查询的答案，所以返回 [-1,-1] 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：s = "1", queries = [[4,5]]
 * 输出：[[0,0]]
 * 解释：这个例子中，端点为 [0,0] 的子字符串对应的十进制值为 1 ，且 1 ^ 4 = 5 。所以答案为 [0,0] 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 10^4
 * s[i] 要么是 '0' ，要么是 '1' 。
 * 1 <= queries.length <= 10^5
 * 0 <= firsti, secondi <= 10^9
 * 
 * 
 * 
 * 
 */
public class Solution
{
    public int[][] SubstringXorQueries(string s, int[][] queries)
    {
        var n = s.Length;
        var memo = new Dictionary<long, int[]>();
        for (var i = 0; i < n; i++)
        {
            if (s[i] == '0')
            {
                if (!memo.ContainsKey(0))
                {
                    memo[0] = new[] { i, i, };
                }
                continue;
            }
            long cur = 0;
            for (var j = 0; i + j < n; j++)
            {
                cur = (cur << 1) | (s[i + j] == '0' ? 0 : 1);
                if (!memo.ContainsKey(cur))
                {
                    memo[cur] = new[] { i, i + j };
                }
            }
        }
        return queries.Select(q =>
        {
            var target = (long)(q[0] ^ q[1]);
            return memo.ContainsKey(target) ? memo[target] : new[] { -1, -1, };
        }).ToArray();
    }
}
