/*
 * @lc app=leetcode.cn id=number-of-days-between-two-dates lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1360] 日期之间隔几天
 *
 * https://leetcode-cn.com/problems/number-of-days-between-two-dates/description/
 *
 * algorithms
 * Easy (50.90%)
 * Total Accepted:    9.7K
 * Total Submissions: 19.1K
 * Testcase Example:  '"2019-06-29"\n"2019-06-30"'
 *
 * 请你编写一个程序来计算两个日期之间隔了多少天。
 *
 * 日期以字符串形式给出，格式为 YYYY-MM-DD，如示例所示。
 *
 *
 *
 * 示例 1：
 *
 * 输入：date1 = "2019-06-29", date2 = "2019-06-30"
 * 输出：1
 *
 *
 * 示例 2：
 *
 * 输入：date1 = "2020-01-15", date2 = "2019-12-31"
 * 输出：15
 *
 *
 *
 *
 * 提示：
 *
 *
 * 给定的日期是 1971 年到 2100 年之间的有效日期。
 *
 *
 */

var monthDays = [12 + 1]int{
	0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31,
}

func daysBetweenDates(date1 string, date2 string) int {
	d1, d2 := days(date1), days(date2)
	return ans(d1 - d2)
}

// days returns days diff between 1970-01-01 and the date.
func days(date string) int {
	bytes := []byte(date)
	year, month, day := bytes[:4], bytes[5:7], bytes[8:]
	y, m, d := itoa(year), itoa(month), itoa(day)
	ans := 0
	for i := 1970; i < y; i++ {
		ans += 365
		if isLeap(i) {
			ans++
		}
	}
	for i := 1; i < m; i++ {
		ans += monthDays[i]
		if i == 2 && isLeap(y) {
			ans++
		}
	}
	ans += d
	return ans
}

func itoa(arr []byte) int {
	ans := 0
	for _, a := range arr {
		ans = ans*10 + int(a-'0')
	}
	return ans
}

func ans(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func isLeap(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}
