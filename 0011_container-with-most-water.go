func maxArea(H []int) int {
	n := len(H)
	ans := 0

	for p, q := 0, n-1; p < q; {
		cur := 0
		if H[p] < H[q] {
			cur = H[p] * (q - p)
			p++
		} else {
			cur = H[q] * (q - p)
			q--
		}

		if cur > ans {
			ans = cur
		}
	}

	return ans
}
