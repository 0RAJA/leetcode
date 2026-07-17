// 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现
// 在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
//
//
// 示例 1:
//
//
// 输入: temperatures = [73,74,75,71,69,72,76,73]
// 输出: [1,1,4,2,1,1,0,0]
//
//
// 示例 2:
//
//
// 输入: temperatures = [30,40,50,60]
// 输出: [1,1,1,0]
//
//
// 示例 3:
//
//
// 输入: temperatures = [30,60,90]
// 输出: [1,1,0]
//
//
//
// 提示：
//
//
// 1 <= temperatures.length <= 10⁵
// 30 <= temperatures[i] <= 100
//
//
// Related Topics 栈 数组 单调栈 👍 2188 👎 0

package p0739

// leetcode submit region begin(Prohibit modification and deletion)
// 维护一个下标构成的单调非递增的栈，表示当前已经连续非升温的天气下标， [4,3,1] 遇到 2 -> [4,3,2]
// 新来一个天气和栈顶比较，发现大于栈顶温度，则说明升温了，则弹出更新间隔天数
func dailyTemperatures(temperatures []int) (res []int) {
	idxStack := make([]int, 0, len(temperatures))
	res = make([]int, len(temperatures))
	for idx, temperature := range temperatures {
		for len(idxStack) > 0 && temperatures[idxStack[len(idxStack)-1]] < temperature {
			lowTemperatureIdx := idxStack[len(idxStack)-1]
			idxStack = idxStack[:len(idxStack)-1]
			res[lowTemperatureIdx] = idx - lowTemperatureIdx
		}
		idxStack = append(idxStack, idx)
	}
	return
}

// leetcode submit region end(Prohibit modification and deletion)
