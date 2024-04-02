/*
 * @lc app=leetcode.cn id=maximum-employees-to-be-invited-to-a-meeting lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2127] 参加会议的最多员工数
 *
 * https://leetcode.cn/problems/maximum-employees-to-be-invited-to-a-meeting/description/
 *
 * algorithms
 * Hard (32.86%)
 * Total Accepted:    3.2K
 * Total Submissions: 9.9K
 * Testcase Example:  '[2,2,1,2]'
 *
 * 一个公司准备组织一场会议，邀请名单上有 n 位员工。公司准备了一张 圆形 的桌子，可以坐下 任意数目 的员工。
 * 
 * 员工编号为 0 到 n - 1 。每位员工都有一位 喜欢 的员工，每位员工 当且仅当 他被安排在喜欢员工的旁边，他才会参加会议。每位员工喜欢的员工 不会
 * 是他自己。
 * 
 * 给你一个下标从 0 开始的整数数组 favorite ，其中 favorite[i] 表示第 i 位员工喜欢的员工。请你返回参加会议的 最多员工数目
 * 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：favorite = [2,2,1,2]
 * 输出：3
 * 解释：
 * 上图展示了公司邀请员工 0，1 和 2 参加会议以及他们在圆桌上的座位。
 * 没办法邀请所有员工参与会议，因为员工 2 没办法同时坐在 0，1 和 3 员工的旁边。
 * 注意，公司也可以邀请员工 1，2 和 3 参加会议。
 * 所以最多参加会议的员工数目为 3 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：favorite = [1,2,0]
 * 输出：3
 * 解释：
 * 每个员工都至少是另一个员工喜欢的员工。所以公司邀请他们所有人参加会议的前提是所有人都参加了会议。
 * 座位安排同图 1 所示：
 * - 员工 0 坐在员工 2 和 1 之间。
 * - 员工 1 坐在员工 0 和 2 之间。
 * - 员工 2 坐在员工 1 和 0 之间。
 * 参与会议的最多员工数目为 3 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 
 * 输入：favorite = [3,0,1,4,1]
 * 输出：4
 * 解释：
 * 上图展示了公司可以邀请员工 0，1，3 和 4 参加会议以及他们在圆桌上的座位。
 * 员工 2 无法参加，因为他喜欢的员工 0 旁边的座位已经被占领了。
 * 所以公司只能不邀请员工 2 。
 * 参加会议的最多员工数目为 4 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == favorite.length
 * 2 <= n <= 10^5
 * 0 <= favorite[i] <= n - 1
 * favorite[i] != i
 * 
 * 
 */
// 参考: https://leetcode.cn/problems/maximum-employees-to-be-invited-to-a-meeting/solution/nei-xiang-ji-huan-shu-tuo-bu-pai-xu-fen-c1i1b/
public class Solution
{
    public int MaximumInvitations(int[] favorite)
    {
        var n = favorite.Length;
        var degree = new int[n];
        foreach (var f in favorite)
        {
            degree[f]++;
        }
        var dist = new int[n];
        var Q = new Queue<int>();
        for (var i = 0; i < n; i++)
        {
            if (degree[i] == 0)
            {
                Q.Enqueue(i);
            }
        }
        while (Q.Count > 0)
        {
            var u = Q.Dequeue();
            dist[u]++;
            var v = favorite[u];
            dist[v] = Math.Max(dist[v], dist[u]);
            degree[v]--;
            if (degree[v] == 0)
            {
                Q.Enqueue(v);
            }
        }
        var maxLoopLength = 0;
        var totPathLength = 0;
        for (var i = 0; i < n; i++)
        {
            if (degree[i] <= 0)
            {
                continue;
            }
            degree[i] = -1;
            var loopLength = 1;
            for (var j = favorite[i]; j != i; j = favorite[j])
            {
                degree[j] = -1;
                loopLength++;
            }
            if (loopLength == 2)
            {
                totPathLength += dist[i] + dist[favorite[i]] + 2;
            }
            else
            {
                maxLoopLength = Math.Max(maxLoopLength, loopLength);
            }
        }
        return Math.Max(maxLoopLength, totPathLength);
    }
}
