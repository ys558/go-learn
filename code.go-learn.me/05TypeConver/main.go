package main

import "fmt"

func main() {
	/*
		go语言是静态语言, 定义, 赋值, 运算必须类型一致
		语法: Type(value)

		常数: 在有需要时,自动转型
		变量: 需要手动转型
	*/
	var a int8 = 10
	// go是强类型语言,不能等于:
	// b = a

	// 应用强制转换:
	var b int16 = int16(a)
	fmt.Println(a, b) //10 10

	f1 := 5.66
	var f2 int = int(f1)
	// 将float转换为int, 只保留整数部分:
	fmt.Println(f1, f2)
	//5.66 5
	f3 := float64(f2)
	fmt.Println(f3)
	// 5

	// b1 := true
	// a := int8(b1)
	// fmt.Println(a)
	// 报错如下: 布尔值和int之间不能转换
	// cannot convert b1 (type bool) to type int8

	// 常数可以自动转型:
	sum := f1 + 100
	fmt.Printf("%T, %f\n", sum, sum)
	// float64, 105.660000

	// 字符串转为字节数组
	s := "big"
	byteArr := []byte(s)
	// fmt.Println(byteArr)
	// [98 105 103]
	byteArr[0] = 'p'
	s = string(byteArr)
	fmt.Println(s)
	// pig

	// 练习, 字符串反转: "hello" -> "olleh"
	s2 := "hello"
	byteArr2 := []byte(s2) // [h e l l o]
	s3 := ""
	for i := len(byteArr2) - 1; i >= 0; i-- {
		// i 为 4 3 2 1 0
		s3 += string(byteArr2[i])
	}
	fmt.Println(s3)
}
