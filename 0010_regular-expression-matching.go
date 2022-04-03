const MAXN = 1000 + 5

func match(s, p []byte, i, j int) bool {
	if i == 0 || j == 0 {
		return i == 0 && j == 0
	}

	if p[j-1] == '.' {
		return true
	}

	return s[i-1] == p[j-1]
}

func isMatch(s string, p string) bool {
	bs, bp := []byte(s), []byte(p)
	m, n := len(bs), len(bp)
	var dp [MAXN][MAXN]bool

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			dp[i][j] = false
		}
	}

	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if bp[j-1] == '*' {
				if j > 1 {
					dp[i][j] = dp[i][j] || dp[i][j-2]
				}
				if match(bs, bp, i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else {
				if match(bs, bp, i, j) {
					dp[i][j] = dp[i][j] || dp[i-1][j-1]
				}
			}
		}
	}

	return dp[m][n]
}
