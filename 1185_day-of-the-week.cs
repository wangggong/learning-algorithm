/*
 * @lc app=leetcode.cn id=day-of-the-week lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1185] 一周中的第几天
 *
 * https://leetcode.cn/problems/day-of-the-week/description/
 *
 * algorithms
 * Easy (62.04%)
 * Total Accepted:    42.1K
 * Total Submissions: 67.4K
 * Testcase Example:  '31\n8\n2019'
 *
 * 给你一个日期，请你设计一个算法来判断它是对应一周中的哪一天。
 * 
 * 输入为三个整数：day、month 和 year，分别表示日、月、年。
 * 
 * 您返回的结果必须是这几个值中的一个 {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday",
 * "Friday", "Saturday"}。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：day = 31, month = 8, year = 2019
 * 输出："Saturday"
 * 
 * 
 * 示例 2：
 * 
 * 输入：day = 18, month = 7, year = 1999
 * 输出："Sunday"
 * 
 * 
 * 示例 3：
 * 
 * 输入：day = 15, month = 8, year = 1993
 * 输出："Sunday"
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 给出的日期一定是在 1971 到 2100 年之间的有效日期。
 * 
 * 
 */
public class Solution
{
    private bool IsLeap(int y) => y % 4 == 0 && (y % 100 != 0 || y % 400 == 0);

    private int GetDays(int d, int m, int y) => (y - 1971) * 365
        + new int[] { 0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31, }[..m]
            .Sum()
        + d
        + Enumerable.Range(1971, y - 1971)
            .Count(IsLeap)
        + (IsLeap(y) && m > 2 ? 1 : 0);

    public string DayOfTheWeek(int day, int month, int year) => new string[]
    {
        "Sunday",
        "Monday",
        "Tuesday",
        "Wednesday",
        "Thursday",
        "Friday",
        "Saturday",
    }[(GetDays(day, month, year) + 4) % 7];
}
