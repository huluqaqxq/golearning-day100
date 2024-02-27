package main

import "fmt"

func main() {
	fmt.Println("导入包 必须要使用")

	var b string = "Runoob"
	var c, a = 1, 2
	fmt.Println(b)
	fmt.Println(c, a)

	d := 30
	fmt.Printf("c type is %T\n", d)

	// 多重赋值
	i, j := 0, 1
	fmt.Printf("i = %d j = %d \n", i, j)

	f, _ := 10, 1
	_, e := 10, 2
	fmt.Printf("f = %d, e = %d", f, e)
}
