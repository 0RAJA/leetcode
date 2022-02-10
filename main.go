package main

import "fmt"

func main() {
	testMap()
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
