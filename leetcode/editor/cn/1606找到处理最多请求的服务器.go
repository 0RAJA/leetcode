//你有 k 个服务器，编号为 0 到 k-1 ，它们可以同时处理多个请求组。每个服务器有无穷的计算能力但是 不能同时处理超过一个请求 。请求分配到服务器的规则
//如下：
//
//
// 第 i （序号从 0 开始）个请求到达。
// 如果所有服务器都已被占据，那么该请求被舍弃（完全不处理）。
// 如果第 (i % k) 个服务器空闲，那么对应服务器会处理该请求。
// 否则，将请求安排给下一个空闲的服务器（服务器构成一个环，必要的话可能从第 0 个服务器开始继续找下一个空闲的服务器）。比方说，如果第 i 个服务器在忙，那
//么会查看第 (i+1) 个服务器，第 (i+2) 个服务器等等。
//
//
// 给你一个 严格递增 的正整数数组 arrival ，表示第 i 个任务的到达时间，和另一个数组 load ，其中 load[i] 表示第 i 个请求的工作
//量（也就是服务器完成它所需要的时间）。你的任务是找到 最繁忙的服务器 。最繁忙定义为一个服务器处理的请求数是所有服务器里最多的。
//
// 请你返回包含所有 最繁忙服务器 序号的列表，你可以以任意顺序返回这个列表。
//
//
//
// 示例 1：
//
//
//
//
//输入：k = 3, arrival = [1,2,3,4,5], load = [5,2,3,3,3]
//输出：[1]
//解释：
//所有服务器一开始都是空闲的。
//前 3 个请求分别由前 3 台服务器依次处理。
//请求 3 进来的时候，服务器 0 被占据，所以它呗安排到下一台空闲的服务器，也就是服务器 1 。
//请求 4 进来的时候，由于所有服务器都被占据，该请求被舍弃。
//服务器 0 和 2 分别都处理了一个请求，服务器 1 处理了两个请求。所以服务器 1 是最忙的服务器。
//
//
// 示例 2：
//
//
//输入：k = 3, arrival = [1,2,3,4], load = [1,2,1,2]
//输出：[0]
//解释：
//前 3 个请求分别被前 3 个服务器处理。
//请求 3 进来，由于服务器 0 空闲，它被服务器 0 处理。
//服务器 0 处理了两个请求，服务器 1 和 2 分别处理了一个请求。所以服务器 0 是最忙的服务器。
//
//
// 示例 3：
//
//
//输入：k = 3, arrival = [1,2,3], load = [10,12,11]
//输出：[0,1,2]
//解释：每个服务器分别处理了一个请求，所以它们都是最忙的服务器。
//
//
// 示例 4：
//
//
//输入：k = 3, arrival = [1,2,3,4,8,9,10], load = [5,2,10,3,1,2,2]
//输出：[1]
//
//
// 示例 5：
//
//
//输入：k = 1, arrival = [1], load = [1]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 1 <= k <= 10⁵
// 1 <= arrival.length, load.length <= 10⁵
// arrival.length == load.length
// 1 <= arrival[i], load[i] <= 10⁹
// arrival 保证 严格递增 。
//
// Related Topics 贪心 数组 有序集合 堆（优先队列） 👍 59 👎 0

package main

import (
	"container/heap"
	"fmt"

	"github.com/0RAJA/Rutils/struct/skip_list"
)

func main() {
	fmt.Println(busiestServers(1, []int{1}, []int{1}))
}

//leetcode submit region begin(Prohibit modification and deletion)
type pair1606 struct {
	end, id int
}
type PH struct {
	pairs []pair1606
}

func (p *PH) Len() int           { return len(p.pairs) }
func (p *PH) Less(i, j int) bool { return p.pairs[i].end < p.pairs[j].end }
func (p *PH) Swap(i, j int)      { p.pairs[i], p.pairs[j] = p.pairs[j], p.pairs[i] }
func (p *PH) Push(v interface{}) { p.pairs = append(p.pairs, v.(pair1606)) }
func (p *PH) Pop() interface{} {
	v := p.pairs[len(p.pairs)-1]
	p.pairs = p.pairs[:len(p.pairs)-1]
	return v
}

type M int

func (m M) CompareTo(comparable skip_list.Comparable) int {
	c, _ := comparable.(M)
	return int(m - c)
}

func busiestServers(k int, arrival []int, load []int) (ret []int) {
	var max int
	sList := skip_list.NewSkipList() //存放可用的服务器的跳表,按照id从小到大
	cnts := make([]int, k)           //统计服务器处理任务次数
	work := &PH{}                    //存放正在处理任务的服务器的优先队列,按照完成时间从小到大
	for i := 0; i < k; i++ {
		heap.Push(work, pair1606{id: i})
	}
	for i, arrivalTime := range arrival {
		for work.Len() > 0 && work.pairs[0].end <= arrivalTime { //先把所有已经完成任务的服务器放到slist中
			p := heap.Pop(work).(pair1606)
			sList.Put(skip_list.NewKV(M(p.id), p.end))
		}
		if sList.Len() == 0 { //如果没有可以使用的服务器就抛弃
			continue
		}
		var id int
		if kv, _ := sList.Get(M(i % k)); kv != nil { //如果能找到i%k或者它的下一个就直接使用，否则使用最小的那个
			id = int(kv.K.(M))
		} else {
			id = int(sList.GetMin().K.(M))
		}
		cnts[id]++
		if cnts[id] > max {
			max = cnts[id]
		}
		sList.Delete(M(id))       //记得从slist中删除已经被使用的服务器
		heap.Push(work, pair1606{ //加入到正在工作的服务器的优先队列中，完成时间是当前处理时间加上处理所需要的时间
			end: arrivalTime + load[i],
			id:  id,
		})
	}
	ret = make([]int, 0, len(cnts))
	for i, v := range cnts {
		if v == max {
			ret = append(ret, i)
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
