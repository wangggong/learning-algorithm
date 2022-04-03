/*
 * @lc app=leetcode.cn id=circular-array-loop lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [457] 环形数组是否存在循环
 *
 * https://leetcode-cn.com/problems/circular-array-loop/description/
 *
 * algorithms
 * Medium (43.58%)
 * Total Accepted:    30.3K
 * Total Submissions: 69.7K
 * Testcase Example:  '[2,-1,1,2,2]'
 *
 * 存在一个不含 0 的 环形 数组 nums ，每个 nums[i] 都表示位于下标 i 的角色应该向前或向后移动的下标个数：
 *
 *
 * 如果 nums[i] 是正数，向前（下标递增方向）移动 |nums[i]| 步
 * 如果 nums[i] 是负数，向后（下标递减方向）移动 |nums[i]| 步
 *
 *
 * 因为数组是 环形 的，所以可以假设从最后一个元素向前移动一步会到达第一个元素，而第一个元素向后移动一步会到达最后一个元素。
 *
 * 数组中的 循环 由长度为 k 的下标序列 seq 标识：
 *
 *
 * 遵循上述移动规则将导致一组重复下标序列 seq[0] -> seq[1] -> ... -> seq[k - 1] -> seq[0] ->
 * ...
 * 所有 nums[seq[j]] 应当不是 全正 就是 全负
 * k > 1
 *
 *
 * 如果 nums 中存在循环，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,-1,1,2,2]
 * 输出：true
 * 解释：存在循环，按下标 0 -> 2 -> 3 -> 0 。循环长度为 3 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [-1,2]
 * 输出：false
 * 解释：按下标 1 -> 1 -> 1 ... 的运动无法构成循环，因为循环的长度为 1 。根据定义，循环的长度必须大于 1 。
 *
 *
 * 示例 3:
 *
 *
 * 输入：nums = [-2,1,-1,-2,-2]
 * 输出：false
 * 解释：按下标 1 -> 2 -> 1 -> ... 的运动无法构成循环，因为 nums[1] 是正数，而 nums[2] 是负数。
 * 所有 nums[seq[j]] 应当不是全正就是全负。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5000
 * -1000 <= nums[i] <= 1000
 * nums[i] != 0
 *
 *
 *
 *
 * 进阶：你能设计一个时间复杂度为 O(n) 且额外空间复杂度为 O(1) 的算法吗？
 *
 */

func circularArrayLoop(nums []int) bool {
	n := len(nums)
	// assert n > 0 && n < 5e3;
	// assert nums[i] >= -1e3 && nums[i] <= 1e3 && nums[i] != 0;
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			continue
		}
		p, q := i, next(nums, i)
		for nums[p]*nums[q] > 0 && nums[p]*nums[next(nums, q)] > 0 {
			if p == q {
				if p == next(nums, p) {
					break
				}
				return true
			}
			p, q = next(nums, p), next(nums, next(nums, q))
		}
		// 发现不是成环的之后反手把所有的 index 都置为 0 (保证原数据不为 0)
		for index := i; nums[index]*nums[next(nums, index)] > 0; index = next(nums, index) {
			nums[index] = 0
		}
	}
	return false
}

/*
 * func circularArrayLoop(nums []int) bool {
 * 	n := len(nums)
 * 	// assert n > 0 && n <= 5e3;
 * 	// assert nums[i] >= -1e3 && nums[i] <= 1e3;
 * 	visit := make([]int, n)
 * 	index := 1
 * 	for i := 0; i < n; i++ {
 * 		if nums[i] == 0 {
 * 			continue
 * 		}
 * 		visit[i] = index
 * 		for true {
 * 			j := next(nums, i)
 * 			if i == j {
 * 				break
 * 			}
 * 			if v := visit[j]; v != 0 {
 * 				if v != index {
 * 					break
 * 				} else {
 * 					return true
 * 				}
 * 			}
 * 			if nums[i]*nums[j] < 0 {
 * 				break
 * 			}
 * 			visit[j] = index
 * 			i = j
 * 		}
 * 		index++
 * 	}
 * 	return false
 * }
 */

func next(arr []int, index int) int {
	n := len(arr)
	// assert index >= 0 && index < n;
	// assert next(arr, index) >= 0 && next(arr, index) < n;
	return ((index+arr[index])%n + n) % n
}
