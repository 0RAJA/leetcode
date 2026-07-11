//在社交媒体网站上有 n 个用户。给你一个整数数组 ages ，其中 ages[i] 是第 i 个用户的年龄。
//
// 如果下述任意一个条件为真，那么用户 x 将不会向用户 y（x != y）发送好友请求：
//
//
// age[y] <= 0.5 * age[x] + 7
// age[y] > age[x]
// age[y] > 100 && age[x] < 100
//
//
// 否则，x 将会向 y 发送一条好友请求。
//
// 注意，如果 x 向 y 发送一条好友请求，y 不必也向 x 发送一条好友请求。另外，用户不会向自己发送好友请求。
//
// 返回在该社交媒体网站上产生的好友请求总数。
//
//
//
// 示例 1：
//
//
//输入：ages = [16,16]
//输出：2
//解释：2 人互发好友请求。
//
//
// 示例 2：
//
//
//输入：ages = [16,17,18]
//输出：2
//解释：产生的好友请求为 17 -> 16 ，18 -> 17 。
//
//
// 示例 3：
//
//
//输入：ages = [20,30,100,110,120]
//输出：3
//解释：产生的好友请求为 110 -> 100 ，120 -> 110 ，120 -> 100 。
//
//
//
//
// 提示：
//
//
// n == ages.length
// 1 <= n <= 2 * 10⁴
// 1 <= ages[i] <= 120
//
// Related Topics 数组 双指针 二分查找 排序 👍 160 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := []int{101, 56, 69, 48, 30}
	fmt.Println(numFriendRequests2(ages))
}

// age[y] <= 0.5 * age[x] + 7
// age[y] > age[x]
// age[y] > 100 && age[x] < 100
/*
1. 找到age[left]>0.5*age[i]+7且l<i的左边界left
2. 找到age[right]==age[i]的右边界right
*/
//leetcode submit region begin(Prohibit modification and deletion)
func numFriendRequests1(ages []int) (sum int) {
	sort.Ints(ages)
	//找到第一个满足条件的值
	search := func(n int, f func(idx int) bool) int {
		l, r := 0, n
		for l < r {
			mid := l + (r-l)/2
			if !f(mid) {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return l
	}
	for i := range ages {
		l := search(len(ages), func(idx int) bool { return float64(ages[idx]) > (0.5*float64(ages[i]) + 7) })
		r := search(len(ages), func(idx int) bool { return ages[idx] > ages[i] }) - 1
		if t := r - l; t > 0 {
			sum += t
		}
	}
	return sum
}

// 1. 找到age[left]>0.5*age[i]+7且l<i的左边界left
// 2. 找到age[right]==age[i]的右边界right
func numFriendRequests2(ages []int) (sum int) {
	sort.Ints(ages)
	l, r := 0, 0
	for _, age := range ages {
		if age < 15 { //由第一点和第二点得出a[left]必须是大于14的
			continue
		}
		for l < len(ages) && ages[l]*2 <= age+14 {
			l++
		}
		for r < len(ages) && ages[r] <= age {
			r++
		}
		sum += r - l - 1
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
