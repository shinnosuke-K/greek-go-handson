package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	// 変数宣言
	var msg string
	msg = "Hello World"
	fmt.Println(msg)

	value := 10
	fmt.Println(value)

	// 関数宣言
	hello()

	// 構造体・型宣言
	user := User{
		Name: "A",
		Age:  10,
	}
	fmt.Println(user)

	// スライス
	num := []int{1, 2, 3, 4, 5}
	fmt.Println(num)
	num = append(num, 6)
	fmt.Println(num)

	// Map
	age := map[string]int{
		"Alice": 10,
		"Bod":   20,
	}
	fmt.Println(age["Alice"])
	fmt.Println(age["C"])

	a, ok := age["Alice"]
	fmt.Println(a, ok)
}

type User struct {
	Name string
	Age  int
}

func hello() {
	fmt.Println("Hello World")
}
