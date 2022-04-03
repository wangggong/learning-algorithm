/*
 * @lc app=leetcode.cn id=russian-doll-envelopes lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [354] 俄罗斯套娃信封问题
 *
 * https://leetcode-cn.com/problems/russian-doll-envelopes/description/
 *
 * algorithms
 * Hard (45.59%)
 * Total Accepted:    78K
 * Total Submissions: 171.6K
 * Testcase Example:  '[[5,4],[6,4],[6,7],[2,3]]'
 *
 * 给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
 *
 * 当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
 *
 * 请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
 *
 * 注意：不允许旋转信封。
 *
 *
 * 示例 1：
 *
 *
 * 输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
 * 输出：3
 * 解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
 *
 * 示例 2：
 *
 *
 * 输入：envelopes = [[1,1],[1,1],[1,1]]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= envelopes.length <= 10^5
 * envelopes[i].length == 2
 * 1 <= wi, hi <= 10^5
 *
 *
 */
import "sort"

func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	}
	sort.Slice(envelopes, func(p, q int) bool {
		if envelopes[p][1] != envelopes[q][1] {
			return envelopes[p][1] < envelopes[q][1]
		}
		return envelopes[p][0] > envelopes[q][0]
	})
	var end []int
	end = append(end, envelopes[0][0])
	ans := 1
	for i := 1; i < n; i++ {
		width := envelopes[i][0]
		if width > end[len(end)-1] {
			end = append(end, envelopes[i][0])
			ans = max(ans, len(end))
			continue
		}
		p, q := 0, len(end)
		for p < q {
			mid := (p + q) >> 1
			if width > end[mid] {
				p = mid + 1
			} else {
				q = mid
			}
		}
		end[p] = width
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
