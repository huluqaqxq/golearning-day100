#  常量的使用
***
> 常量是一个简单值的标识符，在程序运行时，不会被修改的量 \
> 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型、复数） 和 字符串

> 变量：程序运行期间，可以改变的量，声明需要 var 
> 常量：程序运行期间，不可以改变的量，常量声明需要 const
> 
***
# iota
> iota, 特殊的常量，可以认为是一个可以被编译器修改的常量
> iota 在 const 关键字出现时将被重置为 0（const 内部的第一行之前）, const 中每新增一行常量声明将使 iota 计数一次。
```go
const (
	a = iota
	b = iota
	c = iota
)

const d = iota
fmt.Print("a = %d, b = %d,c = %d \n d = %d", a, b, c , d)
```
> iota 同一行 为 同样的值
```go
const (
	a = iota
	b, c = iota, iota
)
fmt.Printf("a = %d, b = %d, c = %d", a, b, c)
```
