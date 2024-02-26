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
}
