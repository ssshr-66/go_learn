# Go 工程能力学习

这个仓库用于记录和实践 Go 语言工程能力学习，不以阅读大型开源项目或直接完成复杂业务为起点。

## 学习目标

通过每天约 1.5 小时、每次一个小型可运行 Demo，逐步掌握：

- Go 基础语法和标准工具链
- 结构体、方法、接口、错误处理和代码组织
- 测试、调试、文件 IO 和网络编程
- goroutine、channel、锁、任务队列和资源生命周期
- 超时、重试、优雅退出以及常见线上问题的定位思路

## 学习方式

每个 Demo 按“理解需求 → 自己设计 → 分步实现 → 测试 → 增加需求 → 复盘”的节奏完成。AI 负责讲解、提问、提示和 Review；在用户尝试之前不直接提供完整实现。

算法题暂时继续使用 Java，工程能力训练集中使用 Go，避免在同一个学习时段切换工程语言。

## Trellis 记录方式

- `.trellis/tasks/go-engineering-learning/`：只保存长期学习目标、能力路线和完成标准，供后续对话恢复上下文
- 真正给学习者看的内容都在仓库根目录：每个学习日或小 Demo 建立一个 `go_XX_xxx/` 文件夹，里面放 `README.md`、Go 源码和必要的测试文件
- 不为普通的每周/每日学习建立复杂的 Trellis 子任务树，除非学习者以后明确提出需要
- 每个 Demo 尽量独立、可运行、可测试；学习过程遵循“理解需求 → 自己设计 → 分步实现 → 测试 → 增加需求 → 复盘”

## 当前入口

- 总路线：[LEARNING_PLAN.md](LEARNING_PLAN.md)
- 第一天：[go_01_basic-syntax/README.md](go_01_basic-syntax/README.md)
- 第一天练习代码：[go_01_basic-syntax/main.go](go_01_basic-syntax/main.go)

当前阶段：先用 4 天唤醒 Go 基础语法，再开始第一个小型工程 Demo。

详细的长期目标和阶段划分见 `.trellis/tasks/go-engineering-learning/prd.md`。
