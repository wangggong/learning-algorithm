/*
 * @lc app=leetcode.cn id=finding-the-users-active-minutes lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1817] 查找用户活跃分钟数
 *
 * https://leetcode.cn/problems/finding-the-users-active-minutes/description/
 *
 * algorithms
 * Medium (70.25%)
 * Total Accepted:    14.5K
 * Total Submissions: 18.9K
 * Testcase Example:  '[[0,5],[1,2],[0,2],[0,5],[1,3]]\n5'
 *
 * 给你用户在 LeetCode 的操作日志，和一个整数 k 。日志用一个二维整数数组 logs 表示，其中每个 logs[i] = [IDi,
 * timei] 表示 ID 为 IDi 的用户在 timei 分钟时执行了某个操作。
 * 
 * 多个用户 可以同时执行操作，单个用户可以在同一分钟内执行 多个操作 。
 * 
 * 指定用户的 用户活跃分钟数（user active minutes，UAM） 定义为用户对 LeetCode 执行操作的 唯一分钟数 。
 * 即使一分钟内执行多个操作，也只能按一分钟计数。
 * 
 * 请你统计用户活跃分钟数的分布情况，统计结果是一个长度为 k 且 下标从 1 开始计数 的数组 answer ，对于每个 j（1 ），answer[j]
 * 表示 用户活跃分钟数 等于 j 的用户数。
 * 
 * 返回上面描述的答案数组 answer 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：logs = [[0,5],[1,2],[0,2],[0,5],[1,3]], k = 5
 * 输出：[0,2,0,0,0]
 * 解释：
 * ID=0 的用户执行操作的分钟分别是：5 、2 和 5 。因此，该用户的用户活跃分钟数为 2（分钟 5 只计数一次）
 * ID=1 的用户执行操作的分钟分别是：2 和 3 。因此，该用户的用户活跃分钟数为 2
 * 2 个用户的用户活跃分钟数都是 2 ，answer[2] 为 2 ，其余 answer[j] 的值都是 0
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：logs = [[1,1],[2,2],[2,3]], k = 4
 * 输出：[1,1,0,0]
 * 解释：
 * ID=1 的用户仅在分钟 1 执行单个操作。因此，该用户的用户活跃分钟数为 1
 * ID=2 的用户执行操作的分钟分别是：2 和 3 。因此，该用户的用户活跃分钟数为 2
 * 1 个用户的用户活跃分钟数是 1 ，1 个用户的用户活跃分钟数是 2 
 * 因此，answer[1] = 1 ，answer[2] = 1 ，其余的值都是 0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 0 i 
 * 1 i 
 * k 的取值范围是 [用户的最大用户活跃分钟数, 10^5]
 * 
 * 
 */
public class Solution
{
    public int[] FindingUsersActiveMinutes(int[][] logs, int k)
    {
        var d = logs.GroupBy(x => x[0])
            .Select(x => x.Select(x => x[1]).Distinct().Count())
            .GroupBy(x => x)
            .ToDictionary(x => x.Key, x => x.Count());
        return Enumerable.Range(1, k).Select(x =>
        {
            d.TryGetValue(x, out var count);
            return count;
        }).ToArray();
    }
}
