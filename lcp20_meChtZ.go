import "container/heap"

/*
 * @lc app=leetcode.cn id=meChtZ lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [LCP 20] 快速公交
 *
 * https://leetcode-cn.com/problems/meChtZ/description/
 *
 * algorithms
 * Hard (32.37%)
 * Total Accepted:    2K
 * Total Submissions: 6.1K
 * Testcase Example:  '31\n5\n3\n[6]\n[10]'
 *
 * 小扣打算去秋日市集，由于游客较多，小扣的移动速度受到了人流影响：
 * - 小扣从 `x` 号站点移动至 `x + 1` 号站点需要花费的时间为 `inc`；
 * - 小扣从 `x` 号站点移动至 `x - 1` 号站点需要花费的时间为 `dec`。
 *
 * 现有 `m` 辆公交车，编号为 `0` 到 `m-1`。小扣也可以通过搭乘编号为 `i` 的公交车，从 `x` 号站点移动至 `jump[i]*x`
 * 号站点，耗时仅为 `cost[i]`。小扣可以搭乘任意编号的公交车且搭乘公交次数不限。
 *
 * 假定小扣起始站点记作 `0`，秋日市集站点记作 `target`，请返回小扣抵达秋日市集最少需要花费多少时间。由于数字较大，最终答案需要对
 * 1000000007 (1e9 + 7) 取模。
 *
 * 注意：小扣可在移动过程中到达编号大于 `target` 的站点。
 *
 * **示例 1：**
 * >输入：`target = 31, inc =  5, dec = 3, jump = [6], cost = [10]`
 * >
 * >输出：`33`
 * >
 * >解释：
 * >小扣步行至 1 号站点，花费时间为 5；
 * >小扣从 1 号站台搭乘 0 号公交至 6 * 1 = 6 站台，花费时间为 10；
 * >小扣从 6 号站台步行至 5 号站台，花费时间为 3；
 * >小扣从 5 号站台搭乘 0 号公交至 6 * 5 = 30 站台，花费时间为 10；
 * >小扣从 30 号站台步行至 31 号站台，花费时间为 5；
 * >最终小扣花费总时间为 33。
 *
 *
 * **示例 2：**
 * >输入：`target = 612, inc =  4, dec = 5, jump = [3,6,8,11,5,10,4], cost =
 * [4,7,6,3,7,6,4]`
 * >
 * >输出：`26`
 * >
 * >解释：
 * >小扣步行至 1 号站点，花费时间为 4；
 * >小扣从 1 号站台搭乘 0 号公交至 3 * 1 = 3 站台，花费时间为 4；
 * >小扣从 3 号站台搭乘 3 号公交至 11 * 3 = 33 站台，花费时间为 3；
 * >小扣从 33 号站台步行至 34 站台，花费时间为 4；
 * >小扣从 34 号站台搭乘 0 号公交至 3 * 34 = 102 站台，花费时间为 4；
 * >小扣从 102 号站台搭乘 1 号公交至 6 * 102 = 612 站台，花费时间为 7；
 * >最终小扣花费总时间为 26。
 *
 *
 * **提示：**
 * 1 <= target <= 10^9
 * 1 <= jump.length, cost.length <= 10
 * 2 <= jump[i] <= 10^6
 * 1 <= inc, dec, cost[i] <= 10^6
 */

const MOD int = 1e9 + 7

/*
 * var memo map[int]int
 * var n int
 *
 * // 参考 https://leetcode-cn.com/problems/meChtZ/solution/go-ji-yi-hua-dfsjie-fa-by-ling-ji-zhi-chu/
 * func busRapidTransit(target int, inc int, dec int, jump []int, cost []int) int {
 * 	memo = make(map[int]int)
 * 	n = len(jump)
 * 	return dfs(target, inc, dec, jump, cost) % MOD
 * }
 */

const MOD int = 1e9 + 7

type Heap struct {
	Arr [][2]int
}

func (h Heap) Len() int            { return len(h.Arr) }
func (h Heap) Less(p, q int) bool  { return h.Arr[p][1] < h.Arr[q][1] }
func (h Heap) Swap(p, q int)       { h.Arr[p], h.Arr[q] = h.Arr[q], h.Arr[p] }
func (h *Heap) Push(v interface{}) { h.Arr = append(h.Arr, v.([2]int)) }
func (h *Heap) Pop() interface{}   { l := len(h.Arr); v := h.Arr[l-1]; h.Arr = h.Arr[:l-1]; return v }

// 优先队列+BFS, 参考: https://leetcode-cn.com/problems/meChtZ/solution/suan-fa-xiao-ai-li-kou-bei-li-jie-zhen-t-tg1z/
//
// 注意:
// 1. 倒着找;
// 2. 不需要 visit, 以下操作保证不会重复访问;
// 3. 剪枝: 已经比当前时间长了就不继续了, 但也不 break;
func busRapidTransit(target int, inc int, dec int, jump []int, cost []int) int {
	h := &Heap{}
	n := len(jump)
	heap.Push(h, [2]int{target, 0})
	ans := target * inc
	for h.Len() > 0 {
		p := heap.Pop(h).([2]int)
		if p[1] > ans {
			continue
		}
		ans = min(ans, p[0]*inc+p[1])
		for i := 0; i < n; i++ {
			div, mod := p[0]/jump[i], p[0]%jump[i]
			if mod == 0 {
				heap.Push(h, [2]int{div, p[1] + cost[i]})
			} else {
				heap.Push(h, [2]int{div, p[1] + cost[i] + mod*inc})
				heap.Push(h, [2]int{div + 1, p[1] + cost[i] + (jump[i]-mod)*dec})
			}
		}
	}
	return ans % MOD
}

func dfs(target, inc, dec int, jump, cost []int) (ans int) {
	if v, ok := memo[target]; ok {
		return v
	}
	defer func() { memo[target] = ans }()
	if target == 1 {
		ans = inc
		return
	}
	if target == 0 {
		ans = 0
		return
	}
	ans = target * inc
	for i := 0; i < n; i++ {
		div, mod := target/jump[i], target%jump[i]
		if mod == 0 {
			ans = min(ans, dfs(div, inc, dec, jump, cost)+cost[i])
			continue
		}
		ans = min(ans, dfs(div, inc, dec, jump, cost)+cost[i]+inc*mod)
		ans = min(ans, dfs(div+1, inc, dec, jump, cost)+cost[i]+dec*(jump[i]-mod))
	}
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
