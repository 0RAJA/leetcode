//假设有一个同时存储文件和目录的文件系统。下图展示了文件系统的一个示例：
//
//
//
// 这里将 dir 作为根目录中的唯一目录。dir 包含两个子目录 subdir1 和 subdir2 。subdir1 包含文件 file1.ext 和子目
//录 subsubdir1；subdir2 包含子目录 subsubdir2，该子目录下包含文件 file2.ext 。
//
// 在文本格式中，如下所示(⟶表示制表符)：
//
//
//dir
//⟶ subdir1
//⟶ ⟶ file1.ext
//⟶ ⟶ subsubdir1
//⟶ subdir2
//⟶ ⟶ subsubdir2
//⟶ ⟶ ⟶ file2.ext
//
//
// 如果是代码表示，上面的文件系统可以写为 "dir
//\tsubdir1
//\t\tfile1.ext
//\t\tsubsubdir1
//\tsubdir2
//\t\tsubsubdir2
//\t\t\tfile2.ext" 。'
//' 和 '\t' 分别是换行符和制表符。
//
// 文件系统中的每个文件和文件夹都有一个唯一的 绝对路径 ，即必须打开才能到达文件/目录所在位置的目录顺序，所有路径用 '/' 连接。上面例子中，指向
//file2.ext 的 绝对路径 是 "dir/subdir2/subsubdir2/file2.ext" 。每个目录名由字母、数字和/或空格组成，每个文件名遵循
//name.extension 的格式，其中 name 和 extension由字母、数字和/或空格组成。
//
// 给定一个以上述格式表示文件系统的字符串 input ，返回文件系统中 指向 文件 的 最长绝对路径 的长度 。 如果系统中没有文件，返回 0。
//
//
//
// 示例 1：
//
//
//输入：input = "dir
//\tsubdir1
//\tsubdir2
//\t\tfile.ext"
//输出：20
//解释：只有一个文件，绝对路径为 "dir/subdir2/file.ext" ，路径长度 20
//
//
// 示例 2：
//
//
//输入：input = "dir
//\tsubdir1
//\t\tfile1.ext
//\t\tsubsubdir1
//\tsubdir2
//\t\tsubsubdir2
//\t\t\tfile2.ext"
//输出：32
//解释：存在两个文件：
//"dir/subdir1/file1.ext" ，路径长度 21
//"dir/subdir2/subsubdir2/file2.ext" ，路径长度 32
//返回 32 ，因为这是最长的路径
//
// 示例 3：
//
//
//输入：input = "a"
//输出：0
//解释：不存在任何文件
//
// 示例 4：
//
//
//输入：input = "file1.txt
//file2.txt
//longfile.txt"
//输出：12
//解释：根目录下有 3 个文件。
//因为根目录中任何东西的绝对路径只是名称本身，所以答案是 "longfile.txt" ，路径长度为 12
//
//
//
//
// 提示：
//
//
// 1 <= input.length <= 10⁴
// input 可能包含小写或大写的英文字母，一个换行符 '
//'，一个制表符 '\t'，一个点 '.'，一个空格 ' '，和数字。
//
// Related Topics 栈 深度优先搜索 字符串 👍 154 👎 0

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"))
}

//leetcode submit region begin(Prohibit modification and deletion)

/*
就是先通过\n分割不同的文件(或目录),通过每个文件的\t的数量判断是第几层
可以使用深搜，每次往后遍历

	如果是下一层就先判断是不是文件，如果不是文件就进到下一层继续搜，否则就继续往后搜
	如果不是下一层就返回，让上一层继续从这个往后深搜
*/
func lengthLongestPath(input string) (ret int) {
	paths := strings.Split(input, "\n")   //路径以\n来进行分割
	var dfs func(idx, level, sum int) int //分别代表当前路径在paths[idx]层，层数为level,此层包括之前的长度为sum
	dfs = func(idx, level, sum int) int {
		for i := idx + 1; i < len(paths); { //从下一个开始查找
			pos := strings.LastIndexByte(paths[i], '\t') + 1
			name := paths[i][pos:]
			if pos == level+1 { //发现这个是下一层的文件或目录
				if strings.Count(name, ".") > 0 { //判断是不是文件
					if sum+len(name) > ret {
						ret = sum + len(name)
					}
					i++
				} else { //如果是目录就进入看看
					i = dfs(i, pos, sum+len(name)+1)
				}
			} else { //发现这个不是这一层的就返回去
				return i
			}
		}
		return len(paths)
	}
	dfs(-1, -1, 0) //-1表示根路径，这样才能把第0层全部搜完
	return
}

//leetcode submit region end(Prohibit modification and deletion)
