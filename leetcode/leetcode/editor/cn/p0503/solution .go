// 给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的 下一个更大元素
// 。
//
// 数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1
// 。
//
//
//
// 示例 1:
//
//
// 输入: nums = [1,2,1]
// 输出: [2,-1,2]
// 解释: 第一个 1 的下一个更大的数是 2；
// 数字 2 找不到下一个更大的数；
// 第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
//
//
// 示例 2:
//
//
// 输入: nums = [1,2,3,4,3]
// 输出: [2,3,4,-1,4]
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
// Related Topics 栈 数组 单调栈 👍 1115 👎 0

package p0503

// leetcode submit region begin(Prohibit modification and deletion)
// 维护一个非单调递增的栈，以及一个记录当前idx是否访问过的map
// 循环遍历nums，
// 1. 每个 num 尝试更新栈中元素
// 2. 判断当前元素还在栈里就停止；没在栈里就继续
func nextGreaterElements(nums []int) (res []int) {
	resMap := make(map[int]int, len(nums))
	idxStack := make([]int, 0, len(nums))
	stackMap := make(map[int]bool, len(nums))
	for i := 0; i <= len(nums); i++ {
		if i == len(nums) {
			i = -1
			continue
		}
		for len(idxStack) > 0 && nums[idxStack[len(idxStack)-1]] < nums[i] {
			popIdx := idxStack[len(idxStack)-1]
			delete(stackMap, popIdx)
			idxStack = idxStack[:len(idxStack)-1]
			resMap[popIdx] = nums[i]
		}
		if _, ok := stackMap[i]; ok {
			break
		}
		stackMap[i] = true
		idxStack = append(idxStack, i)
	}
	for idx := range nums {
		if v, ok := resMap[idx]; ok {
			res = append(res, v)
		} else {
			res = append(res, -1)
		}
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
