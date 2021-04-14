package main

import (
	"fmt"
)

func main() {
	demo3()
}

func demo1() {
	// 使用 make 函数创建切片
	//var slice1 []type = make([]type, len)
	// 简写
	// slice1 := make([]type, len)

}

func demo2() {
	myNUm := []int{10, 20, 30, 40, 50}
	// 创建一个新的切片, 长度为2，容量为4（共享数组地址）
	newNum := myNUm[1:3]
	fmt.Println(myNUm[1])
	fmt.Println(newNum[0])

	myNUm[1] = 100
	fmt.Println(myNUm[1])
	fmt.Println(newNum[0])

	newNum[0] = 111
	fmt.Println(myNUm[1])
	fmt.Println(newNum[0])

}

func demo3() {
	// 切片扩容
	myNum := []int{10, 20, 30, 40, 50}
	// 创建一个新的切片, 长度为2，容量为4（共享数组地址）
	newNum := myNum[1:3]

	fmt.Println(myNum[3])

	// 将新元素赋值为60
	newNum = append(newNum, 60)
	fmt.Println(myNum[3])
	fmt.Println(newNum[2])

	// 切片数组如果容量不足，append()函数会创建一个容量为原来两倍的数组
	newNum = append(myNum, 80)
	fmt.Println(cap(newNum))

	//函数 append() 会智能地处理底层数组的容量增长。在切片的容量小于 1000 个元素时，总是会成倍地增加容量。
	//一旦元素个数超过 1000，容量的增长因子会设为 1.25，也就是会每次增加 25%的容量(随着语言的演化，这种增长算法可能会有所改变)。
}
