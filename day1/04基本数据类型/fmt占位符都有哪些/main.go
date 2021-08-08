package main

import "fmt"

//fmt占位符
func main()  {
	var n = 100
	//查看类型 %T
	fmt.Printf("%T\n",n) //显示类型
	fmt.Printf("%v\n",n)
	fmt.Printf("%b\n",n)
	fmt.Printf("%d\n",n)
	fmt.Printf("%o\n",n)
	fmt.Printf("%x\n",n)
	var s = "沙河"
	fmt.Printf("字符串:%s\n",s)
	fmt.Printf("字符串:%v\n",s)
	fmt.Printf("字符串:%#v\n",s) //%# 如果你是字符串就会给你加个双引号
}
