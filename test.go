package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	demo8()
	//fmt.Println(demo6(6, 2))
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

func demo3() {

	//双向型 chan, 零值 nil
	var ch chan int
	//输入型 chan->
	var ci chan<- int
	//输出型 <-chan
	var co <-chan int
	//make()
	cc := make(chan int, 10)
	ch = cc
	//双向型赋值给单向型正确
	ci = ch
	co = ch
	//单向型赋值给双向型错误
	//ch = ci //❌
	//ch = co //❌
	fmt.Println(ci, co, ch) //0xc00008c000 0xc00008c000 0xc00008c000
}

func demo4() {
	var i = 3
	go func(a int) {
		fmt.Println(a)
		fmt.Println(1)
	}(i)

	fmt.Println("2")
	time.Sleep(1 * time.Second)
}

func demo5() {
	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}

func demo6(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func demo7() {
	a := 100
	b := 200

	fmt.Println("before:\n", a, b)

	swap(&a, &b)

	fmt.Println("after:\n", a, b)

}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func demo8() {
	var a int = 10
	fmt.Printf("a 变量地址 : %x\n", &a)

	var ip *int
	ip = &a
	fmt.Printf("ip 变量存储的指针地址 : %x\n", ip)

	fmt.Printf("*ip 变量的值 : %d\n", *ip)
}
