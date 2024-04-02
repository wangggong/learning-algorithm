/*
 * @lc app=leetcode.cn id=adding-two-negabinary-numbers lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1073] 负二进制数相加
 *
 * https://leetcode.cn/problems/adding-two-negabinary-numbers/description/
 *
 * algorithms
 * Medium (35.47%)
 * Total Accepted:    8.8K
 * Total Submissions: 21.7K
 * Testcase Example:  '[1,1,1,1,1]\n[1,0,1]'
 *
 * 给出基数为 -2 的两个数 arr1 和 arr2，返回两数相加的结果。
 * 
 * 数字以 数组形式 给出：数组由若干 0 和 1 组成，按最高有效位到最低有效位的顺序排列。例如，arr = [1,1,0,1] 表示数字 (-2)^3
 * + (-2)^2 + (-2)^0 = -3。数组形式 中的数字 arr 也同样不含前导零：即 arr == [0] 或 arr[0] == 1。
 * 
 * 返回相同表示形式的 arr1 和 arr2 相加的结果。两数的表示形式为：不含前导零、由若干 0 和 1 组成的数组。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：arr1 = [1,1,1,1,1], arr2 = [1,0,1]
 * 输出：[1,0,0,0,0]
 * 解释：arr1 表示 11，arr2 表示 5，输出表示 16 。
 * 
 * 
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：arr1 = [0], arr2 = [0]
 * 输出：[0]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：arr1 = [0], arr2 = [1]
 * 输出：[1]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 
 * 1 <= arr1.length, arr2.length <= 1000
 * arr1[i] 和 arr2[i] 都是 0 或 1
 * arr1 和 arr2 都没有前导0
 * 
 * 
 */
public class Solution
{
    public int[] AddNegabinary(int[] arr1, int[] arr2)
    {
        var list = new List<int>();
        var C = 0;
        for (var (i, n, m) = (0, arr1.Length, arr2.Length); i < Math.Max(n, m); i++)
        {
            var (p, q) = (i < n ? arr1[n - 1 - i] : 0, i < m ? arr2[m - 1 - i] : 0);
            var v = p + q + C;
            (v, C) = v switch
            {
                2 => (0, -1),
                3 => (1, -1),
                -1 => (1, 1),
                _ => (v, 0),
            };
            list.Add(v);
        }
        if (C != 0)
        {
            list.Add(1);
            if (C == -1)
            {
                list.Add(1);
            }
        }
        var ans = list.ToArray();
        for (var (i, j) = (0, ans.Length - 1); i < j; i++, j--)
        {
            (ans[i], ans[j]) = (ans[j], ans[i]);
        }
        var k = Array.IndexOf(ans, 1);
        return ans[(k < 0 ? ^1 : k) ..];
    }
}
