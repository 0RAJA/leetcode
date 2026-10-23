# LeetCode · 题解与复习笔记

用 Go 记录做过的题，把题目、题型、代码和当时的思路连起来，方便以后回顾。

**[在线复习导航](https://0raja.github.io/leetcode/)** · **[查看题解索引](docs/CATALOG.md)** · **[原始复习笔记](2026/README.md)** · **[返回博客](https://blog.humraja.top/)**

<!-- catalog:start -->

![题解概览](docs/assets/catalog.svg)

当前维护 **80** 题 · 历史归档 **487** 题 · 合并去重 **567** 题

其中 **125** 题关联了思路笔记，**19** 题有测试文件。

| 当前题型 | 收录题数 | 当前题型 | 收录题数 |
| --- | ---: | --- | ---: |
| 数组 | 45 | 动态规划 | 22 |
| 哈希表 | 18 | 字符串 | 17 |
| 双指针 | 16 | 排序 | 10 |
| 链表 | 10 | 深度优先搜索 | 9 |
| 栈 | 8 | 广度优先搜索 | 8 |
| 二分查找 | 6 | 二叉树 | 6 |

同题多版本只统计一次，题型可以重叠。以上为仓库代码记录，未读取或推断在线 AC 状态。

<!-- catalog:end -->

## 怎么回顾

- **在 GitHub 看**：[题解索引](docs/CATALOG.md)按当前维护、历史归档整理，直接跳转到题目、Go 代码、笔记和测试。
- **在线复习**：打开 [GitHub Pages 导航](https://0raja.github.io/leetcode/)，使用最新的题型索引与复习路径。
- **在本地看**：下载或克隆仓库后，用浏览器打开 [`docs/explorer/index.html`](docs/explorer/index.html)。可以搜索题号/题名，按题型、难度筛选，在详情中一起查看思路和代码。
- **留到下次看**：在导航页把题加入待复习，标记保存在当前浏览器，可导出备份；不会作为个人动态发布。
- **连着比较**：切换“关联路径”，对照同一思路下的边界和状态变化；详情会解释相邻题为什么值得一起复习。
- **先想再看**：详情先给回忆线索，笔记和高亮代码主动展开；用“需要再练／能独立写出”记录本次自评。

## 平时怎么记录

继续使用 GoLand LeetCode Editor，一题一个目录：

```text
leetcode/editor/cn/pXXXX/
  solution.go
  solution_test.go       # 按需补充
```

题库元数据负责把插件内部编号转换成页面题号；题型从代码里的 `Related Topics` 读取。同题的新旧实现放在同一条记录中，避免重复统计。

你原来的 [`2026/README.md`](2026/README.md)仍然是复习笔记入口，已有思路会自动关联；也支持写在 `leetcode/editor/cn/doc/note/` 下的非空 Markdown 笔记。不需要再手写一份题目目录。

本地更新导航：

```bash
python scripts/build_catalog.py
```

GitHub Actions 在推送到 `master` 后自动刷新概览、题解索引和 GitHub Pages。运行方式及提交边界见[导航维护说明](docs/catalog-guide.md)和[仓库维护约定](docs/maintenance.md)。

## 目录

| 路径 | 内容 |
| --- | --- |
| [`leetcode/editor/cn/pXXXX/`](leetcode/editor/cn/) | 当前维护的题解，每题独立 package |
| [`2026/`](2026/) | 刷题路线笔记和独立练习 |
| [`leetcode/editor/cn/_legacy/`](leetcode/editor/cn/_legacy/) | 历史题解，在导航中保留版本入口 |
| [`docs/explorer/`](docs/explorer/) | 无需安装依赖的静态题解导航 |
| [`catalog.config.json`](catalog.config.json) | 自定义题目、缺失元数据和题型修正 |

题面缓存、登录态和 IDE 本机配置继续保留在本地，提交边界见[仓库维护约定](docs/maintenance.md)。

## 验证

```bash
go test ./2026/...
go test ./leetcode/editor/cn/...
python -m unittest discover -s scripts -p "test_*.py"
node --test scripts/study-model.test.cjs
python scripts/build_catalog.py --check
```

索引只表示仓库中有对应代码，不代表在线判题通过；“有测试”只表示存在测试文件。历史代码维持归档，不纳入当前 Go 包的测试。
