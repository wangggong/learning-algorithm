/*
 * @lc app=leetcode.cn id=duplicate-zeros lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1089] 复写零
 *
 * https://leetcode.cn/problems/duplicate-zeros/description/
 *
 * algorithms
 * Easy (58.15%)
 * Total Accepted:    34.1K
 * Total Submissions: 56.4K
 * Testcase Example:  '[1,0,2,3,0,4,5,0]'
 *
 * 给你一个长度固定的整数数组 arr，请你将该数组中出现的每个零都复写一遍，并将其余的元素向右平移。
 *
 * 注意：请不要在超过该数组长度的位置写入元素。
 *
 * 要求：请对输入的数组 就地 进行上述修改，不要从函数返回任何东西。
 *
 *
 *
 * 示例 1：
 *
 * 输入：[1,0,2,3,0,4,5,0]
 * 输出：null
 * 解释：调用函数后，输入的数组将被修改为：[1,0,0,2,3,0,0,4]
 *
 *
 * 示例 2：
 *
 * 输入：[1,2,3]
 * 输出：null
 * 解释：调用函数后，输入的数组将被修改为：[1,2,3]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= arr.length <= 10000
 * 0 <= arr[i] <= 9
 *
 *
 */
func duplicateZeros(arr []int) {
	n, k, top := len(arr), 0, 0
	push := func() {
		top++
		return
	}
	popAndCpy := func() {
        // 这里注意一定是先 pop 再 copy, 顺序反了会多 copy 0 的...
		top--
		if 0 <= top && top < n {
			arr[top] = arr[k]
		}
		return
	}
	for ; k < n && top < n; k++ {
		push()
		if arr[k] == 0 {
			push()
		}
	}
	for k > 0 {
		k--
		popAndCpy()
		if arr[k] == 0 {
			popAndCpy()
		}
	}
	return
}
