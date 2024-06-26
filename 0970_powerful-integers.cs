/*
 * @lc app=leetcode.cn id=powerful-integers lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [970] 强整数
 *
 * https://leetcode.cn/problems/powerful-integers/description/
 *
 * algorithms
 * Medium (41.23%)
 * Total Accepted:    16.4K
 * Total Submissions: 38K
 * Testcase Example:  '2\n3\n10'
 *
 * 给定三个整数 x 、 y 和 bound ，返回 值小于或等于 bound 的所有 强整数 组成的列表 。
 * 
 * 如果某一整数可以表示为 x^i + y^j ，其中整数 i >= 0 且 j >= 0，那么我们认为该整数是一个 强整数 。
 * 
 * 你可以按 任何顺序 返回答案。在你的回答中，每个值 最多 出现一次。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：x = 2, y = 3, bound = 10
 * 输出：[2,3,4,5,7,9,10]
 * 解释： 
 * 2 = 2^0 + 3^0
 * 3 = 2^1 + 3^0
 * 4 = 2^0 + 3^1
 * 5 = 2^1 + 3^1
 * 7 = 2^2 + 3^1
 * 9 = 2^3 + 3^0
 * 10 = 2^0 + 3^2
 * 
 * 示例 2：
 * 
 * 
 * 输入：x = 3, y = 5, bound = 15
 * 输出：[2,4,6,8,10,14]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= x, y <= 100
 * 0 <= bound <= 10^6
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     public IList<int> PowerfulIntegers(int x, int y, int bound)
 *     {
 *         if (x == 1)
 *         {
 *             (x, y) = (y, x);
 *             if (x == 1)
 *             {
 *                 var ans = new List<int>();
 *                 if (bound >= 2)
 *                 {
 *                     ans.Add(2);
 *                 }
 *                 return ans;
 *             }
 *         }
 *         var visit = new bool[bound + 1];
 *         for (var i = 1; i < i * x && i <= bound; i *= x)
 *         {
 *             for (var j = 1; i + j <= bound; j *= y)
 *             {
 *                 visit[i + j] = true;
 *                 if (y == 1)
 *                 {
 *                     break;
 *                 }
 *             }
 *         }
 *         return Enumerable.Range(1, bound).Where(i => visit[i]).ToList();
 *     }
 * }
 */

// 注意到 `log(1e6) <= 21`, 枚举 21 次即可.
public class Solution
{
    public IList<int> PowerfulIntegers(int x, int y, int bound)
    {
        var S = new HashSet<int>();
        for (var (i, p) = (0, 1); i <= 21; i++, p *= x)
        {
            for (var (j, q) = (0, 1); j <= 21; j++, q *= y)
            {
                if (p + q <= bound)
                {
                    S.Add((int)(p + q));
                }
                if (q > bound)
                {
                    break;
                }
            }
            if (p > bound)
            {
                break;
            }
        }
        return S.ToList();
    }
}
