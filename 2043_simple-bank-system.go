/*
 * @lc app=leetcode.cn id=simple-bank-system lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2043] 简易银行系统
 *
 * https://leetcode-cn.com/problems/simple-bank-system/description/
 *
 * algorithms
 * Medium (59.04%)
 * Total Accepted:    17.9K
 * Total Submissions: 26.7K
 * Testcase Example:  '["Bank","withdraw","transfer","deposit","transfer","withdraw"]\n' +
  '[[[10,100,20,50,30]],[3,10],[5,1,20],[5,20],[3,4,15],[10,50]]'
 *
 * 你的任务是为一个很受欢迎的银行设计一款程序，以自动化执行所有传入的交易（转账，存款和取款）。银行共有 n 个账户，编号从 1 到 n
 * 。每个账号的初始余额存储在一个下标从 0 开始的整数数组 balance 中，其中第 (i + 1) 个账户的初始余额是 balance[i] 。
 *
 * 请你执行所有 有效的 交易。如果满足下面全部条件，则交易 有效 ：
 *
 *
 * 指定的账户数量在 1 和 n 之间，且
 * 取款或者转账需要的钱的总数 小于或者等于 账户余额。
 *
 *
 * 实现 Bank 类：
 *
 *
 * Bank(long[] balance) 使用下标从 0 开始的整数数组 balance 初始化该对象。
 * boolean transfer(int account1, int account2, long money) 从编号为 account1
 * 的账户向编号为 account2 的账户转帐 money 美元。如果交易成功，返回 true ，否则，返回 false 。
 * boolean deposit(int account, long money) 向编号为 account 的账户存款 money
 * 美元。如果交易成功，返回 true ；否则，返回 false 。
 * boolean withdraw(int account, long money) 从编号为 account 的账户取款 money
 * 美元。如果交易成功，返回 true ；否则，返回 false 。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * ["Bank", "withdraw", "transfer", "deposit", "transfer", "withdraw"]
 * [[[10, 100, 20, 50, 30]], [3, 10], [5, 1, 20], [5, 20], [3, 4, 15], [10,
 * 50]]
 * 输出：
 * [null, true, true, true, false, false]
 *
 * 解释：
 * Bank bank = new Bank([10, 100, 20, 50, 30]);
 * bank.withdraw(3, 10);    // 返回 true ，账户 3 的余额是 $20 ，所以可以取款 $10 。
 * ⁠                        // 账户 3 余额为 $20 - $10 = $10 。
 * bank.transfer(5, 1, 20); // 返回 true ，账户 5 的余额是 $30 ，所以可以转账 $20 。
 * ⁠                        // 账户 5 的余额为 $30 - $20 = $10 ，账户 1 的余额为 $10 + $20 =
 * $30 。
 * bank.deposit(5, 20);     // 返回 true ，可以向账户 5 存款 $20 。
 * ⁠                        // 账户 5 的余额为 $10 + $20 = $30 。
 * bank.transfer(3, 4, 15); // 返回 false ，账户 3 的当前余额是 $10 。
 * ⁠                        // 所以无法转账 $15 。
 * bank.withdraw(10, 50);   // 返回 false ，交易无效，因为账户 10 并不存在。
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == balance.length
 * 1 <= n, account, account1, account2 <= 10^5
 * 0 <= balance[i], money <= 10^12
 * transfer, deposit, withdraw 三个函数，每个 最多调用 10^4 次
 *
 *
*/
type Bank struct {
	Accounts []int64
	Count    int
}

func Constructor(balance []int64) Bank {
	return Bank{Accounts: balance, Count: len(balance)}
}

func (b *Bank) Transfer(from, to int, money int64) bool {
	from--
	to--
	if from < 0 || from >= b.Count || to < 0 || to >= b.Count {
		return false
	}
	if money > b.Accounts[from] {
		return false
	}
	// 傻逼, 如果自己转自己怎么算?
	// b.Accounts[from], b.Accounts[to] = b.Accounts[from]-money, b.Accounts[to]+money
	if from != to {
		b.Accounts[from], b.Accounts[to] = b.Accounts[from]-money, b.Accounts[to]+money
	}
	return true
}

func (b *Bank) Deposit(to int, money int64) bool {
	to--
	if to < 0 || to >= b.Count {
		return false
	}
	b.Accounts[to] = b.Accounts[to] + money
	return true
}

func (b *Bank) Withdraw(from int, money int64) bool {
	from--
	if from < 0 || from >= b.Count {
		return false
	}
	if money > b.Accounts[from] {
		return false
	}
	b.Accounts[from] = b.Accounts[from] - money
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
