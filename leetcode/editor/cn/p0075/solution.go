// 给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地 对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
// 我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
//
//
//
// 必须在不使用库内置的 sort 函数的情况下解决这个问题。
//
//
//
// 示例 1：
//
//
// 输入：nums = [2,0,2,1,1,0]
//
//
// 输出：[0,0,1,1,2,2]
//
// 解释：
//
// 该数组包含两个 0、两个 1 和两个 2。将它们原地排序后，所有 0 排在最前面，接着是所有 1，最后是所有 2。
//
//
// 示例 2：
//
//
// 输入：nums = [2,0,1]
//
//
// 输出：[0,1,2]
//
// 解释：
//
// 数组中有且仅有一个 0、一个 1 和一个 2，按 0、1、2 的顺序原地排列。
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 300
// nums[i] 为 0、1 或 2
//
//
//
//
// 进阶：
//
//
// 你能想出一个仅使用常数空间的一趟扫描算法吗？
//
//
// Related Topics 数组 双指针 冒泡排序 排序 快速排序 👍 2137 👎 0

package p0075

// leetcode submit region begin(Prohibit modification and deletion)
// 这道题最朴素的想法就是：来一个元素发现是 0，我就把它放最左边；来一个元素发现是 2，我就把它放最右边；那么剩下其实就都是 1 了。
// 要划分为 4 个区域，使用左、中、右三条游标：
//  1. 左边：全部都是 0。
//  2. 左边到中间：全部都是 1。
//  3. 中间到右边：全部是待扫描区。
//  4. 右边到结尾：全部都是 2。
//
// 那我就是 left 和 mid 都从左往右，right 从右边往左边遍历。
//  1. 如果 mid 遇到的元素是 0，我就把它和 left 元素交换，然后 left++，mid++。\
//  2. 如果遇到元素是 1，那么就mid++就好了
//  2. 如果遇到元素是 2，我就把它和 right 元素交换。这时候 mid 不能动，right 要减减，因为 right 交换过来的这个元素对于 mid 而言它还没扫描过。
//     一直这样子往后轮询，直到 mid > right，就说明待扫描区没有了，剩下的都是已确定的
func sortColors(nums []int) {
	left, mid := 0, 0
	right := len(nums) - 1
	for mid <= right {
		if nums[mid] == 0 {
			nums[left], nums[mid] = nums[mid], nums[left]
			left++
			mid++
		} else if nums[mid] == 2 {
			nums[mid], nums[right] = nums[right], nums[mid]
			right--
		} else if nums[mid] == 1 {
			mid++
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)
