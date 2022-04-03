/*
 * @lc app=leetcode.cn id=single-number-iii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [260] 只出现一次的数字 III
 *
 * https://leetcode-cn.com/problems/single-number-iii/description/
 *
 * algorithms
 * Medium (73.97%)
 * Total Accepted:    83.3K
 * Total Submissions: 112.7K
 * Testcase Example:  '[1,2,1,3,2,5]'
 *
 * 给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。
 *
 *
 *
 * 进阶：你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,1,3,2,5]
 * 输出：[3,5]
 * 解释：[5, 3] 也是有效的答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [-1,0]
 * 输出：[-1,0]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [0,1]
 * 输出：[1,0]
 *
 *
 * 提示：
 *
 *
 * 2 <= nums.length <= 3 * 1e4
 * -(1 << 31) <= nums[i] <= (1 << 31) - 1
 * 除两个只出现一次的整数外，nums 中的其他数字都出现两次
 *
 *
 */

// NOTE:
// - 先对所有数异或. 这步我还记得;
// - 异或结果就等于 a ^ b;
// - 找到结果为 '1' 的一位 (一定存在, 不存在的话说明 a == b, 矛盾);
// - 重点: 然后考虑该位, 为 '1' 的有 2p+1 个, 为 '0' 的有 2q+1 个, 退化成问题 1.
func singleNumber(nums []int) []int {
	ans := []int{0, 0}
	v := 0
	for i := range nums {
		nums[i] += 1 << 31
	}
	for _, n := range nums {
		v = v ^ n
	}
	index := 0
	for ; index < 32; index++ {
		if v&(1<<index) != 0 {
			break
		}
	}
	for _, n := range nums {
		i := (n >> index) & 1
		ans[i] = ans[i] ^ n
	}
	ans[0], ans[1] = ans[0]-(1<<31), ans[1]-(1<<31)
	return ans
}
