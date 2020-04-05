/*
	字符串:
	1. 概念: 多个byte的集合,理解为一个字符序列
	2. 语法: "" 或 ``
	3. 编码问题:
		ASCII
		Unicode, 保存方式utf8, utf16, utf31 ...
	4. 转义:
		a. 有些字符,有特殊作用,要在字符串里显示他, 可以用\转义为普通字符, 如:
			\"  \`  \\
		b. 有些普通字符, 转义后有特殊作用, 如:
			\n  换行
			\t  制表符
*/
package main

import (
	"fmt"
	"strings"
)

func string() {
	// 字符串
	ss1, ss2 := "A", `
多行文本, 可换行
王二狗`
	// %s 字符串占位符:
	fmt.Printf("%T, %s\n", ss1, ss1) // string, A
	fmt.Printf("%T, %s\n", ss2, ss2)
	// string,
	// 多行文本, 可换行
	// 王二狗

	ss3 := '〇' // 单引号只能用在一个字符
	// %d 字符串编码占位符, %c 或 %q 显示该字符占位符
	fmt.Printf("%T, %d, %c, %q \n", ss3, ss3, ss3, ss3) // int32, 12295, 〇, '〇'

	// 转义:
	// 例如这里 \" 转义掉引号, \\转义掉斜杠 让其不起引号的作用
	fmt.Println("\"c:\\go\"")   // "c:\go"
	fmt.Println("hello\tworld") // hello	world
	fmt.Println(`hello"world"`) // hello"world"
	fmt.Println("hell`owo`rld") // hell`owo`rld

	// 长度: len
	fmt.Println(len(ss2)) // 24

	// 拼接: + 或 Sprintf(字符串模板)
	s3, s4 := "Python", "3.8"
	fmt.Println(s3 + s4)
	s5 := fmt.Sprintf(`%s %s`, s3, s4)
	fmt.Println(s5)

	// 分割: strings.Split
	ret := strings.Split(s5, " ")
	fmt.Println(ret) // [Python 3.8]

	// 包含: strings.Contains
	ret2 := strings.Contains(s5, "3")
	fmt.Println(ret2) // true

	// 判断前后缀:
	ret3 := strings.HasPrefix(s5, "Py") // true
	ret4 := strings.HasSuffix(s5, "8")  // true
	fmt.Println(ret3, ret4)

	// 子串的位置:
	s6 := "applepen"
	fmt.Println(strings.Index(s6, "p"))     // 1
	fmt.Println(strings.LastIndex(s6, "e")) // 6

	// join
	// s7 := []string{"JavaScript", "Python", "PHP", "Ruby", "Golang"}
	// fmt.Println(strings.Join(s7, "~"))
}

func char() {
	s1 := "Golang"
	// '' 表示一个字符,
	var c1 byte = 'G'       // btye = int8, 即查询 ASCII码的位置编号
	var c2 rune = '即'       // rune = int32, 即查询Unicode 的 UTF8, UTF16, UTF32编码的编号
	fmt.Println(s1, c1, c2) // Golang 71 21363

	//

	// 简单应用: 遍历字符串:
	s2 := "hello你好"
	// for i := 0; i < len(s2); i++ {
	// 	fmt.Printf("%c ", s2[i])
	// 	// 如果遍历中英混合的字符串,会出现乱码:
	// 	// h e l l o ä ½   å ¥ ½
	// }

	// for range就按照rune类型去遍历:
	for key, value := range s2 {
		fmt.Printf("%d - %c, ", key, value)
		// 0 - h, 1 - e, 2 - l, 3 - l, 4 - o, 5 - 你, 8 - 好,
	}

	// 修改字符串: 强制类型转换
}

func main() {
	// string()
	char()
}
