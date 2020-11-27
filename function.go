package main

import "fmt"

func plus(a int, b int) int {
	return a+b
}

func plusplus(a ...int) int {
	total := 0
	for _,num := range a{
		total += num
	}
	return total
}

func vals(a,b int)(int, int){
	return plus(a, b*2), plus(b*a, 2)
}

func main() {
	res := plus(1,2)
	fmt.Println(res)

	res = plusplus(1,2,3,4,5,6)
	fmt.Println(res)
}