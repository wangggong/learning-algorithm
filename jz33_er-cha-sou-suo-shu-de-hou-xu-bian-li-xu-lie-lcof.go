/*
 * Medium
 *
 * 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。
 *
 *
 *
 * 参考以下这颗二叉搜索树：
 *
 *      5
 *     / \
 *    2   6
 *   / \
 *  1   3
 * 示例 1：
 *
 * 输入: [1,6,3,2,5]
 * 输出: false
 * 示例 2：
 *
 * 输入: [1,3,2,6,5]
 * 输出: true
 *
 *
 * 提示：
 *
 * 数组长度 <= 1000
 * 通过次数156,352
 * 提交次数279,336
 */
func verifyPostorder(postorder []int) bool {
	return isPostorder(postorder, 0, len(postorder))
}

func isPostorder(arr []int, l, r int) bool {
	if l >= r {
		return true
	}
	root := arr[r-1]
	ind := -1
	for i := l; i < r; i++ {
		if arr[i] > root {
			ind = i
			break
		}
	}
	if ind < l {
		return isPostorder(arr, l, r-1)
	}
	for i := ind; i < r; i++ {
		if arr[i] < root {
			return false
		}
	}
	return isPostorder(arr, l, ind) && isPostorder(arr, ind, r-1)
}
