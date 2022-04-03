/*
 * @lc app=leetcode.cn id=day-of-the-year lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1154] 一年中的第几天
 *
 * https://leetcode-cn.com/problems/day-of-the-year/description/
 *
 * algorithms
 * Easy (64.66%)
 * Total Accepted:    45.8K
 * Total Submissions: 70.8K
 * Testcase Example:  '"2019-01-09"'
 *
 * 给你一个字符串 date ，按 YYYY-MM-DD 格式表示一个 现行公元纪年法 日期。返回该日期是当年的第几天。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：date = "2019-01-09"
 * 输出：9
 * 解释：给定日期是2019年的第九天。
 *
 * 示例 2：
 *
 *
 * 输入：date = "2019-02-10"
 * 输出：41
 *
 *
 *
 *
 * 提示：
 *
 *
 * date.length == 10
 * date[4] == date[7] == '-'，其他的 date[i] 都是数字
 * date 表示的范围从 1900 年 1 月 1 日至 2019 年 12 月 31 日
 *
 *
 */
var monthDays = [12 + 1]int{
	0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31,
}

func dayOfYear(date string) int {
	// assert len(date) == 0 && date[4] == '-' && date[7] == '-';
	bytes := []byte(date)
	year, month, day := bytes[:4], bytes[5:7], bytes[8:]
	y, m, d := atoi(year), atoi(month), atoi(day)
	ans := 0
	for i := 1; i < m; i++ {
		ans += monthDays[i]
	}
	ans += d
	if m > 2 && ((y%4 == 0 && y%100 != 0) || y%400 == 0) {
		ans++
	}
	return ans
}

func atoi(arr []byte) int {
	v := 0
	for _, a := range arr {
		v *= 10
		v += int(a - '0')
	}
	return v
}
