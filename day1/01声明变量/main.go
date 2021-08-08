package main

import "fmt"

// Go语言推荐使用驼峰命名
//声明变量
//var name string
//var age int
//var isOk bool

//批量声明
var (
	name string //""
	age int     // 0
	isOk bool   // false
)

func main()  {
	name = "理想"
	age = 16
	isOk = true
	//var heiheiehei string  声明不使用无法编译
	//Go语言中变量声明必须使用 不使用就编译不下去
	fmt.Print(isOk) //在终端输出要打印的内容
	fmt.Println()
	fmt.Printf("name:%s\n",name) //%s:占位符 使用name这个变量的值去替换占位符
	fmt.Println(age) //打印完指定的内容之后会在后面加一个换行符
	var s1 string = "whm"
	fmt.Println(s1)
	var s2 = "20"
	fmt.Println(s2)
	//简短变量声明
	s3 := "哈哈哈"  //不能在函数外面使用
	fmt.Println(s3)
	// s1 := "10" //同一个作用域中不能重复声明同名的变量

	//匿名变量是一个特殊的变量
}


