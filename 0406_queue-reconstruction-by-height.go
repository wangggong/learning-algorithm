/*
 * @lc app=leetcode.cn id=queue-reconstruction-by-height lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [406] 根据身高重建队列
 *
 * https://leetcode-cn.com/problems/queue-reconstruction-by-height/description/
 *
 * algorithms
 * Medium (74.51%)
 * Total Accepted:    149K
 * Total Submissions: 200K
 * Testcase Example:  '[[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]'
 *
 * 假设有打乱顺序的一群人站成一个队列，数组 people 表示队列中一些人的属性（不一定按顺序）。每个 people[i] = [hi, ki] 表示第
 * i 个人的身高为 hi ，前面 正好 有 ki 个身高大于或等于 hi 的人。
 *
 * 请你重新构造并返回输入数组 people 所表示的队列。返回的队列应该格式化为数组 queue ，其中 queue[j] = [hj, kj]
 * 是队列中第 j 个人的属性（queue[0] 是排在队列前面的人）。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：people = [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
 * 输出：[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
 * 解释：
 * 编号为 0 的人身高为 5 ，没有身高更高或者相同的人排在他前面。
 * 编号为 1 的人身高为 7 ，没有身高更高或者相同的人排在他前面。
 * 编号为 2 的人身高为 5 ，有 2 个身高更高或者相同的人排在他前面，即编号为 0 和 1 的人。
 * 编号为 3 的人身高为 6 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
 * 编号为 4 的人身高为 4 ，有 4 个身高更高或者相同的人排在他前面，即编号为 0、1、2、3 的人。
 * 编号为 5 的人身高为 7 ，有 1 个身高更高或者相同的人排在他前面，即编号为 1 的人。
 * 因此 [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]] 是重新构造后的队列。
 *
 *
 * 示例 2：
 *
 *
 * 输入：people = [[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]
 * 输出：[[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= people.length <= 2000
 * 0 <= hi <= 1e6
 * 0 <= ki < people.length
 * 题目数据确保队列可以被重建
 *
 *
 */

import "sort"

/*
 * // 这样是不对的, 不能按层来.
 * func reconstructQueue(P [][]int) [][]int {
 * 	n := len(P)
 * 	var Q []int
 * 	visit := make([]bool, n)
 * 	count := make([]int, n)
 * 	for i, p := range P {
 * 		count[i] = p[1]
 * 		if count[i] == 0 {
 * 			Q = append(Q, i)
 * 			visit[i] = true
 * 		}
 * 	}
 * 	var ans [][]int
 * 	for len(Q) > 0 {
 * 		var cur [][]int
 * 		for size := len(Q); size > 0; size-- {
 * 			q := Q[0]
 * 			Q = Q[1:]
 * 			cur = append(cur, P[q])
 * 			for i := 0; i < n; i++ {
 * 				if visit[i] {
 * 					continue
 * 				}
 * 				if P[i][0] <= P[q][0] {
 * 					count[i]--
 * 				}
 * 				if count[i] == 0 {
 * 					Q = append(Q, i)
 * 					visit[i] = true
 * 				}
 * 			}
 * 		}
 * 		sort.Slice(cur, func(p, q int) bool { return cur[p][0] < cur[q][0] })
 * 		ans = append(ans, cur...)
 * 	}
 * 	return ans
 * }
 */

type Node struct {
	Val []int
	Next *Node
}

func reconstructQueue(P [][]int) [][]int {
	n := len(P)
	// NOTE: 排序策略:
	// value 从大到小 -> k 从小到大.
	// 排序后就可以直接依次插入了.
	sort.Slice(P, func(p, q int) bool {
		if P[p][0] != P[q][0] {
			return P[p][0] > P[q][0]
		}
		return P[p][1] < P[q][1]
	})
	dummy := &Node{}
	for i := 0; i < n; i++ {
		p := P[i]
		prev, curr := (*Node)(nil), dummy
		for j := p[1]; j >= 0 && curr != nil; j-- {
			prev, curr = curr, curr.Next
		}
		prev.Next = &Node{Val: p, Next: curr}
	}
	var ans [][]int
	for curr := dummy.Next; curr != nil; curr = curr.Next {
		ans = append(ans, curr.Val)
	}
	return ans
}

/*
 * // 数组没有 insert 也太蛋疼了...
 * func reconstructQueue(P [][]int) [][]int {
 * 	n := len(P)
 * 	sort.Slice(P, func(p, q int) bool {
 * 		if P[p][0] != P[q][0] {
 * 			return P[p][0] > P[q][0]
 * 		}
 * 		return P[p][1] < P[q][1]
 * 	})
 * 	var ans [][]int
 * 	for i := 0; i < n; i++ {
 * 		p := P[i]
 * 		if p[1] == 0 {
 * 			ans = append([][]int{p}, ans...)
 * 			continue
 * 		}
 * 		var tmp [][]int
 * 		for i, a := range ans {
 * 			if i == p[1] {
 * 				tmp = append(tmp, p)
 * 			}
 * 			tmp = append(tmp, a)
 * 		}
 * 		if len(tmp) == p[1] {
 * 			tmp = append(tmp, p)
 * 		}
 * 		ans = tmp
 * 	}
 * 	return ans
 * }
 */

/*
 * // 二刷还不会, 欠刷.
 * func reconstructQueue(P [][]int) [][]int {
 *     n := len(P)
 *     sort.Slice(P, func(p, q int) bool {
 *         if P[p][0] != P[q][0] {
 *             return P[p][0] > P[q][0]
 *         }
 *         return P[p][1] < P[q][1]
 *     })
 *     var ans [][]int
 *     for i := 0; i < n; i++ {
 *         if P[i][1] > len(ans) {
 *             ans = append(ans, P[i])
 *             continue
 *         }
 *         var tmp [][]int
 *         ind := P[i][1]
 *         tmp = append(tmp, ans[:ind]...)
 *         tmp = append(tmp, P[i])
 *         tmp = append(tmp, ans[ind:]...)
 *         ans = tmp
 *     }
 *     return ans
 * }
 */
