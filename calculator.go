package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mult(a int, b int) int {
	return a * b
}

func main() {
	//Plus
	res := plus(5, 5)
	fmt.Println(res)
	//Sub
	res = sub(7, 5)
	fmt.Println(res)

	//Mult
	res = mult(5, 5)
	fmt.Println(res)

}
