/*
 * @lc app=leetcode.cn id=binary-watch lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [401] 二进制手表
 *
 * https://leetcode-cn.com/problems/binary-watch/description/
 *
 * algorithms
 * Easy (61.59%)
 * Total Accepted:    57.7K
 * Total Submissions: 93.7K
 * Testcase Example:  '1'
 *
 * 二进制手表顶部有 4 个 LED 代表 小时（0-11），底部的 6 个 LED 代表 分钟（0-59）。每个 LED 代表一个 0 或
 * 1，最低位在右侧。
 * 
 * 
 * 例如，下面的二进制手表读取 "3:25" 。
 * 
 * 
 * 
 * 
 * （图源：WikiMedia - Binary clock samui moon.jpg ，许可协议：Attribution-ShareAlike 3.0
 * Unported (CC BY-SA 3.0) ）
 * 
 * 给你一个整数 turnedOn ，表示当前亮着的 LED 的数量，返回二进制手表可以表示的所有可能时间。你可以 按任意顺序 返回答案。
 * 
 * 小时不会以零开头：
 * 
 * 
 * 例如，"01:00" 是无效的时间，正确的写法应该是 "1:00" 。
 * 
 * 
 * 分钟必须由两位数组成，可能会以零开头：
 * 
 * 
 * 例如，"10:2" 是无效的时间，正确的写法应该是 "10:02" 。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：turnedOn = 1
 * 输出：["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：turnedOn = 9
 * 输出：[]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= turnedOn <= 10
 * 
 * 
 */
func readBinaryWatch(turnedOn int) []string {
	var ans []string
	for i := 0; i <= turnedOn; i++ {
		for h := 0; h < 12; h++ {
			if bits(h) != i {
				continue
			}
			for m := 0; m < 60; m++ {
				if bits(m) != turnedOn-i {
					continue
				}
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return ans
}

func bits(n int) int {
	ans := 0
	for n != 0 {
		ans++
		n = n & (n-1)
	}
	return ans
}
