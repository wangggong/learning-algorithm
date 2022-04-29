#
# @lc app=leetcode.cn id=reaching-points lang=python3
#
# NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
#
# [780] 到达终点
#
# https://leetcode-cn.com/problems/reaching-points/description/
#
# algorithms
# Hard (29.57%)
# Total Accepted:    9.6K
# Total Submissions: 27.4K
# Testcase Example:  '1\n1\n3\n5'
#
# 给定四个整数 sx , sy ，tx 和 ty，如果通过一系列的转换可以从起点 (sx, sy) 到达终点 (tx, ty)，则返回 true，否则返回
# false。
# 
# 从点 (x, y) 可以转换到 (x, x+y)  或者 (x+y, y)。
# 
# 
# 
# 示例 1:
# 
# 
# 输入: sx = 1, sy = 1, tx = 3, ty = 5
# 输出: true
# 解释:
# 可以通过以下一系列转换从起点转换到终点：
# (1, 1) -> (1, 2)
# (1, 2) -> (3, 2)
# (3, 2) -> (3, 5)
# 
# 
# 示例 2:
# 
# 
# 输入: sx = 1, sy = 1, tx = 2, ty = 2 
# 输出: false
# 
# 
# 示例 3:
# 
# 
# 输入: sx = 1, sy = 1, tx = 1, ty = 1 
# 输出: true
# 
# 
# 
# 
# 提示:
# 
# 
# 1 <= sx, sy, tx, ty <= 10^9
# 
# 
#
class Solution:
    def reachingPoints(self, sx: int, sy: int, tx: int, ty: int) -> bool:
        # assert sx >= 0 and sy >= 0 and tx >= 0 and ty >= 0
        if sx > tx or sy > ty:
            return False
        # 其中一个维度如果确定了, 那么另一个维度只能从当前值出发.
        # 举例:
        # 一方面, `sx == tx` => `(sx, sy) => (sx, sy+sx) => (sx, sy+2*sx) => ... => (sx, sy+n*sx) = (tx, ty)`
        # 上面的路径唯一. 因为 `sx == tx` 了, 不能再选 `(sx+sy, sx)` 及以下分支了.
        # 因此 `ty == sy+n*sx` => `(ty - sy) % sx == 0`
        # 另一方面, `sx == tx` => `(tx - sx) % sy == 0`
        if sx == tx or sy == ty:
            return (tx - sx) % sy == 0 and (ty - sy) % sx == 0
        # 由于坐标均为正数, `tx == ty` 路径不可逆, 因此如果原点目标点不重合, 目标点不可达.
        if tx == ty:
            return False
        elif tx > ty:
            tx = tx % ty
        else:
            ty = ty % tx
        return self.reachingPoints(sx, sy, tx, ty)

