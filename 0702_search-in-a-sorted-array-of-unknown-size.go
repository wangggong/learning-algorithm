/*
 * Medium
 *
 * 这是一个交互问题。
 *
 * 您有一个升序整数数组，其长度未知。您没有访问数组的权限，但是可以使用 ArrayReader 接口访问它。你可以调用 ArrayReader.get(i):
 *
 * 返回数组第ith个索引(0-indexed)处的值(即secret[i])，或者
 *
 * 如果 i  超出了数组的边界，则返回 231 - 1
 *
 * 你也会得到一个整数 target。
 *
 * 如果存在secret[k] == target，请返回索引 k 的值；否则返回 -1
 *
 * 你必须写一个时间复杂度为 O(log n) 的算法。
 *
 *
 *
 * 示例 1：
 *
 * 输入: secret = [-1,0,3,5,9,12], target = 9
 * 输出: 4
 * 解释: 9 存在在 nums 中，下标为 4
 * 示例 2：
 *
 * 输入: secret = [-1,0,3,5,9,12], target = 2
 * 输出: -1
 * 解释: 2 不在数组中所以返回 -1
 *
 *
 * 提示：
 *
 * 1 <= secret.length <= 1e4
 * -1e4 <= secret[i], target <= 1e4
 * secret 严格递增
 * 通过次数5,383
 * 提交次数7,245
 */

/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 *
 * func (this *ArrayReader) get(index int) int {}
 */

const MAXN = int(1e4)

/*
 * func search(reader ArrayReader, target int) int {
 * 	// get size;
 * 	// 1 1 1 1 1 1 0 0
 * 	p, q := 0, MAXN+1
 * 	for p < q {
 * 		mid := (p + q + 1) >> 1
 * 		if v := reader.get(mid); v > MAXN {
 * 			q = mid - 1
 * 		} else {
 * 			if v == target {
 * 				return mid
 * 			}
 * 			p = mid
 * 		}
 * 	}
 * 	lastIndex := p
 * 	p, q = 0, lastIndex
 * 	for p <= q {
 * 		mid := (p + q) >> 1
 * 		if v := reader.get(mid); v == target {
 * 			return mid
 * 		} else if v > target {
 * 			q = mid - 1
 * 		} else {
 * 			p = mid + 1
 * 		}
 * 	}
 * 	return -1
 * }
 */

func search(reader ArrayReader, target int) int {
	p, q := 0, MAXN-1
	for p <= q {
		mid := (p + q) >> 1
		if v := reader.get(mid); v == target {
			return mid
		} else if v > target {
			q = mid - 1
		} else {
			p = mid + 1
		}
	}
	return -1
}
