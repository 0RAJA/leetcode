//请实现一个 MyCalendar 类来存放你的日程安排。如果要添加的时间内没有其他安排，则可以存储这个新的日程安排。
//
// MyCalendar 有一个 book(int start, int end)方法。它意味着在 start 到 end 时间内增加一个日程安排，注意，这里
//的时间是半开区间，即 [start, end), 实数 x 的范围为， start <= x < end。
//
// 当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生重复预订。
//
// 每次调用 MyCalendar.book方法时，如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true。否则，返回 false 并且不要将该
//日程安排添加到日历中。
//
// 请按照以下步骤调用 MyCalendar 类: MyCalendar cal = new MyCalendar(); MyCalendar.book(
//start, end)
//
//
//
// 示例:
//
//
//输入:
//["MyCalendar","book","book","book"]
//[[],[10,20],[15,25],[20,30]]
//输出: [null,true,false,true]
//解释:
//MyCalendar myCalendar = new MyCalendar();
//MyCalendar.book(10, 20); // returns true
//MyCalendar.book(15, 25); // returns false ，第二个日程安排不能添加到日历中，因为时间 15 已经被第一个日程安排预
//定了
//MyCalendar.book(20, 30); // returns true ，第三个日程安排可以添加到日历中，因为第一个日程安排并不包含时间 20
//
//
//
//
//
//
// 提示：
//
//
// 每个测试用例，调用 MyCalendar.book 函数最多不超过 1000次。
// 0 <= start < end <= 10⁹
//
//
//
//
// 注意：本题与主站 729 题相同： https://leetcode-cn.com/problems/my-calendar-i/
// Related Topics 设计 线段树 有序集合 👍 20 👎 0

package main

import (
	"github.com/0RAJA/Rutils/struct/skip_list"
)

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type ZNode int

func (z ZNode) CompareTo(comparable skip_list.Comparable) int {
	return int(z - comparable.(ZNode))
}

type MyCalendar struct {
	list *skip_list.SkipList
}

func Constructor() MyCalendar {
	return MyCalendar{list: skip_list.NewSkipList()}
}

/*
	有序集合主要存放 start 和 end
	主要这几种情况
	1. 查找start找到重合的点p1
		1. p1是一个起始点
			return false
		2. p1是一个终止点
			找p1的下一个点
			1. 有下一个点p2（p2一定是一个起始点）
				1. p2>end
					删除p1，加入end
				2. p2==end
					删除p1，删除p2
				3. p2<end
					return false
			2. 没有下一个点
				删除p1,加入end
	2. 查找start没找到重合的点p1，并尝试查找下一个点p2
		1. p2 为 nil
			加入start,end
		2. p2 存在
			1. p2 为终止点
				return false
			2. p2 为起始点
				1. p2 < end
					return false
				2. p2 == end
					加入start,删除p2
				3. p2 > end
					加入start,end

*/
func (this *MyCalendar) Book(start int, end int) bool {
	res, ok := this.list.Get(ZNode(start))
	if ok {
		if res.V.(int) == 0 {
			return false
		} else {
			res1, _ := this.list.Get(ZNode(start + 1))
			if res1 != nil {
				if res1.K.(ZNode) < ZNode(end) {
					return false
				}
				if res1.K.(ZNode) == ZNode(end) {
					this.list.Delete(res.K)
					this.list.Delete(res1.K)
					return true
				}
				if res1.K.(ZNode) > ZNode(end) {
					this.list.Delete(res.K)
					this.list.Put(skip_list.NewKV(ZNode(end), 1))
					return true
				}
			} else {
				this.list.Delete(res.K)
				this.list.Put(skip_list.NewKV(ZNode(end), 1))
				return true
			}
		}
	} else {
		if res == nil {
			this.list.Put(skip_list.NewKV(ZNode(start), 0))
			this.list.Put(skip_list.NewKV(ZNode(end), 1))
			return true
		} else {
			if res.V.(int) == 1 {
				return false
			} else {
				if res.K.(ZNode) < ZNode(end) {
					return false
				}
				if res.K.(ZNode) == ZNode(end) {
					this.list.Put(skip_list.NewKV(ZNode(start), 0))
					this.list.Delete(res.K)
					return true
				}
				if res.K.(ZNode) > ZNode(end) {
					this.list.Put(skip_list.NewKV(ZNode(start), 0))
					this.list.Put(skip_list.NewKV(ZNode(end), 1))
					return true
				}
			}
		}
	}
	return false
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
//leetcode submit region end(Prohibit modification and deletion)
