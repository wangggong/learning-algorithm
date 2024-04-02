/*
 * Hard
 *
 * 给定一个数组，包含从 1 到 N 所有的整数，但其中缺了两个数字。你能在 O(N) 时间内只用 O(1) 的空间找到它们吗？
 *
 * 以任意顺序返回这两个数字均可。
 *
 * 示例 1:
 *
 * 输入: [1]
 * 输出: [2,3]
 * 示例 2:
 *
 * 输入: [2,3]
 * 输出: [1,4]
 * 提示：
 *
 * nums.length <= 30000
 */

func missingTwo(nums []int) []int {
	n := len(nums)
	tot := (n + 2) * (n + 3) / 2
	for _, n := range nums {
		tot -= n
	}
	p, q := 0, n
	for p < q {
		k := partition(nums, p, q)
		if k == nums[k]-1 {
			p = k + 1
		} else {
			q = k
		}
	}
	return []int{p + 1, tot - (p + 1)}
}

func partition(arr []int, l, r int) int {
	start := l
	for i := l + 1; i < r; i++ {
		if arr[i] < arr[l] {
			start++
			arr[i], arr[start] = arr[start], arr[i]
		}
	}
	arr[l], arr[start] = arr[start], arr[l]
	return start
}
