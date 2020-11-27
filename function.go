package main

import "fmt"

func plus(a int, b int) int {
	return a+b
}

func plusplus(a, b, c int) int {
	return a+b+c
}

func vals(a,b int)(int, int){
	return plus(a, b*2), plus(b*a, 2)
}

func main() {
	res := plus(1,2)
	fmt.Println(res)

	res = plusplus(1,2,3,)
	fmt.Println(res)
}