//你需要采用前序遍历的方式，将一个二叉树转换成一个由括号和整数组成的字符串。
//
// 空节点则用一对空括号 "()" 表示。而且你需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。
//
// 示例 1:
//
//
//输入: 二叉树: [1,2,3,4]
//       1
//     /   \
//    2     3
//   /
//  4
//
//输出: "1(2(4))(3)"
//
//解释: 原本将是“1(2(4)())(3())”，
//在你省略所有不必要的空括号对之后，
//它将是“1(2(4))(3)”。
//
//
// 示例 2:
//
//
//输入: 二叉树: [1,2,3,null,4]
//       1
//     /   \
//    2     3
//     \
//      4
//
//输出: "1(2()(4))(3)"
//
//解释: 和第一个示例相似，
//除了我们不能省略第一个对括号来中断输入和输出之间的一对一映射关系。
//
// Related Topics 树 深度优先搜索 字符串 二叉树 👍 313 👎 0

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println()
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *TreeNode) string {
	sb := strings.Builder{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		sb.WriteString(strconv.Itoa(root.Val))
		if root.Left != nil || root.Right != nil {
			sb.WriteByte('(')
			dfs(root.Left)
			sb.WriteByte(')')
		}
		if root.Right != nil {
			sb.WriteByte('(')
			dfs(root.Right)
			sb.WriteByte(')')
		}
	}
	dfs(root)
	return sb.String()
}

//leetcode submit region end(Prohibit modification and deletion)
