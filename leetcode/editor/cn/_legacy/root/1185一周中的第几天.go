//给你一个日期，请你设计一个算法来判断它是对应一周中的哪一天。
//
// 输入为三个整数：day、month 和 year，分别表示日、月、年。
//
// 您返回的结果必须是这几个值中的一个 {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday",
//"Friday", "Saturday"}。
//
//
//
// 示例 1：
//
// 输入：day = 31, month = 8, year = 2019
//输出："Saturday"
//
//
// 示例 2：
//
// 输入：day = 18, month = 7, year = 1999
//输出："Sunday"
//
//
// 示例 3：
//
// 输入：day = 15, month = 8, year = 1993
//输出："Sunday"
//
//
//
//
// 提示：
//
//
// 给出的日期一定是在 1971 到 2100 年之间的有效日期。
//
// Related Topics 数学 👍 80 👎 0

package main

import "fmt"

func main() {
	day, month, year := 3, 1, 2022
	fmt.Println(dayOfTheWeek(day, month, year))
}

// leetcode submit region begin(Prohibit modification and deletion)
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func daysOfMonth(year, month int) int {
	if month == 2 {
		if isLeapYear(year) {
			return 29
		}
		return 28
	}
	m := map[int]struct{}{1: {}, 3: {}, 5: {}, 7: {}, 8: {}, 10: {}, 12: {}}
	if _, ok := m[month]; ok {
		return 31
	}
	return 30
}
func daysOfYear(year int) int {
	if isLeapYear(year) {
		return 366
	}
	return 365
}
func week(days int) string {
	m := [7]string{"Thursday", "Friday", "Saturday", "Sunday", "Monday", "Tuesday", "Wednesday"}
	return m[days%7]
}
func dayOfTheWeek(day int, month int, year int) string {
	sum := 0
	for i := 1971; i < year; i++ {
		sum += daysOfYear(i)
	}
	for i := 1; i < month; i++ {
		sum += daysOfMonth(year, i)
	}
	sum += day
	return week(sum)
}

//leetcode submit region end(Prohibit modification and deletion)
