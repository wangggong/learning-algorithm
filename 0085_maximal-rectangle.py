class Solution:
    def maximalRectangle(self, M: List[List[str]]) -> int:
        if len(M) == 0:
            return 0

        n, m = len(M), len(M[0])
        H = [[0] * m for _ in range(n)]
        for i in range(m):
            H[0][i] = 1 if M[0][i] == "1" else 0
        for i in range(1, n):
            for j in range(m):
                H[i][j] = H[i - 1][j] + 1 if M[i][j] == "1" else 0

        # print(H)
        ans = 0
        for i in range(n):
            SL, SR = [], []
            up, down = [], []
            for j in range(m):
                while len(SL) > 0 and H[i][SL[-1]] >= H[i][j]:
                    SL = SL[:-1]
                up += [-1 if len(SL) == 0 else SL[-1]]
                SL += [j]
            for j in range(m - 1, -1, -1):
                while len(SR) > 0 and H[i][SR[-1]] >= H[i][j]:
                    SR = SR[:-1]
                down = [m if len(SR) == 0 else SR[-1]] + down
                SR += [j]
            for k, (u, d) in enumerate(zip(up, down)):
                ans = max(ans, H[i][k] * (d - u - 1))

            # print("up =", up)
            # print("down =", down)
        return ans
