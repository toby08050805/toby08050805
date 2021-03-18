package main

import "fmt"

var (
	x      int
	y      int
	z      int
	c      bool
	python bool
	java   bool
	sum    int
)

func main() {

	fmt.Println("Hello, World")
	fmt.Println("哈囉！世界！")
	fmt.Println(x, y, z, c, python, java)
	fmt.Println("1" + "1")
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello" + "World")

	a := "Hello World"
	fmt.Printf("%c \n", a[1])
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("%d   %s", sum, a)
	
}
