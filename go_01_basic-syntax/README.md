# Day 1：Go 基础语法唤醒

日期：2026-09-17

今天的目标不是背完 Go 语法，而是重新建立“看得懂、写得出、跑得起来”的感觉。今天先不建复杂工程，也不引入第三方库；我们用一个小命令行程序练习最常用的语法骨架。

## 今天的 90 分钟安排

### 1. 环境确认：0～10 分钟

在当前目录执行：

```bash
go version
go run main.go
```

先观察程序能否运行、输入一个名字后输出什么。今天暂时不要求理解 Go Modules，`go.mod` 会在后面的工程组织阶段引入。

### 2. 语法讲解：10～30 分钟

结合 `main.go`，掌握这些最小知识点：

- `package main`：声明当前文件属于哪个包；可执行程序通常使用 `main` 包。
- `import "fmt"`：引入标准库中的格式化输入输出包。
- `func main()`：程序的入口函数。
- `var name string`：声明一个字符串变量。
- `fmt.Print`、`fmt.Println`：输出内容。
- `fmt.Scan(&name)`：从终端读取输入；`&name` 表示把读到的内容写入变量 `name`。
- `func makeGreeting(name string) string`：声明一个接收字符串、返回字符串的函数。
- `return`：把函数结果交给调用方。

学习时不要只读这份说明。接下来我会在对话中按行解释代码，然后让你合上示例，自己写一个结构相似的方法。

### 3. 控制流程：30～50 分钟

在理解输入输出后，复习并亲手写出一个最小例子：

- `if / else`：根据条件选择分支。
- `for`：Go 中常用的循环语法。
- `switch`：根据多个情况选择分支。

这部分只要求能写出简单条件，不要求今天掌握所有变体。

### 4. 函数仿写：50～70 分钟

先阅读 `makeGreeting`，然后完成下面的练习。不要先让 AI 直接写答案：

在 `main.go` 中实现 `makeFarewell(name string) string`：

1. 接收一个名字。
2. 返回一条包含名字的告别语句，例如 `再见，小明！`。
3. 在 `main` 中调用它，让程序同时打印问候语和告别语。

如果卡住，只先问我一个具体问题，例如“Go 的字符串怎么拼接”或“为什么这里要写 `return`”。

### 5. 运行与复盘：70～90 分钟

完成后执行：

```bash
gofmt -w main.go
go run main.go
```

然后检查：

- 输入名字后，程序是否输出两条符合要求的语句？
- 你能否解释 `main`、`makeGreeting`、`makeFarewell` 的关系？
- 你能否说出 `fmt.Scan(&name)` 中 `&` 的作用？
- 你今天遇到的第一个编译错误是什么，最后如何修复？

## 今天的完成标准

- [ ] 能运行 `go version` 和 `go run main.go`
- [ ] 能解释 `package`、`import`、`func`、`var`、`return`
- [ ] 独立完成 `makeFarewell`
- [ ] 至少写过一个 `if` 或 `for` 小例子
- [ ] 运行过 `gofmt`
- [ ] 用自己的话向我复述一次代码执行流程

## 下一检查点

完成 `makeFarewell` 后，把你的 `main.go` 内容和运行结果发给我。我会先 Review 你的写法，再给第二个小要求；只有完成当天检查点，才进入 Day 2 的数组、切片和 `map`。

总路线见 [LEARNING_PLAN.md](../LEARNING_PLAN.md)。
