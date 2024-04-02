/*
 * @lc app=leetcode.cn id=time-needed-to-inform-all-employees lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1376] 通知所有员工所需的时间
 *
 * https://leetcode.cn/problems/time-needed-to-inform-all-employees/description/
 *
 * algorithms
 * Medium (54.81%)
 * Total Accepted:    19.4K
 * Total Submissions: 34.1K
 * Testcase Example:  '1\n0\n[-1]\n[0]'
 *
 * 公司里有 n 名员工，每个员工的 ID 都是独一无二的，编号从 0 到 n - 1。公司的总负责人通过 headID 进行标识。
 * 
 * 在 manager 数组中，每个员工都有一个直属负责人，其中 manager[i] 是第 i
 * 名员工的直属负责人。对于总负责人，manager[headID] = -1。题目保证从属关系可以用树结构显示。
 * 
 * 公司总负责人想要向公司所有员工通告一条紧急消息。他将会首先通知他的直属下属们，然后由这些下属通知他们的下属，直到所有的员工都得知这条紧急消息。
 * 
 * 第 i 名员工需要 informTime[i] 分钟来通知它的所有直属下属（也就是说在 informTime[i]
 * 分钟后，他的所有直属下属都可以开始传播这一消息）。
 * 
 * 返回通知所有员工这一紧急消息所需要的 分钟数 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 1, headID = 0, manager = [-1], informTime = [0]
 * 输出：0
 * 解释：公司总负责人是该公司的唯一一名员工。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 
 * 
 * 输入：n = 6, headID = 2, manager = [2,2,-1,2,2,2], informTime = [0,0,1,0,0,0]
 * 输出：1
 * 解释：id = 2 的员工是公司的总负责人，也是其他所有员工的直属负责人，他需要 1 分钟来通知所有员工。
 * 上图显示了公司员工的树结构。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 10^5
 * 0 <= headID < n
 * manager.length == n
 * 0 <= manager[i] < n
 * manager[headID] == -1
 * informTime.length == n
 * 0 <= informTime[i] <= 1000
 * 如果员工 i 没有下属，informTime[i] == 0 。
 * 题目 保证 所有员工都可以收到通知。
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     public int NumOfMinutes(int n, int headId, int[] manager, int[] informTime)
 *     {
 *         var G = Enumerable.Range(0, n).Select(_ => new List<int>()).ToArray();
 *         for (var i = 0; i < n; i++)
 *         {
 *             if (i != headId)
 *             {
 *                 G[manager[i]].Add(i);
 *             }
 *         }
 *         var ans = 0;
 *         void dfs(int u, int t)
 *         {
 *             ans = Math.Max(ans, t);
 *             foreach (var v in G[u])
 *             {
 *                 dfs(v, t + informTime[u]);
 *             }
 *         }
 *         dfs(headId, 0);
 *         return ans;
 *     }
 * }
 */

public class Solution
{
    public int NumOfMinutes(int n, int headID, int[] manager, int[] informTime)
        => Enumerable.Range(0, n)
            .Select(k => Dfs(manager, informTime, k))
            .Max();

    private int Dfs(int[] manager, int[] informTime, int k)
        => manager[k] < 0
            ? informTime[k]
            : ((manager[k], informTime[k]) = (-1,
                informTime[k] + Dfs(manager, informTime, manager[k]))).Item2;
}
