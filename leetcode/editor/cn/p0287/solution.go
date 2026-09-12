// 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
//
// 假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
//
// 你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,3,4,2,2]
// 输出：2
//
//
// 示例 2：
//
//
// 输入：nums = [3,1,3,4,2]
// 输出：3
//
//
// 示例 3 :
//
//
// 输入：nums = [3,3,3,3,3]
// 输出：3
//
//
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁵
// nums.length == n + 1
// 1 <= nums[i] <= n
// nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次
//
//
//
//
// 进阶：
//
//
// 如何证明 nums 中至少存在一个重复的数字?
// 你可以设计一个线性级时间复杂度 O(n) 的解决方案吗？
//
//
// Related Topics 位运算 数组 双指针 二分查找 Floyd 判圈算法 抽屉原理 👍 2832 👎 0

package p0287

// leetcode submit region begin(Prohibit modification and deletion)
// 由于数字都在 1-N 内，且最多只有一个重复的，可以考虑把数组想象成链表 value 指向 下一个 key 的链表
// 转换为一个链表里面找相交的节点：先快慢指针找到重合判定有环，然后分别从头和重合处出发一个指针，两者同样速率重合点即为相交点
func findDuplicate(nums []int) (res int) {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	a, b := nums[0], nums[slow]
	for a != b {
		a = nums[a]
		b = nums[b]
	}
	return a
}

// leetcode submit region end(Prohibit modification and deletion)
