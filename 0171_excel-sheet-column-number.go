/*
 * @lc app=leetcode.cn id=excel-sheet-column-number lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [171] Excel 表列序号
 *
 * https://leetcode-cn.com/problems/excel-sheet-column-number/description/
 *
 * algorithms
 * Easy (71.74%)
 * Total Accepted:    122.5K
 * Total Submissions: 171K
 * Testcase Example:  '"A"'
 *
 * 给你一个字符串 columnTitle ，表示 Excel 表格中的列名称。返回 该列名称对应的列序号 。
 * 
 * 例如：
 * 
 * 
 * A -> 1
 * B -> 2
 * C -> 3
 * ...
 * Z -> 26
 * AA -> 27
 * AB -> 28 
 * ...
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入: columnTitle = "A"
 * 输出: 1
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: columnTitle = "AB"
 * 输出: 28
 * 
 * 
 * 示例 3:
 * 
 * 
 * 输入: columnTitle = "ZY"
 * 输出: 701
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= columnTitle.length <= 7
 * columnTitle 仅由大写英文组成
 * columnTitle 在范围 ["A", "FXSHRXW"] 内
 * 
 * 
 */
func titleToNumber(s string) int {
	ans := 0
	for _, b := range s {
		ans = ans * 26 + int(b - 'A') + 1
	}
	return ans
}
