# Go 语言数据类型
***
1. 布尔型
> true or false 列子： var b bool = ture.
2. 数字类型
> 整型 int 和 浮点型 float32、float64, GO
3. 字符串类型
> 字符串就是一串固定长度的字符连起来的字符序列。GO 语言的字节使用 UTF-8 编码标识 Unicode 文本
4. 派生类型
> - 指针类型（Pointer）
> - 数组类型
> - 结构化类型
> - Channel 类型
> - 函数类型
> - 切片类型
> - 接口类型 （interface）
> - Map 类型
***
# 变量的声明
1. 声明格式 var 变量名 类型,变量声明了必须使用
2. 只是声明没有初始化的变量,默认值为0
3. 同一个 {} 里, 声明的变量名是唯一的
4. 可以同时声明多个变量
```go
 var b int 10 b=20
```
***
# 变量的赋值
1. 变量的初始化，声明变量时，同时赋值
```go
var b int 10 b = 20
fmt.Println("b =", b)
```
2. 自动推导类型，必须初始化，通过初始化的值确定类型
```go
c := 30
fmt.Printf("c type is %T\n", c)
```
3. 赋值前必须声明变量类型
```go
a = 10
fmt.Printf("a type is %T\n", a)
```
***
# Printf 和 Println 的区别
```go
a := 10
b := 20
fmt.Println("a = ", a, "b = ", b)
fmt.Printf("a = %d, b = %d", a, b)
```
***
# 匿名变量
```go
f, _ := 10, 1
_, e := 10, 2
fmt.Printf("f = %d, e = %d", f, e)
```
