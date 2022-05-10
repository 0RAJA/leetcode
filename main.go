/*
 * @Author: Raja
 * @Description:
 * @Date: 2022-02-12 07:17:21
 * @LastEditTime: 2022-02-14 04:34:06
 * @FilePath: /leetcode/main.go
 */
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// testMap()
	// testMMM()
	// testSearch()
	// testSlice()
	testSort()
}

func testSort() {
	s := []string{"az", "bz", "cz", "dz", "zzzzz"}
	sort.Slice(s, func(i, j int) bool {
		return strings.Compare(s[i], s[j]) < 0
	})
	fmt.Println(s)
}

func testSlice() {
	var nums []int
	test := func(nums []int) {
		nums = append(nums, 1, 2, 3, 4)
	}
	test(nums)
	fmt.Println(nums)
}
func testSearch() {
	/*
		0,1,2,3,4
	*/
	nums := []int{1, 3, 4}
	fmt.Println(sort.SearchInts(nums, 2))
}

func testMap() {
	type Num struct {
		a, b int
	}
	m := make(map[Num]bool)
	fmt.Println(m[Num{a: 1, b: 2}])
	m[Num{a: 1, b: 2}] = true
	fmt.Println(m[Num{a: 1, b: 2}])
}

func testMMM() {
	fmt.Println("testMaa")
}
