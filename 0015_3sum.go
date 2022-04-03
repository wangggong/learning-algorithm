/*
 * @lc app=leetcode.cn id=3sum lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (34.46%)
 * Total Accepted:    795K
 * Total Submissions: 2.3M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0
 * 且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-1,0,1,2,-1,-4]
 * 输出：[[-1,-1,2],[-1,0,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [0]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= nums.length <= 3000
 * -1e5 <= nums[i] <= 1e5
 *
 *
 */
import "sort"

type _3Pair struct{ x, y, z int }

/*
 * func threeSum(nums []int) [][]int {
 * 	n := len(nums)
 * 	m := make(map[int]int)
 * 	mp := make(map[_3Pair]struct{})
 * 	for i, v := range nums {
 * 		m[v] = i
 * 	}
 * 	var ans [][]int
 * 	for i := 0; i < n; i++ {
 * 		for j := i + 1; j < n; j++ {
 * 			v, ok := m[-nums[i]-nums[j]]
 * 			if !ok {
 * 				continue
 * 			}
 * 			if v == i || v == j {
 * 				continue
 * 			}
 * 			a := []int{nums[i], nums[j], nums[v]}
 * 			sort.Ints(a)
 * 			p := _3Pair{a[0], a[1], a[2]}
 * 			if _, ok := mp[p]; ok {
 * 				continue
 * 			}
 * 			mp[p] = struct{}{}
 * 			ans = append(ans, a)
 * 		}
 * 	}
 * 	return ans
 * }
 */

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int
	for i := 0; i+2 < n; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		p, q := i+1, n-1
		for p < q {
			if p > i+1 && nums[p-1] == nums[p] {
				p++
				continue
			}
			switch c := nums[p] + nums[q] + nums[i]; {
			case c == 0:
				ans = append(ans, []int{nums[p], nums[q], nums[i]})
				p, q = p+1, q-1
			case c > 0:
				q--
			case c < 0:
				p++
			default:
				// unreachable
			}
		}
	}
	return ans
}
