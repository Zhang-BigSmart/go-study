package main

import (
	"fmt"
)

type Books struct {
	title      string
	author     string
	subject    string
	subject_id int
}

func main() {
	demo2()
	//demo3(Books{title: "design pattern"})
}

func demo1() {
	// 创建新的结构体
	fmt.Println(Books{"Golang", "edison", "program", 1})

	// 使用 key:value 格式
	fmt.Println(Books{title: "JVM", author: "zzm", subject: "program", subject_id: 2})

	// 忽略字段为空
	fmt.Println(Books{title: "MySQL", subject: "database"})
}

func demo2() {
	var book Books

	book.title = "高并发"

	fmt.Println(book)

	// var ip *Books
	// ip = &book
	//fmt.Println(user, &user) 使用的是默认格式的打印方式%v。
	//对于 struct 默认输出格式是 {field0 field1 ...}。
	//想要输出地址，需要使用 %p，fmt.Printf("%v %p\n", user, &user)。
	fmt.Printf("%v, %p", book, &book)
}

func demo3(book Books) {
	fmt.Println(book)
}
