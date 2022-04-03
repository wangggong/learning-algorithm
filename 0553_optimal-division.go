/*
 * @lc app=leetcode.cn id=optimal-division lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [553] 最优除法
 *
 * https://leetcode-cn.com/problems/optimal-division/description/
 *
 * algorithms
 * Medium (61.09%)
 * Total Accepted:    8K
 * Total Submissions: 12.9K
 * Testcase Example:  '[1000,100,10,2]'
 *
 * 给定一组正整数，相邻的整数之间将会进行浮点除法操作。例如， [2,3,4] -> 2 / 3 / 4 。
 *
 *
 * 但是，你可以在任意位置添加任意数目的括号，来改变算数的优先级。你需要找出怎么添加括号，才能得到最大的结果，并且返回相应的字符串格式的表达式。你的表达式不应该含有冗余的括号。
 *
 * 示例：
 *
 *
 * 输入: [1000,100,10,2]
 * 输出: "1000/(100/10/2)"
 * 解释:
 * 1000/(100/10/2) = 1000/((100/10)/2) = 200
 * 但是，以下加粗的括号 "1000/((100/10)/2)" 是冗余的，
 * 因为他们并不影响操作的优先级，所以你需要返回 "1000/(100/10/2)"。
 *
 * 其他用例:
 * 1000/(100/10)/2 = 50
 * 1000/(100/(10/2)) = 50
 * 1000/100/10/2 = 0.5
 * 1000/100/(10/2) = 2
 *
 *
 * 说明:
 *
 *
 * 输入数组的长度在 [1, 10] 之间。
 * 数组中每个元素的大小都在 [2, 1000] 之间。
 * 每个测试用例只有一个最优除法解。
 *
 *
 */

import (
	"strconv"
	"strings"
)

const MAXN = 10 + 10

/*
 * func optimalDivision(nums []int) string {
 * 	n := len(nums)
 * 	// assert n > 0 && n <= 10;
 * 	if n == 1 {
 * 		return strconv.Itoa(nums[0])
 * 	}
 *
 * 	// 1000 100 10 2
 * 	index := -1
 * 	maxVal := float64(0)
 * 	for i := 1; i < n; i++ {
 * 		num := float64(nums[0])
 * 		for j := 1; j < i; j++ {
 * 			num /= float64(nums[j])
 * 		}
 * 		den := float64(nums[i])
 * 		for j := i; j < n; j++ {
 * 			den /= float64(nums[j])
 * 		}
 * 		if v := num / den; v > maxVal {
 * 			index = i
 * 			maxVal = v
 * 		}
 * 	}
 * 	var strs []string
 * 	for _, v := range nums {
 * 		strs = append(strs, strconv.Itoa(v))
 * 	}
 * 	if len(strs[index:]) == 1 {
 * 		return fmt.Sprintf("%s/%s", strings.Join(strs[:index], "/"), strings.Join(strs[index:], "/"))
 * 	}
 * 	return fmt.Sprintf("%s/(%s)", strings.Join(strs[:index], "/"), strings.Join(strs[index:], "/"))
 * }
 */

func optimalDivision(nums []int) string {
	switch len(nums) {
	case 0:
		return ""
	case 1:
		return strconv.Itoa(nums[0])
	case 2:
		return strconv.Itoa(nums[0]) + "/" + strconv.Itoa(nums[1])
	default:
		var arr []string
		for _, n := range nums[1:] {
			arr = append(arr, strconv.Itoa(n))
		}
		return strconv.Itoa(nums[0]) + "/(" + strings.Join(arr, "/") + ")"
	}
}
