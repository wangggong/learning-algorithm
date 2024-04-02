/*
 * @lc app=leetcode.cn id=maximum-tastiness-of-candy-basket lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6271] 礼盒的最大甜蜜度
 *
 * https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/description/
 *
 * algorithms
 * Medium (56.94%)
 * Total Accepted:    1K
 * Total Submissions: 1.7K
 * Testcase Example:  '[13,5,1,8,21,2]\n3'
 *
 * 给你一个正整数数组 price ，其中 price[i] 表示第 i 类糖果的价格，另给你一个正整数 k 。
 * 
 * 商店组合 k 类 不同 糖果打包成礼盒出售。礼盒的 甜蜜度 是礼盒中任意两种糖果 价格 绝对差的最小值。
 * 
 * 返回礼盒的 最大 甜蜜度。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：price = [13,5,1,8,21,2], k = 3
 * 输出：8
 * 解释：选出价格分别为 [13,5,21] 的三类糖果。
 * 礼盒的甜蜜度为 min(|13 - 5|, |13 - 21|, |5 - 21|) = min(8, 8, 16) = 8 。
 * 可以证明能够取得的最大甜蜜度就是 8 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：price = [1,3,1], k = 2
 * 输出：2
 * 解释：选出价格分别为 [1,3] 的两类糖果。 
 * 礼盒的甜蜜度为 min(|1 - 3|) = min(2) = 2 。
 * 可以证明能够取得的最大甜蜜度就是 2 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：price = [7,7,7,7], k = 2
 * 输出：0
 * 解释：从现有的糖果中任选两类糖果，甜蜜度都会是 0 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= price.length <= 10^5
 * 1 <= price[i] <= 10^9
 * 2 <= k <= price.length
 * 
 * 
 */
public class Solution
{
    public int MaximumTastiness(int[] price, int k)
    {
        var n = price.Length;
        Array.Sort(price);
        var p = 0;
        var q = (int) 1e9 + 1;
        bool check(int v)
        {
            for (int p = 0, q = 0, rest = k - 1; p < n; p = q)
            {
                for (q = p; q < n && price[q] < price[p] + v; q++) { }
                if (q >= n)
                {
                    break;
                }
                rest--;
                if (rest == 0)
                {
                    return true;
                }
            }
            return false;
        }
        while (p < q)
        {
            var mid = (p + q) >> 1;
            if (!check(mid))
            {
                q = mid;
            }
            else
            {
                p = mid + 1;
            }
        }
        return p - 1;
    }
}
