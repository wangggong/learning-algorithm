/*
 * @lc app=leetcode.cn id=range-sum-query-immutable lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [303] 区域和检索 - 数组不可变
 *
 * https://leetcode-cn.com/problems/range-sum-query-immutable/description/
 *
 * algorithms
 * Easy (73.54%)
 * Total Accepted:    143.8K
 * Total Submissions: 195.5K
 * Testcase Example:  '["NumArray","sumRange","sumRange","sumRange"]\n' +
  '[[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]'
 *
 * 给定一个整数数组  nums，处理以下类型的多个查询:
 *
 *
 * 计算索引 left 和 right （包含 left 和 right）之间的 nums 元素的 和 ，其中 left <= right
 *
 *
 * 实现 NumArray 类：
 *
 *
 * NumArray(int[] nums) 使用数组 nums 初始化对象
 * int sumRange(int i, int j) 返回数组 nums 中索引 left 和 right 之间的元素的 总和 ，包含 left 和
 * right 两点（也就是 nums[left] + nums[left + 1] + ... + nums[right] )
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：
 * ["NumArray", "sumRange", "sumRange", "sumRange"]
 * [[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
 * 输出：
 * [null, 1, -1, -3]
 *
 * 解释：
 * NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
 * numArray.sumRange(0, 2); // return 1 ((-2) + 0 + 3)
 * numArray.sumRange(2, 5); // return -1 (3 + (-5) + 2 + (-1))
 * numArray.sumRange(0, 5); // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= nums.length <= 10^4
 * -10^5 <= nums[i] <= 10^5
 * 0 <= i <= j < nums.length
 * 最多调用 10^4 次 sumRange 方法
 *
 *
*/
type NumArray struct {
	Nums      []int
	PrefixSum []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i, n := range nums {
		prefixSum[i+1] = prefixSum[i] + n
	}
	return NumArray{Nums: nums, PrefixSum: prefixSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	// assert left >= 0 && left < right && right < len(this.Nums);
	// X	-2	0	3	-5	2	-1
	// 0	-2	-2	1	-4	-2	-3
	return this.PrefixSum[right+1] - this.PrefixSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
