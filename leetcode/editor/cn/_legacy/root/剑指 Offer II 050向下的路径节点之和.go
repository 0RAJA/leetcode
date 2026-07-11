// 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
//
// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//
//
//
// 示例 1：
//
//
//
//
// 输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
// 输出：3
// 解释：和等于 8 的路径有 3 条，如图所示。
//
//
// 示例 2：
//
//
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// 输出：3
//
//
//
//
// 提示:
//
//
// 二叉树的节点个数的范围是 [0,1000]
// -10⁹
// -1000 <= targetSum <= 1000
//
//
//
//
// 注意：本题与主站 437 题相同：https://leetcode-cn.com/problems/path-sum-iii/
// Related Topics 树 深度优先搜索 二叉树 👍 33 👎 0

package main

func main() {

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
func pathSum(root *TreeNode, targetSum int) (ret int) {
	var dfs func(root *TreeNode, nums []int)
	dfs = func(root *TreeNode, nums []int) {
		if root == nil {
			return
		}
		nums = append(nums, root.Val)
		sum := 0
		for i := len(nums) - 1; i >= 0; i-- {
			sum += nums[i]
			if sum == targetSum {
				ret++
			}
		}
		length := len(nums)
		dfs(root.Left, nums[:length])
		dfs(root.Right, nums[:length])
	}
	nums := make([]int, 0, 1010)
	dfs(root, nums)
	return
}

// leetcode submit region end(Prohibit modification and deletion)
