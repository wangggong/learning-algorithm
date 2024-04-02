/*
 * @lc app=leetcode.cn id=find-a-peak-element-ii lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1901] 寻找峰值 II
 *
 * https://leetcode.cn/problems/find-a-peak-element-ii/description/
 *
 * algorithms
 * Medium (58.22%)
 * Total Accepted:    9.1K
 * Total Submissions: 15.4K
 * Testcase Example:  '[[1,4],[3,2]]'
 *
 * 一个 2D 网格中的 峰值 是指那些 严格大于 其相邻格子(上、下、左、右)的元素。
 * 
 * 给你一个 从 0 开始编号 的 m x n 矩阵 mat ，其中任意两个相邻格子的值都 不相同 。找出 任意一个 峰值 mat[i][j] 并
 * 返回其位置 [i,j] 。
 * 
 * 你可以假设整个矩阵周边环绕着一圈值为 -1 的格子。
 * 
 * 要求必须写出时间复杂度为 O(m log(n)) 或 O(n log(m)) 的算法
 * 
 * 
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 
 * 
 * 输入: mat = [[1,4],[3,2]]
 * 输出: [0,1]
 * 解释: 3 和 4 都是峰值，所以[1,0]和[0,1]都是可接受的答案。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 
 * 
 * 输入: mat = [[10,20,15],[21,30,14],[7,16,32]]
 * 输出: [1,1]
 * 解释: 30 和 32 都是峰值，所以[1,1]和[2,2]都是可接受的答案。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == mat.length
 * n == mat[i].length
 * 1 <= m, n <= 500
 * 1 <= mat[i][j] <= 10^5
 * 任意两个相邻元素均不相等.
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/find-a-peak-element-ii/solutions/2571587/tu-jie-li-yong-xing-zui-da-zhi-pan-duan-r4e0n/?envType=daily-question&envId=2023-12-19
//
// 巧妙的二分. 重点是找中间列的最大值.
public class Solution
{
    public int[] FindPeakGrid(int[][] M)
    {
        var (n, m) = (M.Length, M[0].Length);
        int get(int i, int j) => 0 <= i && i < n && 0 <= j && j < m
            ? M[i][j]
            : -1;
        for (var (p, q) = (0, m); true; )
        {
            var mid = (p + q) >> 1;
            var r = -1;
            for (var i = 0; i < n; i++)
            {
                if (get(i, mid) > get(r, mid)) { r = i; }
            }
            if (get(r, mid) > get(r, mid - 1) && get(r, mid) > get(r, mid + 1))
            { return new int[] { r, mid, }; }
            if (get(r, mid) < get(r, mid - 1)) { q = mid; }
            else { p = mid; }
        }
    }
}
