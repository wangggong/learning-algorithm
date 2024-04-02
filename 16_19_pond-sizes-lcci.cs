/*
 * Medium
 * 
 * 你有一个用于表示一片土地的整数矩阵land，该矩阵中每个点的值代表对应地点的海拔高度。若值为0则表示水域。由垂直、水平或对角连接的水域为池塘。池塘的大小是指相连接的水域的个数。编写一个方法来计算矩阵中所有池塘的大小，返回值需要从小到大排序。
 * 
 * 示例：
 * 
 * 输入：
 * [
 *   [0,2,1,0],
 *   [0,1,0,1],
 *   [1,1,0,1],
 *   [0,1,0,1]
 * ]
 * 输出： [1,2,4]
 * 提示：
 * 
 * 0 < len(land) <= 1000
 * 0 < len(land[i]) <= 1000
 */

public class Solution
{
    private int dfs(int[][] land, int x, int y) => 0 <= x && x < land.Length
        && 0 <= y && y < land[0].Length
        && land[x][y] is 0
            ? (land[x][y] = 1) + Enumerable
                .Range(-1, 3)
                .SelectMany(dx => Enumerable
                    .Range(-1, 3)
                    .Select(dy => dfs(land, x + dx, y + dy)))
                .Sum()
            : 0;
    
    public int[] PondSizes(int[][] land) => Enumerable
        .Range(0, land.Length)
        .SelectMany(x => Enumerable
            .Range(0, land[0].Length)
            .Where(y => land[x][y] is 0)
            .Select(y => dfs(land, x, y)))
        .OrderBy(x => x)
        .ToArray();
}
