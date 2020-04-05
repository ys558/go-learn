### Windows子系统Ubuntu安装：
- 1.下载地址：https://golang.google.cn/dl/ 选择Linux版本，为tar包， 默认下载到 /mnt/c/Users/yuyi/Downloads/ 路径：
![clipboard](C:\Users\yuyi\AppData\Local\YNote\data\yuyi.gz@163.com\5b6260852dd949ceb9d5d612789c4243\clipboard.png)

- 打开wsl Ubuntu子系统，将其解压到 /usr/local/ 目录, -C为解压到指定目录， pwd可以查看当前路径：

1. hello world

新建`touch hello.go`，输入如下：

```go
package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}
```
>注意，go语言有大小写之分;

> 其他注意点:
> import时包名可以写为 `import (. "fmt")`, 用.代替 包名fmt


运行方法一:

命令行里`go run hello.go` 可直接运行; `go build hello.go`编译为二进制码，此时生成一个`hello.exe`文件. 

cmd里直接输入exe的文件名`hello`可以直接运行:

```cmd
D:\study\code\go-learn\hello>hello
hello, world
```

git bash或wsl里输入`./hello.exe`运行
```shell
yuyi@□□□ĵ□□□ MINGW64 /d/study/code/go-learn/hello
$ ./hello.exe 
hello, world
```
运行方法二:
在任意路径下运行:
`go install 01hello`
或, 可进入项目文件夹(应用包)的路径,然后运行:
`go install`

指定exe名:
`go build -o haha.exe`

命令源码文件不能放在同一个包

查看go build或go install的详细操作:
`go build -n`

`go install -n`

`go get -x github.com/go-errors/errors`

`C:\Users\yuyi\go\src\github.com\go-errors`