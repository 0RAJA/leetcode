// Curated review routes, independent of the automatically generated catalog.
// These are small comparison exercises based on the repository's current Go
// solutions and 2026 notes, not an official syllabus or a record of mastery.
// Maintainers: keep only a few useful routes here. Add a step only when its
// connection explains which invariant, state or boundary changes; use the
// displayed frontend question ID. New solutions still enter the main catalog
// automatically and do not require a route entry. Tests check every reference.
window.LEETCODE_PATHS = {
  "version": 1,
  "paths": [
    {
      "id": "binary-search",
      "title": "二分：每次排除哪一半",
      "description": "先固定搜索区间，再分别改变命中后的处理方式和有序性的判断条件。",
      "steps": [
        {
          "id": "704",
          "focus": "维护搜索区间，明确左右端点是否可达"
        },
        {
          "id": "34",
          "focus": "命中后继续收缩，分别找到左右边界",
          "connection": "704 命中即可返回；34 存在重复值，命中后仍要朝一侧收缩。对照两次二分的更新条件，说明为什么不会漏掉最左或最右的位置。"
        },
        {
          "id": "33",
          "focus": "先识别有序半边，再决定排除范围",
          "connection": "沿用排除半个区间的思路，但旋转数组不再整体有序。先判断哪半边有序，再判断 target 是否落在其中，比较这一步如何替代普通二分的大小判断。"
        }
      ]
    },
    {
      "id": "sum-search",
      "title": "求和：从补数到多元组合",
      "description": "比较哈希补数与有序双指针，再把多元求和拆成外层枚举和内层搜索。",
      "steps": [
        {
          "id": "1",
          "focus": "查找 target − 当前值，记录已出现元素"
        },
        {
          "id": "15",
          "focus": "固定一个数，用有序双指针寻找另外两个",
          "connection": "把两数补数查询换成排序后的相向双指针，再用外层枚举固定第三个数。说明和偏小或偏大时移动哪个指针，并区分索引答案与需要去重的数值组合。"
        },
        {
          "id": "18",
          "focus": "再固定一个数，复用内层双指针和分层去重",
          "connection": "四数之和增加一层固定元素的枚举，剩余两数仍按和的大小推进。对照代码中每层枚举与左右指针的去重位置，说明它们分别防止哪一种重复组合。"
        }
      ]
    },
    {
      "id": "linked-fast-slow",
      "title": "链表：快慢指针停在哪里",
      "description": "同样的一快一慢，可以服务于判环、找入口、定位中点和分治拆分。重点比较终止条件。",
      "steps": [
        {
          "id": "141",
          "focus": "区分快慢相遇与快指针到达链尾"
        },
        {
          "id": "142",
          "focus": "相遇后等速前进，把判环扩展为找入口",
          "connection": "141 的相遇只证明有环；142 还需定位入口。把一个指针移回 head，再让两者等速前进，用入环前距离与环内路程解释第二次相遇的位置。"
        },
        {
          "id": "234",
          "focus": "无环时定位中点，用递归回退比较两半",
          "connection": "把相遇判环的目标换成 fast 到达链尾，slow 就用于定位中点。本仓库接着利用递归回退比较两半；检查奇数长度为什么需要跳过中间节点。"
        },
        {
          "id": "148",
          "focus": "把中点用于断链，再递归排序和合并",
          "connection": "继续复用定位中点，但目标从镜像比较变成拆分排序。对照 slow、fast 的初始位置和断链位置，确认奇偶长度的两半都会缩小，最后再合并两个有序链表。"
        }
      ]
    },
    {
      "id": "knapsack-state",
      "title": "背包：先问状态表示什么",
      "description": "依次改变可用次数、优化目标和是否区分顺序，用同一容量维度对照状态与遍历方向。",
      "steps": [
        {
          "id": "416",
          "focus": "0-1 背包，dp 表示某个和是否可达"
        },
        {
          "id": "494",
          "focus": "保留每个数只用一次，把可达性改为方案数",
          "connection": "都先把问题化成子集目标和；494 用 P = (sum + target) / 2 转换正负号。状态从布尔值改为计数，每个数仍只用一次，因此对照容量逆序如何避免重复使用当前数。"
        },
        {
          "id": "322",
          "focus": "硬币可以重复，dp 改为最少硬币数",
          "connection": "把每个元素至多使用一次改成硬币可重复使用，容量遍历由逆序变为正序。状态也从方案数改成最少个数，比较不可达状态的初值和 min 更新为什么都要随之改变。"
        },
        {
          "id": "518",
          "focus": "保持可重复使用，改为统计无序组合",
          "connection": "使用次数限制与 322 相同，但最小值改为方案计数。比较 dp[0] 从零枚硬币变为一种空组合的含义，并把硬币放在外层，避免把不同选取顺序重复计数。"
        },
        {
          "id": "377",
          "focus": "顺序也算不同答案，让目标和成为外层循环",
          "connection": "518 把同一组数字的不同顺序视为一个组合；377 要分别计数。把目标和放到外层，再枚举最后一个选的数，用 [1, 2] 凑出 3 对照组合与排列的差别。"
        }
      ]
    },
    {
      "id": "monotonic-next",
      "title": "单调栈：等到更大的那一个",
      "description": "保持出栈条件，分别改变答案的形式、查询方式和数组的边界。",
      "steps": [
        {
          "id": "739",
          "focus": "保存尚未等到更高温度的下标，出栈时填距离"
        },
        {
          "id": "496",
          "focus": "同样找下一个更大值，把结果映射回查询数组",
          "connection": "739 返回下标距离，496 返回更大的元素值，并且只回答 nums1 中的查询。保留遇到更大值时出栈的条件，比较栈里该存什么、答案应如何通过映射关联到查询。"
        },
        {
          "id": "503",
          "focus": "扩展为循环数组，明确重复扫描的停止条件",
          "connection": "保留下一更大元素的出栈逻辑，但数组右端不再是终点。再次经过前缀补齐环形右侧，对照如何限制重复入栈，以及一圈后仍无更大值的元素如何结束等待。"
        }
      ]
    },
    {
      "id": "interval-boundaries",
      "title": "区间：交集、并集与边界",
      "description": "区间题的排序外形相近，真正要比较的是保留什么边界，以及端点相等时是否冲突。",
      "steps": [
        {
          "id": "435",
          "focus": "冲突时保留更早结束的区间，给后面留下空间"
        },
        {
          "id": "452",
          "focus": "让一支箭落在一组区间的公共交集中",
          "connection": "从保留互不重叠的区间，改成用一个位置覆盖一组重叠区间。特别对照端点相等的情形：435 的相接区间可以共存，452 的同一支箭可以命中两个闭区间的公共端点。"
        },
        {
          "id": "763",
          "focus": "用字符出现范围形成区间，合并后确定切分点",
          "connection": "452 关心公共交集，763 则把每个字符的出现范围连成不可拆开的并集。比较边界从取更小的结束位置，变成扩展到更远的结束位置，说明何时才允许切出一个片段。"
        }
      ]
    }
  ]
};
