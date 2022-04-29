/*
 * @lc app=leetcode.cn id=random-pick-index lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [398] 随机数索引
 *
 * https://leetcode-cn.com/problems/random-pick-index/description/
 *
 * algorithms
 * Medium (67.23%)
 * Total Accepted:    29.1K
 * Total Submissions: 40.8K
 * Testcase Example:  '["Solution","pick","pick","pick"]\n[[[1,2,3,3,3]],[3],[1],[3]]'
 *
 * 给定一个可能含有重复元素的整数数组，要求随机输出给定的数字的索引。 您可以假设给定的数字一定存在于数组中。
 *
 * 注意：
 * 数组大小可能非常大。 使用太多额外空间的解决方案将不会通过测试。
 *
 * 示例:
 *
 *
 * int[] nums = new int[] {1,2,3,3,3};
 * Solution solution = new Solution(nums);
 *
 * // pick(3) 应该返回索引 2,3 或者 4。每个索引的返回概率应该相等。
 * solution.pick(3);
 *
 * // pick(1) 应该返回 0。因为只有nums[0]等于1。
 * solution.pick(1);
 *
 *
 */
import "math/rand"

/*
 * type Solution map[int][]int
 *
 * func Constructor(nums []int) Solution {
 * 	s := make(map[int][]int)
 * 	for i, n := range nums {
 * 		s[n] = append(s[n], i)
 * 	}
 * 	return Solution(s)
 * }
 *
 * func (s *Solution) Pick(target int) int {
 * 	// assert len(s[target]) > 0;
 * 	arr := (*s)[target]
 * 	n := len(arr)
 * 	return arr[rand.Intn(n)]
 * }
 */

// 水塘抽样
type Solution []int

func Constructor(nums []int) Solution {
	return Solution(nums)
}

func (s *Solution) Pick(target int) int {
	cnt := 0
	ans := -1
	for i, v := range *s {
		if v == target {
			cnt++
			if ans < 0 {
				ans = i
			} else {
				if rand.Intn(cnt) == 0 {
					ans = i
				}
			}
		}
	}
	return ans
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
