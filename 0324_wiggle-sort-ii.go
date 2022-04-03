/*
 * @lc app=leetcode.cn id=wiggle-sort-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [324] 摆动排序 II
 *
 * https://leetcode-cn.com/problems/wiggle-sort-ii/description/
 *
 * algorithms
 * Medium (38.28%)
 * Total Accepted:    30.3K
 * Total Submissions: 79.1K
 * Testcase Example:  '[1,5,1,1,6,4]'
 *
 * 给你一个整数数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。
 *
 * 你可以假设所有输入数组都可以得到满足题目要求的结果。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,5,1,1,6,4]
 * 输出：[1,6,1,5,1,4]
 * 解释：[1,4,1,5,1,6] 同样是符合题目要求的结果，可以被判题程序接受。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,3,2,2,3,1]
 * 输出：[2,3,1,3,1,2]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5 * 1e4
 * 0 <= nums[i] <= 5000
 * 题目数据保证，对于给定的输入 nums ，总能产生满足题目要求的结果
 *
 *
 *
 *
 * 进阶：你能用 O(n) 时间复杂度和 / 或原地 O(1) 额外空间来实现吗？
 *
 */

/*
 * func wiggleSort(nums []int) {
 * 	sort.Ints(nums)
 * 	n := len(nums)
 * 	var tmp []int
 * 	start := (n + 1) >> 1
 * 	p, q := start-1, n-1
 * 	// fmt.Printf("%v %v %v %v\n", nums, p, q, start)
 * 	for ; p >= 0 || q >= start; p, q = p-1, q-1 {
 * 		if p >= 0 {
 * 			tmp = append(tmp, nums[p])
 * 		}
 * 		if q >= start {
 * 			tmp = append(tmp, nums[q])
 * 		}
 * 	}
 * 	copy(nums, tmp)
 * 	return
 * }
 */

// 参考: https://leetcode-cn.com/problems/wiggle-sort-ii/solution/yi-bu-yi-bu-jiang-shi-jian-fu-za-du-cong-onlognjia/
func wiggleSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	K := (n + 1) >> 1
	quickSelect(nums, K)
	v := nums[K]
	// 3-way partition
	// 荷兰国旗?
	// 给定数放在中间, 小于给定数的放在左边, 大于给定数的放在右边
	i, j, k := 0, 0, n-1
	for j <= k {
		switch c := nums[j] - v; {
		case c > 0:
			nums[j], nums[k] = nums[k], nums[j]
			k--
		case c < 0:
			nums[j], nums[i] = nums[i], nums[j]
			i++
			j++
		case c == 0:
			j++
		default:
			// unreachable
		}
	}
	var tmp []int
	left, right := nums[:K], nums[K:]
	p, q := len(left)-1, len(right)-1
	for p >= 0 || q >= 0 {
		if p >= 0 {
			tmp = append(tmp, left[p])
			p--
		}
		if q >= 0 {
			tmp = append(tmp, right[q])
			q--
		}
	}
	// fmt.Printf("%v %v %v %v\n", nums, left, right, tmp)
	copy(nums, tmp)
	return
}

func quickSelect(arr []int, k int) {
	n := len(arr)
	p, q := 0, n-1
	for p < q {
		ind := partition(arr, p, q)
		if ind == k {
			return
		}
		if ind < k {
			p = ind + 1
		} else {
			q = ind - 1
		}
	}
}

func partition(arr []int, p, q int) int {
	v := arr[p]
	start := p
	for i := p + 1; i <= q; i++ {
		if arr[i] < v {
			start++
			arr[start], arr[i] = arr[i], arr[start]
		}
	}
	arr[start], arr[p] = arr[p], arr[start]
	return start
}
