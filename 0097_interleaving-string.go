const MAXN = 1000 + 10

func isInterleave(s1 string, s2 string, s3 string) bool {
	var dp [MAXN][MAXN]bool

	b1, b2, b3 := []byte(s1), []byte(s2), []byte(s3)
	m, n := len(b1), len(b2)

	if len(b3) != m+n {
		return false
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
				continue
			}

			p, q, r := byte(0), byte(0), byte(0)
			if i > 0 {
				p = b1[i-1]
			}

			if j > 0 {
				q = b2[j-1]
			}

			if i+j > 0 {
				r = b3[i+j-1]
			}

			if p == r {
				dp[i][j] = dp[i][j] || dp[i-1][j]
			}

			if q == r {
				dp[i][j] = dp[i][j] || dp[i][j-1]
			}
		}
	}

	return dp[m][n]
}
