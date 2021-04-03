package main

import (
	"fmt"
	"reflect"
)

func main() {

	demo2()
}

func demo1() {
	str1 := "Golang"
	str2 := "Go语言"
	fmt.Println(reflect.TypeOf(str2[2]).Kind())
	fmt.Println(str1[2], string(str1[2]))
	fmt.Printf("%d %c\n", str2[2], str2[2])
	fmt.Println("len(str2)：", len(str2))
}

func demo2() {
	s := make([]int, 5, 10)
	fmt.Println(len(s), cap(s))
	fmt.Println(s)
	s = append(s, 1, 2, 3, 4, 5, 6)
	fmt.Println(s)
	s = append(s)
	fmt.Println(s)
}
