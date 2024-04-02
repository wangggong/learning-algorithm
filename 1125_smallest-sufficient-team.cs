/*
 * @lc app=leetcode.cn id=smallest-sufficient-team lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1125] 最小的必要团队
 *
 * https://leetcode.cn/problems/smallest-sufficient-team/description/
 *
 * algorithms
 * Hard (51.10%)
 * Total Accepted:    7.7K
 * Total Submissions: 13.6K
 * Testcase Example:  '["java","nodejs","reactjs"]\n[["java"],["nodejs"],["nodejs","reactjs"]]'
 *
 * 作为项目经理，你规划了一份需求的技能清单 req_skills，并打算从备选人员名单 people 中选出些人组成一个「必要团队」（ 编号为 i
 * 的备选人员 people[i] 含有一份该备选人员掌握的技能列表）。
 * 
 * 所谓「必要团队」，就是在这个团队中，对于所需求的技能列表 req_skills
 * 中列出的每项技能，团队中至少有一名成员已经掌握。可以用每个人的编号来表示团队中的成员：
 * 
 * 
 * 例如，团队 team = [0, 1, 3] 表示掌握技能分别为 people[0]，people[1]，和 people[3] 的备选人员。
 * 
 * 
 * 请你返回 任一 规模最小的必要团队，团队成员用人员编号表示。你可以按 任意顺序 返回答案，题目数据保证答案存在。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：req_skills = ["java","nodejs","reactjs"], people =
 * [["java"],["nodejs"],["nodejs","reactjs"]]
 * 输出：[0,2]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：req_skills = ["algorithms","math","java","reactjs","csharp","aws"],
 * people =
 * [["algorithms","math","java"],["algorithms","math","reactjs"],["java","csharp","aws"],["reactjs","csharp"],["csharp","math"],["aws","java"]]
 * 输出：[1,2]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 1 
 * req_skills[i] 由小写英文字母组成
 * req_skills 中的所有字符串 互不相同
 * 1 
 * 0 
 * 1 
 * people[i][j] 由小写英文字母组成
 * people[i] 中的所有字符串 互不相同
 * people[i] 中的每个技能是 req_skills 中的技能
 * 题目数据保证「必要团队」一定存在
 * 
 * 
 */
public class Solution
{
    public int[] SmallestSufficientTeam(string[] skills, IList<IList<string>> people)
	{
		var n = skills.Length;
		var S = 1 << n;
		var d = Enumerable.Range(0, n).ToDictionary(i => skills[i], i => i);
		var m = people.Count();
		var dp = new int[S];
		var prev = new (int, int)[S];
		Array.Fill(dp, m + 1);
		Array.Fill(prev, (-1, -1));
		dp[0] = 0;
		for (var i = 0; i < m; i++)
		{
			var p = people[i];
			var cur = 0;
			foreach (var skill in p)
			{
				cur |= 1 << d[skill];
			}
			for (var s = S - 1; s >= 0; s--)
			{
				var ns = s | cur;
				if (dp[ns] > dp[s] + 1)
				{
					dp[ns] = dp[s] + 1;
					prev[ns] = (s, i);
				}
			}
		}
		var ans = new List<int>();
		for (var c = S - 1; c != -1; c = prev[c].Item1)
		{
			ans.Add(prev[c].Item2);
		}
		return ans.OrderBy(x => x).Skip(1).ToArray();
	}
}
