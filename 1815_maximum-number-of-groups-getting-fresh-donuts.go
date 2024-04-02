/*
 * @lc app=leetcode.cn id=maximum-number-of-groups-getting-fresh-donuts lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1815] 得到新鲜甜甜圈的最多组数
 *
 * https://leetcode.cn/problems/maximum-number-of-groups-getting-fresh-donuts/description/
 *
 * algorithms
 * Hard (32.44%)
 * Total Accepted:    5.2K
 * Total Submissions: 11.8K
 * Testcase Example:  '3\n[1,2,3,4,5,6]'
 *
 * 有一个甜甜圈商店，每批次都烤 batchSize 个甜甜圈。这个店铺有个规则，就是在烤一批新的甜甜圈时，之前 所有
 * 甜甜圈都必须已经全部销售完毕。给你一个整数 batchSize 和一个整数数组 groups ，数组中的每个整数都代表一批前来购买甜甜圈的顾客，其中
 * groups[i] 表示这一批顾客的人数。每一位顾客都恰好只要一个甜甜圈。
 * 
 * 当有一批顾客来到商店时，他们所有人都必须在下一批顾客来之前购买完甜甜圈。如果一批顾客中第一位顾客得到的甜甜圈不是上一组剩下的，那么这一组人都会很开心。
 * 
 * 你可以随意安排每批顾客到来的顺序。请你返回在此前提下，最多 有多少组人会感到开心。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：batchSize = 3, groups = [1,2,3,4,5,6]
 * 输出：4
 * 解释：你可以将这些批次的顾客顺序安排为 [6,2,4,5,1,3] 。那么第 1，2，4，6 组都会感到开心。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：batchSize = 4, groups = [1,3,2,5,2,2,1,6]
 * 输出：4
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 1 
 * 1 
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/maximum-number-of-groups-getting-fresh-donuts/solution/by-endlesscheng-r5ve/
const mask = (1 << 5) - 1

func maxHappyGroups(m int, groups []int) int {
    count := make([]int, m)
    tot := 0
    for _, g := range groups {
        g %= m
        if g == 0 {
            tot++
        } else if count[m - g] > 0 {
            count[m - g]--
            tot++
        } else {
            count[g]++
        }
    }
    val, status := 0, 0
    for i, n := 1, len(count); i < n; i++ {
        if count[i] > 0 {
            val = (val << 5) | i
            status = (status << 5) | count[i]
        }
    }
    // fmt.Println("\t", count, val, status)
    memo := make(map[[2]int]int)
    var dfs func(int, int) int
    dfs = func(left, status int) int {
        key := [2]int{left, status}
        if _, ok := memo[key]; !ok {
            for i := 0; i < 5; i++ {
                d := i*5
                if (status >> d) & mask != 0 {
                    v := (val >> d) & mask
                    res := dfs((left + v) % m, status - (1 << d))
                    if left == 0 {
                        res++
                    }
                    if res > memo[key] {
                        memo[key] = res
                    }
                }
            }
        }
        // fmt.Println(key, memo[key])
        return memo[key]
    }
    return tot + dfs(0, status)
}
