/*
 * Hard
 *
 * 给你一个文件 file ，并且该文件只能通过给定的 read4 方法来读取，请实现一个方法使其能够使 read 读取 n 个字符。注意：你的 read 方法可能会被调用多次。
 *
 * read4 的定义：
 *
 * read4 API 从文件中读取 4 个连续的字符，然后将这些字符写入缓冲区数组 buf4 。
 *
 * 返回值是读取的实际字符数。
 *
 * 请注意，read4() 有其自己的文件指针，类似于 C 中的 FILE * fp 。
 *
 *     参数类型: char[] buf4
 *     返回类型: int
 *
 * 注意: buf4[] 是目标缓存区不是源缓存区，read4 的返回结果将会复制到 buf4[] 当中。
 * 下列是一些使用 read4 的例子：
 *
 *
 *
 * File file("abcde"); // 文件名为 "abcde"， 初始文件指针 (fp) 指向 'a'
 * char[] buf4 = new char[4]; // 创建一个缓存区使其能容纳足够的字符
 * read4(buf4); // read4 返回 4。现在 buf4 = "abcd"，fp 指向 'e'
 * read4(buf4); // read4 返回 1。现在 buf4 = "e"，fp 指向文件末尾
 * read4(buf4); // read4 返回 0。现在 buf4 = ""，fp 指向文件末尾
 *
 *
 * read 方法：
 *
 * 通过使用 read4 方法，实现 read 方法。该方法可以从文件中读取 n 个字符并将其存储到缓存数组 buf 中。您 不能 直接操作 file 。
 *
 * 返回值为实际读取的字符。
 *
 * read 的定义：
 *
 *     参数类型:  char[] buf, int n
 *     返回类型:  int
 *
 * 注意: buf[] 是目标缓存区不是源缓存区，你需要将结果写入 buf[] 中。
 * 注意：
 *
 * 你 不能 直接操作该文件，文件只能通过 read4 获取而 不能 通过 read。
 * read  函数可以被调用 多次。
 * 请记得 重置 在 Solution 中声明的类变量（静态变量），因为类变量会 在多个测试用例中保持不变，影响判题准确。请 查阅 这里。
 * 你可以假定目标缓存数组 buf 保证有足够的空间存下 n 个字符。
 * 保证在一个给定测试用例中，read 函数使用的是同一个 buf。
 *
 *
 * 示例 1：
 *
 * 输入： file = "abc"， queries = [1,2,1]
 * 输出：[1,2,0]
 * 解释：测试用例表示以下场景:
 * File file("abc");
 * Solution sol;
 * sol.read (buf, 1); // 调用 read 方法后，buf 应该包含 “a”。我们从文件中总共读取了 1 个字符，所以返回 1。
 * sol.read (buf, 2); // 现在 buf 应该包含 "bc"。我们从文件中总共读取了 2 个字符，所以返回 2。
 * sol.read (buf, 1); // 我们已经到达文件的末尾，不能读取更多的字符。所以返回 0。
 * 假设已经分配了 buf ，并保证有足够的空间存储文件中的所有字符。
 * 示例 2：
 *
 * 输入：file = "abc"， queries = [4,1]
 * 输出：[3,0]
 * 解释：测试用例表示以下场景:
 * File file("abc");
 * Solution sol;
 * sol.read (buf, 4); // 调用 read 方法后，buf 应该包含 “abc”。我们从文件中总共读取了 3 个字符，所以返回 3。
 * sol.read (buf, 1); // 我们已经到达文件的末尾，不能读取更多的字符。所以返回 0。
 *
 *
 * 提示：
 *
 * 1 <= file.length <= 500
 * file 由英语字母和数字组成
 * 1 <= queries.length <= 10
 * 1 <= queries[i] <= 500
 * 通过次数3,519
 * 提交次数5,923
 */

/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf4 []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf4 := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf4) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf4) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf4) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */

var readn = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	var buffer []byte
	return func(buf []byte, n int) int {
		// assert buffer[0:k] is available.
		for len(buffer) < n {
			tmpBuff := make([]byte, 4)
			v := read4(tmpBuff)
			if v < 4 {
				buffer = append(buffer, tmpBuff[:v]...)
				break
			}
			buffer = append(buffer, tmpBuff[:]...)
		}
		// assert len(buffer) >= n or no other bytes to read.
		k := n
		if len(buffer) < n {
			k = len(buffer)
		}
		for i := 0; i < k; i++ {
			buf[i] = buffer[i]
		}
		buffer = buffer[k:]
		return k
	}
}

var solution = readn

/*
 * var solution = func(read4 func([]byte) int) func([]byte, int) int {
 *     // implement read below.
 *     var buffer []byte
 *     return func(buf []byte, n int) int {
 *         for len(buffer) < n {
 *             tmp := make([]byte, 4)
 *             v := read4(tmp)
 *             buffer = append(buffer, tmp[:v]...)
 *             if v < 4 {
 *                 break
 *             }
 *         }
 *         k := min(n, len(buffer))
 *         copy(buf, buffer[:k])
 *         buffer = buffer[k:]
 *         return k
 *     }
 * }
 *
 * func min(x, y int) int {
 *     if x < y {
 *         return x
 *     }
 *     return y
 * }
 */
