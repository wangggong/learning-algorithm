/*
 * @lc app=leetcode.cn id=mini-parser lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [385] 迷你语法分析器
 *
 * https://leetcode-cn.com/problems/mini-parser/description/
 *
 * algorithms
 * Medium (42.28%)
 * Total Accepted:    14.9K
 * Total Submissions: 29.5K
 * Testcase Example:  '"324"'
 *
 * 给定一个字符串 s 表示一个整数嵌套列表，实现一个解析它的语法分析器并返回解析的结果 NestedInteger 。
 *
 * 列表中的每个元素只可能是整数或整数嵌套列表
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "324",
 * 输出：324
 * 解释：你应该返回一个 NestedInteger 对象，其中只包含整数值 324。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "[123,[456,[789]]]",
 * 输出：[123,[456,[789]]]
 * 解释：返回一个 NestedInteger 对象包含一个有两个元素的嵌套列表：
 * 1. 一个 integer 包含值 123
 * 2. 一个包含两个元素的嵌套列表：
 * ⁠   i.  一个 integer 包含值 456
 * ⁠   ii. 一个包含一个元素的嵌套列表
 * ⁠        a. 一个 integer 包含值 789
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 5 * 10^4
 * s 由数字、方括号 "[]"、负号 '-' 、逗号 ','组成
 * 用例保证 s 是可解析的 NestedInteger
 * 输入中的所有值的范围是 [-10^6, 10^6]
 *
 *
 */
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

func NewNestedInteger() *NestedInteger { return &NestedInteger{} }

/*
 * func deserialize(s string) *NestedInteger {
 * 	arr := []byte(s)
 * 	return _deserialize(arr, 0, len(arr))
 * }
 *
 * func _deserialize(arr []byte, l, r int) *NestedInteger {
 * 	if l >= r {
 * 		return nil
 * 	}
 * 	neg := false
 * 	switch arr[l] {
 * 	case '-':
 * 		neg = true
 * 		l++
 * 		fallthrough
 * 	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
 * 		ans := 0
 * 		i := l
 * 		for ; i < r && '0' <= arr[i] && arr[i] <= '9'; i++ {
 * 			ans = ans*10 + int(arr[i]-'0')
 * 		}
 * 		if neg {
 * 			ans *= -1
 * 		}
 * 		ni := NewNestedInteger()
 * 		ni.SetInteger(ans)
 * 		return ni
 * 	case '[':
 * 		// assert arr[r-1] == ']'
 * 		ni := NewNestedInteger()
 * 		for i := l + 1; i < r; i++ {
 * 			j := i + 1
 * 			if arr[i] == '-' || ('0' <= arr[i] && arr[i] <= '9') {
 * 				for ; j < r && '0' <= arr[j] && arr[j] <= '9'; j++ {
 * 				}
 * 			} else {
 * 				k := 1
 * 				for ; j < r && k > 0; j++ {
 * 					if arr[j] == '[' {
 * 						k++
 * 					} else if arr[j] == ']' {
 * 						k--
 * 					}
 * 				}
 * 			}
 * 			if elem := _deserialize(arr, i, j); elem != nil {
 * 				ni.Add(*elem)
 * 			}
 * 			i = j
 * 		}
 * 		return ni
 * 	default:
 * 		// unreachable
 * 	}
 * 	// unreachable
 * 	return nil
 * }
 */

func deserialize(s string) *NestedInteger {
	arr := []byte(s)
	ni, _ := _deserialize(arr, 0)
	return ni
}

func _deserialize(arr []byte, k int) (*NestedInteger, int) {
	n := len(arr)
	if k >= n {
		return nil, k
	}
	ni := NewNestedInteger()
	neg := false
	switch arr[k] {
	case '-':
		// 符号正常处理. 我们认定 `arr` 合法, 那负号后一定是数字, 接着搞就成.
		neg = true
		k++
		fallthrough
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		// 数字就直接算呗. 为了处理方便, 仅处理数字. 后面接什么都不管, 那应该是外层调用管的事.
		ans := 0
		for ; k < n && '0' <= arr[k] && arr[k] <= '9'; k++ {
			ans = ans*10 + int(arr[k]-'0')
		}
		if neg {
			ans *= -1
		}
		ni.SetInteger(ans)
	case '[':
		// 找到匹配的括号.
		m := 1
		k++
		right := k
		for ; right < n && m > 0; right++ {
			if arr[right] == '[' {
				m++
			} else if arr[right] == ']' {
				m--
			}
		}
		// `arr[k:right+1]` 为当前的 token. 这时需要遍历其中每个元素, 往 NestedInteger 中添加.
		for k < right {
			// 注意, 如果出现了 "[[]]" 的情况, 在内层括号解析完成后返回的位置将是 `3` 号位置, 是个 "]". 这时得给它跳过去 :facepalm:
			if arr[k] == ']' {
				k++
				continue
			}
			elem, next := _deserialize(arr, k)
			if elem != nil {
				ni.Add(*elem)
			}
			k = next
		}
	}
	// 返回的 `k+1` 始终代表当前 token 结束位置的 **下个** 位置, 就能跳过 "," 了.
	return ni, k + 1
}
