// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
//
//
//
// 示例 1:
//
//
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]
//
//
// 示例 2:
//
//
// 输入: nums = [0]
// 输出: [0]
//
//
//
// 提示:
//
//
//
// 1 <= nums.length <= 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
//
//
//
//
// 进阶：你能尽量减少完成的操作次数吗？
//
// Related Topics 数组 双指针 👍 3006 👎 0

package p0283

// leetcode submit region begin(Prohibit modification and deletion)
// 双指针：low 指针指向 0，fast 指针指向 非0，两者交换
// 注意细节点：low < fast && fast < len(nums)；在此基础上满足条件则交换，否则推动继续
func moveZeroes(nums []int) {
	low, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] == 0 {
			fast++
			continue
		}
		if nums[low] != 0 && low < fast {
			low++
			continue
		}
		if low == fast {
			fast++
		} else {
			nums[low], nums[fast] = nums[fast], nums[low]
		}
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
