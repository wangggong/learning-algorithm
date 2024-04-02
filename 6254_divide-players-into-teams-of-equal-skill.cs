/*
 * @lc app=leetcode.cn id=divide-players-into-teams-of-equal-skill lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6254] 划分技能点相等的团队
 *
 * https://leetcode.cn/problems/divide-players-into-teams-of-equal-skill/description/
 *
 * algorithms
 * Medium (52.82%)
 * Total Accepted:    3.8K
 * Total Submissions: 7.3K
 * Testcase Example:  '[3,2,5,1,3,4]'
 *
 * 给你一个正整数数组 skill ，数组长度为 偶数 n ，其中 skill[i] 表示第 i 个玩家的技能点。将所有玩家分成 n / 2 个 2
 * 人团队，使每一个团队的技能点之和 相等 。
 * 
 * 团队的 化学反应 等于团队中玩家的技能点 乘积 。
 * 
 * 返回所有团队的 化学反应 之和，如果无法使每个团队的技能点之和相等，则返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：skill = [3,2,5,1,3,4]
 * 输出：22
 * 解释：
 * 将玩家分成 3 个团队 (1, 5), (2, 4), (3, 3) ，每个团队的技能点之和都是 6 。
 * 所有团队的化学反应之和是 1 * 5 + 2 * 4 + 3 * 3 = 5 + 8 + 9 = 22 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：skill = [3,4]
 * 输出：12
 * 解释：
 * 两个玩家形成一个团队，技能点之和是 7 。
 * 团队的化学反应是 3 * 4 = 12 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：skill = [1,1,2,3]
 * 输出：-1
 * 解释：
 * 无法将玩家分成每个团队技能点都相等的若干个 2 人团队。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= skill.length <= 10^5
 * skill.length 是偶数
 * 1 <= skill[i] <= 1000
 * 
 * 
 */
public class Solution
{
    public long DividePlayers(int[] skill)
    {
        var n = skill.Length;
        var tot = skill.Select(s => s).Sum();
        var target = tot / (n / 2);
        var count = new Dictionary<int, int>();
        long ans = 0;
        foreach (var s in skill)
        {
            if (count.ContainsKey(target - s))
            {
                count[target - s]--;
                if (count[target - s] == 0) { count.Remove(target - s); }
                ans += (long) s * (long) (target - s);
            }
            else
            {
                count[s] = (count.ContainsKey(s) ? count[s] : 0) + 1;
            }
        }
        if (count.Count() > 0) { return -1; }
        return ans;
    }
}
