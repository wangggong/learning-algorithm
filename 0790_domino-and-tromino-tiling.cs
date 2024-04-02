/*
 * @lc app=leetcode.cn id=domino-and-tromino-tiling lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [790] 多米诺和托米诺平铺
 *
 * https://leetcode.cn/problems/domino-and-tromino-tiling/description/
 *
 * algorithms
 * Medium (46.58%)
 * Total Accepted:    11.5K
 * Total Submissions: 21.9K
 * Testcase Example:  '3'
 *
 * 有两种形状的瓷砖：一种是 2 x 1 的多米诺形，另一种是形如 "L" 的托米诺形。两种形状都可以旋转。
 * 
 * 
 * 
 * 给定整数 n ，返回可以平铺 2 x n 的面板的方法的数量。返回对 10^9 + 7 取模 的值。
 * 
 * 平铺指的是每个正方形都必须有瓷砖覆盖。两个平铺不同，当且仅当面板上有四个方向上的相邻单元中的两个，使得恰好有一个平铺有一个瓷砖占据两个正方形。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 
 * 
 * 输入: n = 3
 * 输出: 5
 * 解释: 五种不同的方法如上所示。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: n = 1
 * 输出: 1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 1000
 * 
 * 
 */
public class Solution
{
    private const long Mod = (long) 1e9 + 7;

    public int NumTilings(int n)
    {
        long[,] dp = new long[3, 3];
        dp[1, 2] = 1;
        dp[2, 0] = 1;
        dp[2, 1] = 1;
        dp[2, 2] = 2;
        for (int i = 3; i <= n; i++)
        {
            dp[i % 3, 0] = (dp[(i - 2) % 3, 2] + dp[(i - 1) % 3, 1]) % Mod;
            dp[i % 3, 1] = (dp[(i - 2) % 3, 2] + dp[(i - 1) % 3, 0]) % Mod;
            dp[i % 3, 2] = (dp[(i - 2) % 3, 2] + dp[(i - 1) % 3, 0] + dp[(i - 1) % 3, 1] + dp[(i - 1) % 3, 2]) % Mod;

        }
        return (int) dp[n % 3, 2];
    }
}
