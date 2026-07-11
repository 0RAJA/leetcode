// 给你一个数组 orders，表示客户在餐厅中完成的订单，确切地说， orders[i]=[customerNamei,tableNumberi,foodIt
// emi] ，其中 customerNamei 是客户的姓名，tableNumberi 是客户所在餐桌的桌号，而 foodItemi 是客户点的餐品名称。
//
// 请你返回该餐厅的 点菜展示表 。在这张表中，表中第一行为标题，其第一列为餐桌桌号 “Table” ，后面每一列都是按字母顺序排列的餐品名称。接下来每一行中
// 的项则表示每张餐桌订购的相应餐品数量，第一列应当填对应的桌号，后面依次填写下单的餐品数量。
//
// 注意：客户姓名不是点菜展示表的一部分。此外，表中的数据行应该按餐桌桌号升序排列。
//
// 示例 1：
//
// 输入：orders = [["David","3","Ceviche"],["Corina","10","Beef Burrito"],["David",
// "3","Fried Chicken"],["Carla","5","Water"],["Carla","5","Ceviche"],["Rous","3","
// Ceviche"]]
// 输出：[["Table","Beef Burrito","Ceviche","Fried Chicken","Water"],["3","0","2","1
// ","0"],["5","0","1","0","1"],["10","1","0","0","0"]]
// 解释：
// 点菜展示表如下所示：
// Table,Beef Burrito,Ceviche,Fried Chicken,Water
// 3    ,0           ,2      ,1            ,0
// 5    ,0           ,1      ,0            ,1
// 10   ,1           ,0      ,0            ,0
// 对于餐桌 3：David 点了 "Ceviche" 和 "Fried Chicken"，而 Rous 点了 "Ceviche"
// 而餐桌 5：Carla 点了 "Water" 和 "Ceviche"
// 餐桌 10：Corina 点了 "Beef Burrito"
//
// 示例 2：
//
// 输入：orders = [["James","12","Fried Chicken"],["Ratesh","12","Fried Chicken"],[
// "Amadeus","12","Fried Chicken"],["Adam","1","Canadian Waffles"],["Brianna","1","
// Canadian Waffles"]]
// 输出：[["Table","Canadian Waffles","Fried Chicken"],["1","2","0"],["12","0","3"]]
//
// 解释：
// 对于餐桌 1：Adam 和 Brianna 都点了 "Canadian Waffles"
// 而餐桌 12：James, Ratesh 和 Amadeus 都点了 "Fried Chicken"
//
// 示例 3：
//
// 输入：orders = [["Laura","2","Bean Burrito"],["Jhon","2","Beef Burrito"],["Melis
// sa","2","Soda"]]
// 输出：[["Table","Bean Burrito","Beef Burrito","Soda"],["2","1","1","1"]]
//
// 提示：
//
// 1 <= orders.length <= 5 * 10^4
// orders[i].length == 3
// 1 <= customerNamei.length, foodItemi.length <= 20
// customerNamei 和 foodItemi 由大小写英文字母及空格字符 ' ' 组成。
// tableNumberi 是 1 到 500 范围内的整数。
//
// Related Topics 数组 哈希表 字符串 有序集合 排序
// 👍 52 👎 0
package main

import (
	"sort"
	"strconv"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type MyN struct {
	dish map[string]int //每个桌子的菜的数量
}

func displayTable(orders [][]string) (ret [][]string) {
	id := make([]int, 0)           //所有桌子的id
	table := make(map[int]MyN)     //所有桌子
	dishSlice := make([]string, 0) //所有菜名
	dish := make(map[string]bool)  //用来判断这个菜名是否重复出现过
	for _, v := range orders {
		myId, _ := strconv.Atoi(v[1])
		x, ok := table[myId] //看看这个桌子id是否出现过
		if ok == false {     //没出现过就加到id里面
			id = append(id, myId)
			table[myId] = MyN{dish: make(map[string]int, 0)}
			x = table[myId]
		}
		if dish[v[2]] == false { //看看菜名是否出现过,没出现过就加到dishSlice里面
			dishSlice = append(dishSlice, v[2])
			dish[v[2]] = true
		}
		x.dish[v[2]]++ //此桌子的菜的数量+1
	}
	sort.Slice(id, func(i, j int) bool { //给id排序
		return id[i] < id[j]
	})
	sort.Slice(dishSlice, func(i, j int) bool { //给菜名排序
		return dishSlice[i] < dishSlice[j]
	})
	ret = make([][]string, len(id)+1) //填充返回值
	ret = append(ret)
	ret[0] = append(ret[0], "Table")
	ret[0] = append(ret[0], dishSlice...)
	for i, myId := range id {
		ret[i+1] = append(ret[i+1], strconv.Itoa(myId))
		for _, s := range dishSlice {
			ret[i+1] = append(ret[i+1], strconv.Itoa(table[myId].dish[s]))
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
