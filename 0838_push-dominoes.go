/*
 * @lc app=leetcode.cn id=push-dominoes lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [838] 推多米诺
 *
 * https://leetcode-cn.com/problems/push-dominoes/description/
 *
 * algorithms
 * Medium (50.42%)
 * Total Accepted:    12.9K
 * Total Submissions: 23.7K
 * Testcase Example:  '"RR.L"'
 *
 * n 张多米诺骨牌排成一行，将每张多米诺骨牌垂直竖立。在开始时，同时把一些多米诺骨牌向左或向右推。
 *
 * 每过一秒，倒向左边的多米诺骨牌会推动其左侧相邻的多米诺骨牌。同样地，倒向右边的多米诺骨牌也会推动竖立在其右侧的相邻多米诺骨牌。
 *
 * 如果一张垂直竖立的多米诺骨牌的两侧同时有多米诺骨牌倒下时，由于受力平衡， 该骨牌仍然保持不变。
 *
 * 就这个问题而言，我们会认为一张正在倒下的多米诺骨牌不会对其它正在倒下或已经倒下的多米诺骨牌施加额外的力。
 *
 * 给你一个字符串 dominoes 表示这一行多米诺骨牌的初始状态，其中：
 *
 *
 * dominoes[i] = 'L'，表示第 i 张多米诺骨牌被推向左侧，
 * dominoes[i] = 'R'，表示第 i 张多米诺骨牌被推向右侧，
 * dominoes[i] = '.'，表示没有推动第 i 张多米诺骨牌。
 *
 *
 * 返回表示最终状态的字符串。
 *
 *
 * 示例 1：
 *
 *
 * 输入：dominoes = "RR.L"
 * 输出："RR.L"
 * 解释：第一张多米诺骨牌没有给第二张施加额外的力。
 *
 *
 * 示例 2：
 *
 *
 * 输入：dominoes = ".L.R...LR..L.."
 * 输出："LL.RR.LLRRLL.."
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == dominoes.length
 * 1 <= n <= 10^5
 * dominoes[i] 为 'L'、'R' 或 '.'
 *
 *
 */

// 参考: https://leetcode-cn.com/problems/push-dominoes/solution/gong-shui-san-xie-yi-ti-shuang-jie-bfs-y-z52w/
const MAXN = 1e5 + 10

/*
 * // 双指针: 先计算两侧受力, 再分情况讨论. 愣写.
 *
 * var left, right [MAXN]byte
 *
 * func pushDominoes(dominoes string) string {
 * 	bytes := []byte(dominoes)
 * 	n := len(bytes)
 * 	for i := 0; i < MAXN; i++ {
 * 		left[i] = '.'
 * 		right[i] = '.'
 * 	}
 * 	first, last := -1, -1
 * 	for i := 0; i < n; i++ {
 * 		if bytes[i] != '.' {
 * 			first = i
 * 		}
 * 		if first >= 0 {
 * 			left[i] = bytes[first]
 * 		}
 * 	}
 * 	for i := n - 1; i >= 0; i-- {
 * 		if bytes[i] != '.' {
 * 			last = i
 * 		}
 * 		if last >= 0 {
 * 			right[i] = bytes[last]
 * 		}
 * 	}
 * 	// fmt.Printf("%s\n%s\n", left, right)
 * 	for i := 0; i < n; i++ {
 * 		if bytes[i] != '.' {
 * 			continue
 * 		}
 * 		j := i + 1
 * 		for ; j < n && bytes[j] == '.'; j++ {
 * 		}
 * 		j--
 * 		li, rj := left[i], right[j]
 * 		if li == '.' && rj == '.' {
 * 			// do nothing
 * 		} else if li == '.' {
 * 			if rj == 'L' {
 * 				update(bytes, i, j, rj, rj)
 * 			}
 * 		} else if rj == '.' {
 * 			if li == 'R' {
 * 				update(bytes, i, j, li, li)
 * 			}
 * 		} else if !(li == 'L' && rj == 'R') {
 * 			update(bytes, i, j, li, rj)
 * 		}
 * 	}
 * 	return string(bytes)
 * }
 *
 * func update(arr []byte, p, q int, left, right byte) {
 * 	for ; p <= q; p, q = p+1, q-1 {
 * 		if left != right && p == q {
 * 			continue
 * 		}
 * 		arr[p], arr[q] = left, right
 * 	}
 * }
 */

// 多源 BFS.

var visit [MAXN]int

type Info struct {
	Direction byte
	Pos, Time int
}

func pushDominoes(dominoes string) string {
	bytes := []byte(dominoes)
	n := len(bytes)
	for i := range visit {
		visit[i] = 0
	}
	var Q []Info
	for i, b := range bytes {
		if b == '.' {
			continue
		}
		Q = append(Q, Info{b, i, 1})
		visit[i] = 1
	}
	for len(Q) > 0 {
		q := Q[0]
		Q = Q[1:]
		var next int
		switch q.Direction {
		case 'L':
			next = q.Pos - 1
		case 'R':
			next = q.Pos + 1
		default:
			// unreachable
		}
		if next < 0 || next >= n || bytes[q.Pos] == '.' {
			continue
		}
		if visit[next] == 0 {
			bytes[next] = q.Direction
			Q = append(Q, Info{q.Direction, next, q.Time + 1})
			visit[next] = q.Time + 1
			continue
		}
		if visit[next] == q.Time+1 {
			bytes[next] = '.'
			continue
		}
	}
	return string(bytes)
}
