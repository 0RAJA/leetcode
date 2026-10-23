# 维护指南

日常只维护题解、测试和原始笔记。统计、索引和导航数据都从这些文件生成。

## 日常流程

1. 用 GoLand LeetCode Editor 生成 `leetcode/editor/cn/pXXXX/solution.go`，按需添加同目录的 `*_test.go`，运行改动包的 Go 测试。
2. 思路继续写在 `2026/README.md` 的顶层编号条目中，以页面题号和题名开头；也可写入 `leetcode/editor/cn/doc/note/` 下以题号、题名命名的非空 Markdown 文件。空占位和只有标题的路线条目不算题解笔记。
3. 将新文件加入 Git，提交并推送到 `master`。Actions 自动生成索引、运行检查、发布 GitHub Pages，并在需要时回写生成文件；拉取请求只检查，不发布。

提交前检查变更范围，保留题解、测试、笔记及必要的配置；题面缓存、IDE 状态、cookie 和 token 留在本地。提交信息使用 `feat:`、`fix:`、`docs:` 或 `chore:` 加简述。

## 需要修改哪里

| 内容 | 修改位置 |
| --- | --- |
| 题解与测试 | 原来的 Go 文件和同目录测试文件 |
| 思路、边界与易错点 | 原始笔记，不另抄一份目录 |
| 缺少映射或其他平台题 | `catalog.config.json` 的 `sources`，参照现有条目补充路径、`id`、`title`，按需补 `url`、`topics` |
| 题名、题型等修正 | 同配置的 `overrides`，键使用页面展示题号 |
| 精选比较路线（可选） | `docs/explorer/learning-paths.js`；每一步写清楚比较什么，相邻步骤说明变化理由 |

标准目录中的数字是插件内部 `questionId`，通过已入库的 `leetcode/editor/cn/all.json` 转换成页面题号，不要直接当成页面题号。该元数据文件由插件维护，索引不使用其中的通过状态。其他平台的编号与 LeetCode 重复时，笔记题名须与 `sources` 配置一致。普通题目关联按题型计算，无需逐题连线。

`docs/explorer/data.js` 和 `docs/CATALOG.md` 是生成结果，不直接编辑。导航统计只从题目记录计算；README 不保存题数、题型统计表或统计图片。收录表示有代码记录，“有测试”表示存在测试文件，都不代表在线判题或测试已通过。

## 本地预览（可选）

新题解和笔记先用 `git add <文件路径>` 纳入 Git，再运行：

```bash
python scripts/build_catalog.py
```

用浏览器打开 `docs/explorer/index.html`。生成脚本只读取 Git 已跟踪的相关文件，不扫描缓存；无需安装前端依赖。

修改导航或生成脚本后，运行：

```bash
python scripts/build_catalog.py --check
python -m unittest discover -s scripts -p "test_*.py"
node --test scripts/study-model.test.cjs
```

`--check` 只检查是否过期，不写文件；若提示过期，重新生成。Actions 也会运行这些检查。题解代码另跑改动包的 Go 测试。

## 遇到问题

- **新题没出现**：先确认文件已加入 Git，再检查元数据映射；无法识别的文件列在 [待补映射](CATALOG.md#待补映射)，按上表补配置后重新生成。
- **线上没更新**：查看仓库 Actions 的 `Build catalog and deploy Pages`；确认推送分支是 `master`、工作流已启用，Pages 来源是 GitHub Actions。工作流不能回写时，在本地生成并将生成文件一并提交。
- **生成文件有冲突**：保留正确的题解、笔记和配置后重新生成，不逐条手改索引。

插件模板与配置分别见[模板说明](leetcode-editor-template.md)、[配置示例](leetcode-editor-config.example.md)。历史题解保留在 `_legacy/`；需要继续维护时迁入标准题解目录。
