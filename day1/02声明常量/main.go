package main

import "fmt"

//常量
const pi = 3.1415926

//批量声明常量
const (
	statusOK = 200
	notFOUND = 404
)

//批量声明常量时,如果某一行没有赋值 那么值和前一行一致
const (
	n1 = 100
	n2
	n3
)

//iota   当const 出现 iota一定被置为0
const (
	a1 = iota //0
	a2 = iota //1
	a3 = iota //3
)

const (
	b1 = iota //
	b2        //当我的常量里
	_
	b3
)

const (
	_  = iota
	KB = 1 <<(10 * iota)
	MB = 1 <<(10 * iota)
	GB = 1 <<(10 * iota)
	TB = 1 <<(10 * iota)
	PB = 1 <<(10 * iota)
)

func main()  {
	//pi = 123
	//fmt.Println("n1",n1)
	//fmt.Println("n1",n2)
	//fmt.Println("n1",n3)
	//
	//fmt.Println("a1",a1)
	//fmt.Println("a2",a2)
	//fmt.Println("a3",a3)

	//fmt.Println("b1",b1)
	//fmt.Println("b2",b2)
	//fmt.Println("b3",b3)

	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)

}

