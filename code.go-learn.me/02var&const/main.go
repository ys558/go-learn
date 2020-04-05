package main

import "fmt"

// 1. 函数外面定义的变量称为包内变量, go没有全局变量的概念:
// var aa = 3
// var bb = "kkk"
// 也可简写为:
var (
	aa = 3
	bb = "kkk"
)

func var_base() {
	var a int
	var s string
	// 2. Printf表示打印格式, 类似于js的字符串模板``,
	// Println表示只打印能值
	// %q是quotation的意思, 把双引号""打印出来
	/*
		!!! 注意, go语言的变量声明了一定要使用, 否则会报错:
	*/
	fmt.Printf("%d %q \n", a, s)
}

func var_init_value() {
	var s string = "abc"
	var i int = 99
	// 3.同时定义多个变量:
	var a, b int = 3, 4
	fmt.Println(i, s, a, b)
}

func var_type_deduction() {
	// 4. 不指定类型定义变量:
	var u = "a"
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(u, a, b, c, s)
}

func var_shortter() {
	// 5. := 定义一个变量,省略前面的var, 只能用于函数内部,不能定义于包变量
	a, b, c, s := 3, 4, false, "l"
	// 第二次赋值b, 不可再用:=, 直接b=88
	b = 88
	fmt.Println(a, b, c, s)
}

// go语言变量可用返回多个返回值, 如下定义:
func var_return_multi_value() (string, int) {
	return "YY", 29
}

// 常量常定义于包变量内:
const e = 2.7182
const (
	filename = "a.txt"
	kk, ee   = 45, 12
)

// 常量的简写: 这里dd也为1:
const (
	cc = 1111
	dd
)

// 常量枚举类型 iota : 每增加一行,iota依次增加1
const (
	zero = iota
	one
	_ //跳过一个值
	two
)
const (
	qq, ww = iota + 1, iota + 2 // 1 2
	ii, rr
	tt, yy
)
const (
	_  = iota           //0
	Kb = 1 << 10 * iota // 二进制数1往左移10位, 即2的10次方, 1024
	Mb = 1 << 10 * iota // 二进制数1往左移20位, 即2的20次方, 2048
	Gb = 1 << 10 * iota
	Tb = 1 << 10 * iota
	Pb = 1 << 10 * iota
)

func main() {
	// 变量:
	fmt.Println(aa, bb)
	var_base()
	var_init_value()
	var_type_deduction()
	var_shortter()
	// 调用 var_return_multi_value()里的赋值:
	// 用变量a赋值给"YY",
	// 不想用上面29, 用_代替, 称匿名变量
	a, _ := var_return_multi_value()
	fmt.Println(a)
	_, b := var_return_multi_value()
	fmt.Println(b)

	fmt.Println(e, filename, kk, ee) //2.7182 a.txt 45 12
	fmt.Println(cc, dd)              //1111 1111
	fmt.Println(zero, one, two)      //0 1 3
	fmt.Println(Kb, Mb, Gb, Tb, Pb)  //1024 2048 3072 4096 5120

	fmt.Println(qq, ww, ii, rr, tt, yy)
}
