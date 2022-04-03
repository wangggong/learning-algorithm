/*
 * @lc app=leetcode.cn id=range-sum-query-mutable lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [307] 区域和检索 - 数组可修改
 *
 * https://leetcode-cn.com/problems/range-sum-query-mutable/description/
 *
 * algorithms
 * Medium (52.75%)
 * Total Accepted:    27K
 * Total Submissions: 51.1K
 * Testcase Example:  '["NumArray","sumRange","update","sumRange"]\n[[[1,3,5]],[0,2],[1,2],[0,2]]'
 *
 * 给你一个数组 nums ，请你完成两类查询。
 *
 *
 * 其中一类查询要求 更新 数组 nums 下标对应的值
 * 另一类查询要求返回数组 nums 中索引 left 和索引 right 之间（ 包含 ）的nums元素的 和 ，其中 left <= right
 *
 *
 * 实现 NumArray 类：
 *
 *
 * NumArray(int[] nums) 用整数数组 nums 初始化对象
 * void update(int index, int val) 将 nums[index] 的值 更新 为 val
 * int sumRange(int left, int right) 返回数组 nums 中索引 left 和索引 right 之间（ 包含
 * ）的nums元素的 和 （即，nums[left] + nums[left + 1], ..., nums[right]）
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：
 * ["NumArray", "sumRange", "update", "sumRange"]
 * [[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]
 * 输出：
 * [null, 9, null, 8]
 *
 * 解释：
 * NumArray numArray = new NumArray([1, 3, 5]);
 * numArray.sumRange(0, 2); // 返回 1 + 3 + 5 = 9
 * numArray.update(1, 2);   // nums = [1,2,5]
 * numArray.sumRange(0, 2); // 返回 1 + 2 + 5 = 8
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 3 * 10^4
 * -100 <= nums[i] <= 100
 * 0 <= index < nums.length
 * -100 <= val <= 100
 * 0 <= left <= right < nums.length
 * 调用 pdate 和 sumRange 方法次数不大于 3 * 10^4
 *
 *
 */
// 树状数组
// 参考: https://leetcode-cn.com/problems/range-sum-query-mutable/solution/guan-yu-ge-lei-qu-jian-he-wen-ti-ru-he-x-41hv/

// 应用场景总结如下:
// - 数组不变，求区间和：「前缀和」、「树状数组」、「线段树」
// - 多次修改某个数，求区间和：「树状数组」、「线段树」
// - 多次整体修改某个区间，求区间和：「线段树」、「树状数组」（看修改区间的数据范围）
// - 多次将某个区间变成同一个数，求区间和：「线段树」、「树状数组」（看修改区间的数据范围）

func lowBit(x int) int { return x & -x }

type NumArray struct {
	Size int
	Nums []int
	Tree []int
}

func Constructor(nums []int) NumArray {
	size := len(nums)
	tree := make([]int, size+1)
	na := NumArray{Size: size, Nums: nums, Tree: tree}
	for i, n := range nums {
		na.add(i+1, n)
	}
	return na
}

// ["NumArray","sumRange","update","sumRange"]\n[[[-1]],[0,0],[0,1],[0,0]]
func (this *NumArray) query(x int) int {
	ans := 0
	for i := x; i > 0; i -= lowBit(i) {
		// fmt.Printf("query: %v %v\n", i, this.Tree[i])
		ans += this.Tree[i]
	}
	return ans
}

func (this *NumArray) add(x, v int) {
	for i := x; i <= this.Size; i += lowBit(i) {
		// fmt.Printf("add: %v %v\n", i, this.Tree[i])
		this.Tree[i] += v
	}
}

func (this *NumArray) Update(index int, val int) {
	// assert index < this.Size;
	this.add(index+1, val-this.Nums[index])
	this.Nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	// assert left >= 0 && left < right && right < len(this.Nums);
	return this.query(right+1) - this.query(left)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
