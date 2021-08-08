package main

import (
	"fmt"
	"strings"
)
func main()  {
	name := "理想"
	age := "18"
	ss := name + age
	fmt.Printf(ss)

	ss1 := fmt.Sprintf("\n%s%s\n",name,age)
	fmt.Printf(ss1)
	fmt.Println(len(ss))
	path := `D:\\Go\\src\\code\\study`
	//分割
	ret := strings.Split(path,"\\")
	fmt.Println(ret)
	//包含
	fmt.Println(strings.Contains(ss,"理想"))
	//前缀: 判断是否以理想开头的
	fmt.Println(strings.HasPrefix(ss,"理想"))
	//后缀 判断是否以18 结尾的
	fmt.Println(strings.HasSuffix(ss,"18"))
	//判断位置
	s4 := "abcdef"
	fmt.Println(strings.Index(s4,"c"))
	fmt.Println(strings.LastIndex(s4,"b"))
	//拼接
	fmt.Println(strings.Join(ret,"+"))
}
