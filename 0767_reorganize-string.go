/*
 * @lc app=leetcode.cn id=reorganize-string lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [767] 重构字符串
 *
 * https://leetcode.cn/problems/reorganize-string/description/
 *
 * algorithms
 * Medium (47.97%)
 * Total Accepted:    44.7K
 * Total Submissions: 93.2K
 * Testcase Example:  '"aab"'
 *
 * 给定一个字符串 s ，检查是否能重新排布其中的字母，使得两相邻的字符不同。
 *
 * 返回 s 的任意可能的重新排列。若不可行，返回空字符串 "" 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "aab"
 * 输出: "aba"
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "aaab"
 * 输出: ""
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= s.length <= 500
 * s 只包含小写字母
 *
 *
 */

import (
	"container/heap"
)

const alphaCnt int = 26

type byteCnt struct {
	Byte byte
	Cnt  int
}

type Heap []byteCnt

func (h Heap) Len() int              { return len(h) }
func (h Heap) Less(p, q int) bool    { return h[p].Cnt > h[q].Cnt }
func (h Heap) Swap(p, q int)         { h[p], h[q] = h[q], h[p]; return }
func (h *Heap) Push(v interface{})   { *h = append(*h, v.(byteCnt)) }
func (h *Heap) Pop() (v interface{}) { l := len(*h); v = (*h)[l-1]; *h = (*h)[:l-1]; return }

func reorganizeString(s string) string {
	arr := []byte(s)
	var cnt [alphaCnt]int
	for _, b := range arr {
		cnt[int(b-'a')]++
	}
	h := &Heap{}
	for b := byte('a'); b <= 'z'; b++ {
		c := cnt[int(b-'a')]
		if c == 0 {
			continue
		}
		heap.Push(h, byteCnt{Byte: b, Cnt: c})
	}
	var ans []byte
	for h.Len() > 1 {
		p := heap.Pop(h).(byteCnt)
		q := heap.Pop(h).(byteCnt)
		ans = append(ans, p.Byte, q.Byte)
		p.Cnt, q.Cnt = p.Cnt-1, q.Cnt-1
		if p.Cnt > 0 {
			heap.Push(h, p)
		}
		if q.Cnt > 0 {
			heap.Push(h, q)
		}
	}
	if h.Len() > 0 {
		p := heap.Pop(h).(byteCnt)
		ans = append(ans, p.Byte)
		p.Cnt = p.Cnt - 1
		if p.Cnt > 0 {
			return ""
		}
	}
	return string(ans)
}
