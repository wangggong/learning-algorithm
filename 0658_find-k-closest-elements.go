/*
 * @lc app=leetcode.cn id=find-k-closest-elements lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [658] 找到 K 个最接近的元素
 *
 * https://leetcode.cn/problems/find-k-closest-elements/description/
 *
 * algorithms
 * Medium (45.92%)
 * Total Accepted:    68.2K
 * Total Submissions: 142.6K
 * Testcase Example:  '[1,2,3,4,5]\n4\n3'
 *
 * 给定一个 排序好 的数组 arr ，两个整数 k 和 x ，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。
 *
 * 整数 a 比整数 b 更接近 x 需要满足：
 *
 *
 * |a - x| < |b - x| 或者
 * |a - x| == |b - x| 且 a < b
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：arr = [1,2,3,4,5], k = 4, x = 3
 * 输出：[1,2,3,4]
 *
 *
 * 示例 2：
 *
 *
 * 输入：arr = [1,2,3,4,5], k = 4, x = -1
 * 输出：[1,2,3,4]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= arr.length
 * 1 <= arr.length <= 10^4
 * arr 按 升序 排列
 * -10^4 <= arr[i], x <= 10^4
 *
 *
 */

/*
 * import "sort"
 *
 * func findClosestElements(arr []int, k int, x int) []int {
 * 	sort.Slice(arr, func(p, q int) bool {
 * 		return abs(arr[p]-x) < abs(arr[q]-x) || (abs(arr[p]-x) == abs(arr[q]-x) && arr[p] < arr[q])
 * 	})
 * 	ans := make([]int, k)
 * 	copy(ans, arr[:k])
 * 	sort.Ints(ans)
 * 	return ans
 * }
 *
 */

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	p, q := 0, n-1
	for p < q {
		mid := (p + q + 1) >> 1
		if arr[mid] <= x {
			p = mid
		} else {
			q = mid - 1
		}
	}
	if p+1 < n && abs(arr[p+1]-x) < abs(arr[p]-x) {
		p++
	}
	p, q = p-1, p+1
	for q-p-1 < k {
		if q >= n {
			p--
			continue
		}
		if p < 0 {
			q++
			continue
		}
		if abs(arr[q]-x) < abs(arr[p]-x) {
			q++
		} else {
			p--
		}
	}
	return arr[p+1 : q]
}
