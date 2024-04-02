/*
 * Easy
 * 
 * 泰波那契序列 Tn 定义如下： 
 * 
 * T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2
 * 
 * 给你整数 n，请返回第 n 个泰波那契数 Tn 的值。
 * 
 *  
 * 
 * 示例 1：
 * 
 * 输入：n = 4
 * 输出：4
 * 解释：
 * T_3 = 0 + 1 + 1 = 2
 * T_4 = 1 + 1 + 2 = 4
 * 示例 2：
 * 
 * 输入：n = 25
 * 输出：1389537
 *  
 * 
 * 提示：
 * 
 * 0 <= n <= 37
 * 答案保证是一个 32 位整数，即 answer <= 2^31 - 1。
 * 通过次数 174.1K
 * 提交次数 285.6K
 * 通过率 61.0%
 * 
 */
public class Solution
{
    public int Tribonacci(int n)
    {
        if (n < 3) { return n is 0 ? 0 : 1; }
        var dp = new int[n + 1];
        (dp[0], dp[1], dp[2]) = (0, 1, 1);
        for (var i = 3; i <= n; i++) { dp[i] = dp[i - 1] + dp[i - 2] + dp[i - 3]; }
        return dp.Last();
    }
}

/*
 * public class Solution
 * {
 *     public int Tribonacci(int n) => n < 3
 *         ? (n is 0 ? 0 : 1)
 *         : Mul(Pow(new int[][]
 *         {
 *             new int[] { 1, 1, 1, },
 *             new int[] { 1, 0, 0, },
 *             new int[] { 0, 1, 0, },
 *         }, n - 2), new int[] { 1, 1, 0, })[0];
 * 
 *     private int[][] Pow(int[][] M, int k)
 *     {
 *         var n = M.Length;
 *         var ans = new int[n][];
 *         for (var i = 0; i < n; i++)
 *         {
 *             ans[i] = new int[n];
 *             ans[i][i] = 1;
 *         }
 *         for (; k > 0; k >>= 1)
 *         {
 *             if ((k & 1) is not 0) { ans = Mul(M, ans); }
 *             M = Mul(M, M);
 *         }
 *         return ans;
 *     }
 * 
 *     private int[][] Mul(int[][] A, int[][] B)
 *     {
 *         var n = A.Length;
 *         var ans = new int[n][];
 *         for (var i = 0; i < n; i++) { ans[i] = new int[n]; }
 *         for (var i = 0; i < n; i++)
 *         {
 *             for (var j = 0; j < n; j++)
 *             {
 *                 for (var k = 0; k < n; k++) { ans[i][j] += A[i][k] * B[k][j]; }
 *             }
 *         }
 *         return ans;
 *     }
 * 
 *     private int[] Mul(int[][] M, int[] v)
 *     {
 *         var n = M.Length;
 *         var ans = new int[n];
 *         for (var i = 0; i < n; i++)
 *         {
 *             for (var j = 0; j < n; j++) { ans[i] += M[i][j] * v[j]; }
 *         }
 *         return ans;
 *     }
 * }
 */
