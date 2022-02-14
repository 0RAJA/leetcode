/*
 * @Author: Raja
 * @Description:
 * @Date: 2022-02-12 07:17:21
 * @LastEditTime: 2022-02-14 04:34:06
 * @FilePath: /leetcode/main.go
 */
package main

import "fmt"

func main() {
	testMap()
	testMMM()
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
