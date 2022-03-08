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
)

func main() {
	//testMap()
	//testMMM()
	testSearch()
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
