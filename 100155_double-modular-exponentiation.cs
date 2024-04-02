/*
 * @lc app=leetcode.cn id=double-modular-exponentiation lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100155] 双模幂运算
 *
 * https://leetcode.cn/problems/double-modular-exponentiation/description/
 *
 * algorithms
 * Medium (43.66%)
 * Total Accepted:    3.3K
 * Total Submissions: 7.5K
 * Testcase Example:  '[[2,3,3,10],[3,3,3,1],[6,1,1,4]]\n2'
 *
 * 给你一个下标从 0 开始的二维数组 variables ，其中 variables[i] = [ai, bi, ci, mi]，以及一个整数
 * target 。
 * 
 * 如果满足以下公式，则下标 i 是 好下标：
 * 
 * 
 * 0 <= i < variables.length
 * ((ai^bi % 10)^ci) % mi == target
 * 
 * 
 * 返回一个由 好下标 组成的数组，顺序不限 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：variables = [[2,3,3,10],[3,3,3,1],[6,1,1,4]], target = 2
 * 输出：[0,2]
 * 解释：对于 variables 数组中的每个下标 i ：
 * 1) 对于下标 0 ，variables[0] = [2,3,3,10] ，(2^3 % 10)^3 % 10 = 2 。
 * 2) 对于下标 1 ，variables[1] = [3,3,3,1] ，(3^3 % 10)^3 % 1 = 0 。
 * 3) 对于下标 2 ，variables[2] = [6,1,1,4] ，(6^1 % 10)^1 % 4 = 2 。
 * 因此，返回 [0,2] 作为答案。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：variables = [[39,3,1000,1000]], target = 17
 * 输出：[]
 * 解释：对于 variables 数组中的每个下标 i ：
 * 1) 对于下标 0 ，variables[0] = [39,3,1000,1000] ，(39^3 % 10)^1000 % 1000 = 1 。
 * 因此，返回 [] 作为答案。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= variables.length <= 100
 * variables[i] == [ai, bi, ci, mi]
 * 1 <= ai, bi, ci, mi <= 10^3
 * 0 <= target <= 10^3
 * 
 * 
 */
public class Solution
{
    private int Pow(int n, int k, int mod)
    {
        var ans = 1;
        for (; k > 0; k >>= 1)
        {
            if ((k & 1) is not 0) { ans = ans * n % mod; }
            n = n * n % mod;
        }
        return ans;
    }

    public IList<int> GetGoodIndices(int[][] variables, int target) => variables
        .Select((row, i) => (row, i))
        .Where(x => Pow(Pow(x.row[0], x.row[1], 10), x.row[2], x.row[3]) == target)
        .Select(x => x.i)
        .ToList();
}
