/*
 * @lc app=leetcode.cn id=array-of-doubled-pairs lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [954] 二倍数对数组
 *
 * https://leetcode-cn.com/problems/array-of-doubled-pairs/description/
 *
 * algorithms
 * Medium (30.91%)
 * Total Accepted:    15.7K
 * Total Submissions: 44.1K
 * Testcase Example:  '[3,1,3,6]'
 *
 * 给定一个长度为偶数的整数数组 arr，只有对 arr 进行重组后可以满足 “对于每个 0 <= i < len(arr) / 2，都有 arr[2 *
 * i + 1] = 2 * arr[2 * i]” 时，返回 true；否则，返回 false。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：arr = [3,1,3,6]
 * 输出：false
 *
 *
 * 示例 2：
 *
 *
 * 输入：arr = [2,1,2,6]
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：arr = [4,-2,2,-4]
 * 输出：true
 * 解释：可以用 [-2,-4] 和 [2,4] 这两组组成 [-2,-4,2,4] 或是 [2,4,-2,-4]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= arr.length <= 3 * 10^4
 * arr.length 是偶数
 * -10^5 <= arr[i] <= 10^5
 *
 *
 */
import "sort"

/*
 * func canReorderDoubled(arr []int) bool {
 * 	n := len(arr)
 * 	sort.Slice(arr, func(p, q int) bool { return abs(arr[p]) < abs(arr[q]) })
 * 	count := make(map[int][]int)
 * 	for i, a := range arr {
 * 		count[a] = append(count[a], i)
 * 	}
 * 	for i := 0; i < n; i++ {
 * 		p := count[arr[i]]
 * 		if len(p) == 0 {
 * 			continue
 * 		}
 * 		k := count[arr[i]*2]
 * 		if len(k) == 0 {
 * 			return false
 * 		}
 * 		count[arr[i]*2] = k[1:]
 * 		count[arr[i]] = p[1:]
 * 	}
 * 	return true
 * }
 */

func canReorderDoubled(arr []int) bool {
	count := make(map[int]int)
	for _, a := range arr {
		count[a]++
	}
	vals := make([]int, 0, len(arr))
	for v := range count {
		vals = append(vals, v)
	}
	sort.Slice(vals, func(p, q int) bool { return abs(vals[p]) < abs(vals[q]) })
	for _, v := range vals {
		if count[2*v] < count[v] {
			return false
		}
		count[2*v] -= count[v]
	}
	return true
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
