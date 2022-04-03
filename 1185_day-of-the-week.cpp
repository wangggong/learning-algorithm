/*
 * @lc app=leetcode.cn id=day-of-the-week lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1185] 一周中的第几天
 *
 * https://leetcode-cn.com/problems/day-of-the-week/description/
 *
 * algorithms
 * Easy (62.67%)
 * Total Accepted:    34.6K
 * Total Submissions: 55.2K
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
class Solution {
  string weekdays[7] = {
    "Sunday",
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday"
  };
  int monthDays[12+1] = {
    0,
    31,
    28,
    31,
    30,
    31,
    30,
    31,
    31,
    30,
    31,
    30,
    31
  };
public:
  string dayOfTheWeek(int d, int m, int y) {
    auto ans = 4;
    for (auto i = 1971; i < y; i++) {
      auto isLeap = i % 4 == 0 && (i % 100 != 0 || i % 400 == 0);
      ans += 365;
      ans += isLeap ? 1 : 0;
      // cout << "y" << i << " " << ans << endl;
    }
    for (auto i = 1; i < m; i++) {
      ans += monthDays[i];
      auto isLeap = y % 4 == 0 && (y % 100 != 0 || y % 400 == 0);
      ans += (i == 2 && isLeap) ? 1 : 0;
      // cout << "m" << i << " " << ans << endl;
    }
    ans += d;
    // cout << ans << endl;
    return weekdays[ans % 7];
  }
};
