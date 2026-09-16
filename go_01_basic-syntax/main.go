package main

import "fmt"

func main() {
	var name string

	fmt.Print("请输入名字：")
	fmt.Scan(&name)

	fmt.Println(makeGreeting(name))
	fmt.Println(makeFarewell(name))
}

func makeGreeting(name string) string {
	return "你好，" + name + "！"
}

func makeFarewell(name string) string {
	// TODO: Day 1 练习：返回一条包含 name 的告别语句。
	return ""
}
