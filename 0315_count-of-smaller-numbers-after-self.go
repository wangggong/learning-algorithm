import "sort"

/*
 * @lc app=leetcode.cn id=count-of-smaller-numbers-after-self lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [315] 计算右侧小于当前元素的个数
 *
 * https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/description/
 *
 * algorithms
 * Hard (41.82%)
 * Total Accepted:    55K
 * Total Submissions: 131.5K
 * Testcase Example:  '[5,2,6,1]'
 *
 * 给你一个整数数组 nums ，按要求返回一个新数组 counts 。数组 counts 有该性质： counts[i] 的值是  nums[i]
 * 右侧小于 nums[i] 的元素的数量。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [5,2,6,1]
 * 输出：[2,1,1,0]
 * 解释：
 * 5 的右侧有 2 个更小的元素 (2 和 1)
 * 2 的右侧仅有 1 个更小的元素 (1)
 * 6 的右侧有 1 个更小的元素 (1)
 * 1 的右侧有 0 个更小的元素
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [-1]
 * 输出：[0]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [-1,-1]
 * 输出：[0,0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 */

func lowBit(x int) int { return x & -x }

type BIT []int

func (b BIT) query(index int) int {
	ans := 0
	for ; index > 0; index -= lowBit(index) {
		ans += b[index]
	}
	return ans
}

func (b BIT) add(index int, val int) {
	for ; index < len(b); index += lowBit(index) {
		b[index] += val
	}
}

func countSmaller(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	// 离散化
	m := make(map[int]int)
	for _, v := range nums {
		m[v] = 0
	}
	mn := len(m)
	var vals []int
	for v := range m {
		vals = append(vals, v)
	}
	sort.Ints(vals)
	for i, v := range vals {
		m[v] = i
	}
	for i, v := range nums {
		nums[i] = m[v]
	}
	// 转为计算前缀和
	tr := BIT(make([]int, mn+1))
	for i := n - 1; i >= 0; i-- {
		ans[i] = tr.query(nums[i])
		tr.add(nums[i]+1, 1)
	}
	return ans
}
