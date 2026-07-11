//求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
//
//
//
// 示例 1：
//
// 输入: n = 3
//输出: 6
//
//
// 示例 2：
//
// 输入: n = 9
//输出: 45
//
//
//
//
// 限制：
//
//
// 1 <= n <= 10000
//
// Related Topics 位运算 递归 脑筋急转弯 👍 417 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func sumNums(n int) (ret int) {
	var dfs func(n int) bool
	dfs = func(n int) bool {
		ret += n
		return n > 0 && dfs(n-1)
	}
	dfs(n)
	return
}

//leetcode submit region end(Prohibit modification and deletion)
