/*
 * @lc app=leetcode.cn id=zigzag-conversion lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6] Z 字形变换
 *
 * https://leetcode-cn.com/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (51.06%)
 * Total Accepted:    371.6K
 * Total Submissions: 722.5K
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
 *
 * 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
 *
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 * 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
 *
 * 请你实现这个将字符串进行指定行数变换的函数：
 *
 *
 * string convert(string s, int numRows);
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "PAYPALISHIRING", numRows = 3
 * 输出："PAHNAPLSIIGYIR"
 *
 * 示例 2：
 *
 *
 * 输入：s = "PAYPALISHIRING", numRows = 4
 * 输出："PINALSIGYAHRPI"
 * 解释：
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "A", numRows = 1
 * 输出："A"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 由英文字母（小写和大写）、',' 和 '.' 组成
 * 1 <= numRows <= 1000
 *
 *
 */

/*
 * P   A   H     N
 * A P L S I  I  G
 * Y   I   R
 *
 * 0   4   8     12
 * 1 3 5 7 9  11 13
 * 2   6   10
 */

const MAXN = 1000 + 10

var zigzag [MAXN][]byte

/*
 * func convert(s string, numRows int) string {
 * 	bs := []byte(s)
 * 	// n := len(bs)
 * 	for i := range zigzag {
 * 		zigzag[i] = nil
 * 	}
 * 	ind, direction := 0, 1
 * 	for _, b := range bs {
 * 		next := ind + direction
 * 		if next == 0 {
 * 			direction = 1
 * 		} else if next == numRows-1 {
 * 			direction = -1
 * 		}
 * 		zigzag[ind] = append(zigzag[ind], b)
 * 		ind = next
 * 	}
 * 	var ans []byte
 * 	for _, z := range zigzag {
 * 		ans = append(ans, z...)
 * 	}
 * 	return string(ans)
 * }
 */

func convert(s string, numRows int) string {
	// NOTE: 数学
	// T = 2*numRows - 2
	// - 第一行: T*k + 0
	// - 最后一行: T*k + (T-1)
	// - 中间 (对称的两个值): T*k + p & T*k + (T-p)
	bs := []byte(s)
	n := len(bs)
	if numRows == 1 || numRows >= n {
		return s
	}
	t := 2*numRows - 2
	var ans []byte
	for p := 0; p < numRows; p++ {
		for k := 0; k+p < n; k += t {
			ans = append(ans, bs[k+p])
			if p == 0 || p+1 == numRows {
				continue
			}
			if k+(t-p) < n {
				ans = append(ans, bs[k+(t-p)])
			}
		}
	}
	return string(ans)
}
