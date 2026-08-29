// 小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
//
// 除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接
// 相连的房子在同一天晚上被打劫 ，房屋将自动报警。
//
// 给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。
//
//
//
// 示例 1:
//
//
//
//
// 输入: root = [3,2,3,null,3,null,1]
// 输出: 7
// 解释: 小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7
//
// 示例 2:
//
//
//
//
// 输入: root = [3,4,5,1,3,null,1]
// 输出: 9
// 解释: 小偷一晚能够盗取的最高金额 4 + 5 = 9
//
//
//
//
// 提示：
//
//
//
//
// 树的节点数在 [1, 10⁴] 范围内
// 0 <= Node.val <= 10⁴
//
//
// Related Topics 树 深度优先搜索 动态规划 二叉树 树形 DP 👍 2196 👎 0

package p0337

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode submit region begin(Prohibit modification and deletion)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 后序遍历 返回两个值 一个是选择当前节点，一个是不选择当前节点 我们选择其中最大的那个
// 分别算出左右两边，偷或者不偷的值后 有几种可能性
// 1. 偷当前节点：root.Val + noRobValLeft + noRobValRight
// 2. 不偷当前节点：从一下可能中获取 max 值
//    1. 不偷左边+偷右边：noRobValLeft+robValRight
//    2. 不偷左边+不偷右边：noRobValLeft+noRobValRight
//    3. 偷左边+偷右边：robValLeft+robValRight
//    4. 偷左边+不偷右边：robValLeft+noRobValRight
func rob(root *TreeNode) int {
	var dfs func(root *TreeNode) (noRobVal, robVal int)
	dfs = func(root *TreeNode) (noRobVal, robVal int) {
		if root == nil {
			return 0, 0
		}
		noRobValLeft, robValLeft := dfs(root.Left)
		noRobValRight, robValRight := dfs(root.Right)
		return max(noRobValLeft+noRobValRight, noRobValLeft+robValRight, robValLeft+noRobValRight, robValLeft+robValRight), root.Val + noRobValLeft + noRobValRight
	}
	return max(dfs(root))
}

// leetcode submit region end(Prohibit modification and deletion)
