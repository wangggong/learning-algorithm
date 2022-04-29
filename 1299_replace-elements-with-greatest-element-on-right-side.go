/*
 * @lc app=leetcode.cn id=replace-elements-with-greatest-element-on-right-side lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1299] 将每个元素替换为右侧最大元素
 *
 * https://leetcode-cn.com/problems/replace-elements-with-greatest-element-on-right-side/description/
 *
 * algorithms
 * Easy (78.69%)
 * Total Accepted:    33.7K
 * Total Submissions: 42.8K
 * Testcase Example:  '[17,18,5,4,6,1]'
 *
 * 给你一个数组 arr ，请你将每个元素用它右边最大的元素替换，如果是最后一个元素，用 -1 替换。
 *
 * 完成所有替换操作后，请你返回这个数组。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：arr = [17,18,5,4,6,1]
 * 输出：[18,6,6,6,1,-1]
 * 解释：
 * - 下标 0 的元素 --> 右侧最大元素是下标 1 的元素 (18)
 * - 下标 1 的元素 --> 右侧最大元素是下标 4 的元素 (6)
 * - 下标 2 的元素 --> 右侧最大元素是下标 4 的元素 (6)
 * - 下标 3 的元素 --> 右侧最大元素是下标 4 的元素 (6)
 * - 下标 4 的元素 --> 右侧最大元素是下标 5 的元素 (1)
 * - 下标 5 的元素 --> 右侧没有其他元素，替换为 -1
 *
 *
 * 示例 2：
 *
 *
 * 输入：arr = [400]
 * 输出：[-1]
 * 解释：下标 0 的元素右侧没有其他元素。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= arr.length <= 1e4
 * 1 <= arr[i] <= 1e5
 *
 *
 */
func replaceElements(arr []int) []int {
	n := len(arr)
	S := make([]int, n)
	S[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		S[i] = S[i+1]
		if arr[i] > S[i] {
			S[i] = arr[i]
		}
	}
	S = S[1:]
	S = append(S, -1)
	return S
}
