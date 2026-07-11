// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
// 
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 
// 和 "192.168@1.1" 是 无效 IP 地址。
// 
//
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新
// 排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
//
// 
//
// 示例 1： 
//
// 
// 输入：s = "25525511135"
// 输出：["255.255.11.135","255.255.111.35"]
// 
//
// 示例 2： 
//
// 
// 输入：s = "0000"
// 输出：["0.0.0.0"]
// 
//
// 示例 3： 
//
// 
// 输入：s = "101023"
// 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 20 
// s 仅由数字组成 
// 
//
// Related Topics 字符串 回溯 👍 1606 👎 0
package main

import (
	"strconv"
	"strings"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 其实就是组合问题 切割一个符合条件的值 然后剩下的再切割 直到整体符合预期则加入结果
// 考虑两个点： 符合条件的（合法 ip 段）+ 剪枝（无法达成合法 ip 段或者ip 地址 就退出）
func restoreIpAddresses(s string) (ret []string) {
	// 判断当前部分是否是一个合法的 ip 段
	isValidIpPart := func(s string) bool {
		// 存在前导零 false
		if len(s) > 1 && s[0] == '0' {
			return false
		}
		// 满足 0 - 255
		num, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return num >= 0 && num <= 255
	}
	// 剪枝: 判断当前层是否需要直接 break 了，不能再往下构造了
	needToBreak := func(path []string, cur, remain string) bool {
		// 前面已经满了 or 当前 ip 段已经超过长度了 or 剩余的已经不够构造了 or 太多了分不下了
		// 剩余的 ip 段数（去掉了当前正在构造的）
		remainSlots := 4 - len(path) - 1
		if len(path) >= 4 || len(cur) >= 4 || len(remain) < remainSlots || len(remain)-(3-len(cur)) > remainSlots*3 {
			return true
		}
		return false
	}
	var dfs func(path []string, idx int)
	dfs = func(path []string, idx int) {
		// 刚好走到最后了说明当前 path 是一个合法的 ip 段组成的 ip 地址
		if idx == len(s) && len(path) == 4 {
			ret = append(ret, strings.Join(path, "."))
			return
		}
		for i := idx; i < len(s); i++ {
			cur := s[idx : i+1]
			// 当前组合无法构造出有效的结果直接就 break 剪枝
			if needToBreak(path, cur, s[i+1:]) {
				break
			}
			// 当前 ip 段合法则继续组合
			if isValidIpPart(cur) {
				dfs(append(path, cur), i+1)
			}
		}
	}
	dfs([]string{}, 0)
	return
}

// leetcode submit region end(Prohibit modification and deletion)
