package main

import "fmt"

func main() {
	// 变量：程序运行期间，可以改变的量，声明需要 var
	// 常量：程序运行期间，不可以改变的量，常量声明需要 const
	const a int = 10
	// 常量不允许被更改
	//a = 20
	fmt.Printf("a =%d", a)

	const b = 11.2 // 没有使用 :=
	fmt.Printf("b = %f \n", b)

	// 定义多个变量
	var (
		c int
		d float64
	)
	c, d = 10, 3.14
	fmt.Printf("c = %d, d = %.2f \n", c, d)

	// 定义多个常量
	const (
		i int     = 10
		j float64 = 3.14
	)
	fmt.Println("i = ", i)
	fmt.Println("j = ", j)

	// iota
	// 每隔一行子等累加 1
	const (
		x = iota
		y = iota
		z = iota
	)
	fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)

	// 同一行 iota 值是同一个
	const (
		a1     = iota
		b1, c1 = iota, iota
	)
	fmt.Printf("a1 = %d, b1 = %d, c1 = %d", a1, b1, c1)
}
