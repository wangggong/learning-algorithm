/*
 * @lc app=leetcode.cn id=sum-of-numbers-with-units-digit-k lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [5218] 个位数字为 K 的整数之和
 *
 * https://leetcode.cn/problems/sum-of-numbers-with-units-digit-k/description/
 *
 * algorithms
 * Medium (20.27%)
 * Total Accepted:    3.8K
 * Total Submissions: 18.8K
 * Testcase Example:  '58\n9'
 *
 * 给你两个整数 num 和 k ，考虑具有以下属性的正整数多重集：
 *
 *
 * 每个整数个位数字都是 k 。
 * 所有整数之和是 num 。
 *
 *
 * 返回该多重集的最小大小，如果不存在这样的多重集，返回 -1 。
 *
 * 注意：
 *
 *
 * 多重集与集合类似，但多重集可以包含多个同一整数，空多重集的和为 0 。
 * 个位数字 是数字最右边的数位。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：num = 58, k = 9
 * 输出：2
 * 解释：
 * 多重集 [9,49] 满足题目条件，和为 58 且每个整数的个位数字是 9 。
 * 另一个满足条件的多重集是 [19,39] 。
 * 可以证明 2 是满足题目条件的多重集的最小长度。
 *
 *
 * 示例 2：
 *
 *
 * 输入：num = 37, k = 2
 * 输出：-1
 * 解释：个位数字为 2 的整数无法相加得到 37 。
 *
 * 示例 3：
 *
 *
 * 输入：num = 0, k = 7
 * 输出：0
 * 解释：空多重集的和为 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= num <= 3000
 * 0 <= k <= 9
 *
 *
 */
func minimumNumbers(num int, k int) int {
	if num == 0 {
		return 0
	}
	for i := 1; i <= 10; i++ {
		if v := k * i; v%10 == num%10 && v <= num {
			return i
		}
	}
	return -1
}
