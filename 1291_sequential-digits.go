/*
 * @lc app=leetcode.cn id=sequential-digits lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1291] 顺次数
 *
 * https://leetcode-cn.com/problems/sequential-digits/description/
 *
 * algorithms
 * Medium (52.65%)
 * Total Accepted:    8.3K
 * Total Submissions: 15.7K
 * Testcase Example:  '100\n300'
 *
 * 我们定义「顺次数」为：每一位上的数字都比前一位上的数字大 1 的整数。
 *
 * 请你返回由 [low, high] 范围内所有顺次数组成的 有序 列表（从小到大排序）。
 *
 *
 *
 * 示例 1：
 *
 * 输出：low = 100, high = 300
 * 输出：[123,234]
 *
 *
 * 示例 2：
 *
 * 输出：low = 1000, high = 13000
 * 输出：[1234,2345,3456,4567,5678,6789,12345]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 10 <= low <= high <= 10^9
 *
 *
 */
func sequentialDigits(low int, high int) []int {
	// assert 10 <= low && low <= high && high <= 1e9;
	arr := buildSequentialDigits()
	n := len(arr)
	var r, s int
	{
		// 0 0 0 0 0 1 1 1
		p, q := 0, n
		for p < q {
			mid := (p + q) >> 1
			if arr[mid] >= low {
				q = mid
			} else {
				p = mid + 1
			}
		}
		r = p
	}
	{
		// 1 1 1 1 1 0 0 0
		//
		// 傻逼! 二分模型记反了!
		p, q := -1, n-1
		for p < q {
			mid := (p + q + 1) >> 1
			if arr[mid] > high {
				q = mid - 1
			} else {
				p = mid
			}
		}
		s = p
	}
	return arr[r : s+1]
}

func buildSequentialDigits() []int {
	var arr, ans []int
	for i := 1; i < 10; i++ {
		ans = append(ans, i)
	}
	for i := 2; i <= 9; i++ {
		k := 0
		for i, a := range ans {
			v := a % 10
			if v == 9 {
				break
			}
			ans[i] = a*10 + (v + 1)
			k = i + 1
		}
		arr = append(arr, ans[:k]...)
	}
	return arr
}
