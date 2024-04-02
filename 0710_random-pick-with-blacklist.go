/*
 * @lc app=leetcode.cn id=random-pick-with-blacklist lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [710] 黑名单中的随机数
 *
 * https://leetcode.cn/problems/random-pick-with-blacklist/description/
 *
 * algorithms
 * Hard (38.20%)
 * Total Accepted:    12.8K
 * Total Submissions: 31.8K
 * Testcase Example:  '["Solution","pick","pick","pick","pick","pick","pick","pick"]\n' +
  '[[7,[2,3,5]],[],[],[],[],[],[],[]]'
 *
 * 给定一个整数 n 和一个 无重复 黑名单整数数组 blacklist 。设计一种算法，从 [0, n - 1] 范围内的任意整数中选取一个 未加入
 * 黑名单 blacklist 的整数。任何在上述范围内且不在黑名单 blacklist 中的整数都应该有 同等的可能性 被返回。
 *
 * 优化你的算法，使它最小化调用语言 内置 随机函数的次数。
 *
 * 实现 Solution 类:
 *
 *
 * Solution(int n, int[] blacklist) 初始化整数 n 和被加入黑名单 blacklist 的整数
 * int pick() 返回一个范围为 [0, n - 1] 且不在黑名单 blacklist 中的随机整数
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入
 * ["Solution", "pick", "pick", "pick", "pick", "pick", "pick", "pick"]
 * [[7, [2, 3, 5]], [], [], [], [], [], [], []]
 * 输出
 * [null, 0, 4, 1, 6, 1, 0, 4]
 *
 * 解释
 * Solution solution = new Solution(7, [2, 3, 5]);
 * solution.pick(); // 返回0，任何[0,1,4,6]的整数都可以。注意，对于每一个pick的调用，
 * ⁠                // 0、1、4和6的返回概率必须相等(即概率为1/4)。
 * solution.pick(); // 返回 4
 * solution.pick(); // 返回 1
 * solution.pick(); // 返回 6
 * solution.pick(); // 返回 1
 * solution.pick(); // 返回 0
 * solution.pick(); // 返回 4
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= n <= 10^9
 * 0 <= blacklist.length <= min(10^5, n - 1)
 * 0 <= blacklist[i] < n
 * blacklist 中所有值都 不同
 * pick 最多被调用 2 * 10^4 次
 *
 *
*/

import "math/rand"

/*
 * type Solution struct {
 * 	N            int
 * 	UseBlacklist bool
 * 	Blacklist    map[int]struct{}
 * 	Whitelist    []int
 * }
 *
 * func Constructor(n int, blacklist []int) Solution {
 * 	s := Solution{N: n}
 * 	m := make(map[int]struct{})
 * 	for _, b := range blacklist {
 * 		m[b] = struct{}{}
 * 	}
 * 	if len(blacklist) < n>>1 {
 * 		s.Blacklist = m
 * 		s.UseBlacklist = true
 * 	} else {
 * 		for i := 0; i < n; i++ {
 * 			if _, ok := m[i]; !ok {
 * 				s.Whitelist = append(s.Whitelist, i)
 * 			}
 * 		}
 * 	}
 * 	return s
 * }
 *
 * func (s *Solution) Pick() int {
 * 	if s.UseBlacklist {
 * 		for true {
 * 			k := rand.Intn(s.N)
 * 			if _, ok := s.Blacklist[k]; !ok {
 * 				return k
 * 			}
 * 		}
 * 	} else {
 * 		return s.Whitelist[rand.Intn(len(s.Whitelist))]
 * 	}
 * 	// unreachable
 * 	return -1
 * }
 */

// 黑名单映射, 很漂亮的解法.
type Solution struct {
	N, M  int
	Index map[int]int
}

func Constructor(n int, blacklist []int) Solution {
	s := Solution{
		N:     n,
		M:     len(blacklist),
		Index: make(map[int]int),
	}
	m := make(map[int]struct{})
	for _, b := range blacklist {
		m[b] = struct{}{}
	}
	for i, j := 0, s.N; i < s.M; i++ {
		if blacklist[i] >= s.N-s.M {
			continue
		}
		for j--; j >= 0; j-- {
			if _, ok := m[j]; !ok {
				break
			}
		}
		s.Index[blacklist[i]] = j
	}
	return s
}

func (s *Solution) Pick() int {
	k := rand.Intn(s.N - s.M)
	if v, ok := s.Index[k]; ok {
		return v
	}
	return k
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
