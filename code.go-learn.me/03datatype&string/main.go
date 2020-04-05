/*
+ 基本数据类型:
1. 布尔: (bool), 布尔型无法参与数值运算, 也无法与其他类型进行转换
	取值: true, false
2. 数值: (digit)
整形: int
		!!! int类型, 则根据电脑是多少位来决定, 比如64位的电脑, 实际上是int64, 即能显示2^64-1位的数字:
		有符号: 最高位表示符号, 0为正数, 1为负数,其余表示数值:
		int8: (-128 ~ 127)
		int16: (-32768 ~ 32767)
		int32: (-2147483648 ~ 2147483647)
		int64: (-9223372036854775808 ~ 9223372036854775807)
		uint8: (0 ~ 255),
		uint16: (0 ~ 65535),
		uint32: (0 ~ 4294967295)
		uint64: (0 ~ 18446744073709551615)

		byte: uinit8
		rune: int32 当需要处理中文,日文或者其他字符时,需要用到

		进制
		二进制: 占位符: %b
		八进制: 数字以0开头, 占位符: %o
		十六进制: 数字以ox开头, 占位符: %x

		占位符 %p 表示内存地址:
	浮点:
	复数: complex
3. 字符串: (string)
*/
package main

import (
	"fmt"
	"math"
)

func dataType() {
	// 1. 布尔 bool
	b1 := true
	// %t布尔占位符
	fmt.Printf("%T, %t\n", b1, b1)

	// 2. 整数
	var i1 int8 = 100 // 报错
	var i2 uint8 = 0xAD
	fmt.Println(i1)
	fmt.Printf("i2: %x \n", i2)
	fmt.Println(i2) // 173

	var i3 int = 0776
	fmt.Printf("i3: %o \n", i3)
	fmt.Println(i3) // 510
	// var i4 int64
	/* go为强类型语言, 不同种类的数据类型不能相互赋值: */
	// i4 = i3
	// fmt.Println(i4)
	// 报错如下:
	// src\code.go-learn.me\03datatype\main.go:38:5: cannot use i3 (type int) as type int64 in assignment

	// 显示i3的内存地址:
	fmt.Printf("%p \n", &i3)
	// 0xc000012100

	var i5 uint8 = 100
	var i6 byte
	i6 = i5
	var i7 int32 = -2147483648
	var i8 rune
	i8 = i7
	// byte 和 uint8 可以互换, rune 和 int32 可以互换
	fmt.Println(i6 == i5) // true
	// !!!! %d为数值型的占位符
	fmt.Printf("%T, %d\n", i6, i6) // uint8, 100
	fmt.Println(i7 == i8)          // true

	// 浮点:
	var f1 = 3.14
	// !!! 浮点类型占位符: %f, 默认保存后6位, .3表示保留小数点后3位小数
	fmt.Printf("%T, %.3f\n", f1, f1)
	// 直接打印, 写几位就是几位:
	fmt.Println(f1)

	// 各种数常量:
	fmt.Println(math.MaxInt8)    // 127
	fmt.Println(math.MaxInt16)   // 32767
	fmt.Println(math.MaxInt32)   // 2147483647
	fmt.Println(math.MaxInt64)   // 9223372036854775807
	fmt.Println(math.MaxUint8)   // 255
	fmt.Println(math.MaxUint16)  // 65535
	fmt.Println(math.MaxUint32)  // 42949967295
	fmt.Println(math.MaxFloat32) // 3.4028234663852886e+38
	fmt.Println(math.MaxFloat64) // 1.7976931348623157e+308
}

func main() {
	dataType()
}
