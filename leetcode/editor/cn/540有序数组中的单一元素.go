//给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。
//
// 请你找出并返回只出现一次的那个数。
//
// 你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,1,2,3,3,4,4,8,8]
//输出: 2
//
//
// 示例 2:
//
//
//输入: nums =  [3,3,7,7,10,11,11]
//输出: 10
//
//
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 10⁵
// 0 <= nums[i] <= 10⁵
//
// Related Topics 数组 二分查找 👍 336 👎 0

package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 2}
	fmt.Print(singleNonDuplicate2(nums))
}

//leetcode submit region begin(Prohibit modification and deletion)
func singleNonDuplicate(nums []int) (ret int) {
	var search func(i, j int) int
	search = func(i, j int) int {
		mid := (i + j) / 2                       //中间的数
		if mid > 0 && nums[mid] == nums[mid-1] { //看是否和左边匹配上
			if (j-mid)%2 != 0 { //右边是奇数去右边找
				return search(mid+1, j)
			}
			return search(i, mid-2)
		} else if mid < len(nums)-1 && nums[mid] == nums[mid+1] { //看是否和右边匹配上
			if (mid-i)%2 != 0 { //左边是奇数去左边找
				return search(i, mid-1)
			}
			return search(mid+2, j)
		} else { //都匹配不上就是单独的数
			return nums[mid]
		}
	}
	return search(0, len(nums)-1)
}

func singleNonDuplicate2(nums []int) (ret int) {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if mid > 0 && nums[mid] == nums[mid-1] { //看是否和左边匹配上
			if (r-mid)%2 != 0 { //右边是奇数去右边找
				l = mid + 1
			} else {
				r = mid - 2
			}
		} else if mid < len(nums)-1 && nums[mid] == nums[mid+1] { //看是否和右边匹配上
			if (mid-l)%2 != 0 { //左边是奇数去左边找
				r = mid - 1
			} else {
				l = mid + 2
			}
		} else { //都匹配不上就是单独的数
			return nums[mid]
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
