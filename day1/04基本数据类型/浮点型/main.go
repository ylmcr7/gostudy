package main

import "fmt"

//浮点型
func main()  {
	//math.MaxFloat32 //
	f1 := 1.23456
	fmt.Printf("%T\n",f1)  //默认Go语言中的小数都是float64类型
	f2 := float32(1.23456)
	fmt.Printf("%T",f2)
	f1 = f2 //float32类型的值不能直接赋值给float64
}

