// 1. Start with functions :), why not

// we are now in main package
package main

import "fmt" // we want to import another package: fmt
// fmt is like stdio in c, standart input and output etc...

// bot arguments have type defined, and int as result
func add(x int, y int) int {
	return x + y
}

// if arguments are the same
func add2(x, y int) int {
	return x + y
}

// multiple results returned
func swap(x, y string) (string, string) {
	return y, x
}

// multiple named results
// x and y are named resuts of this function
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return  
}

func main() {
	
	fmt.Println(add(10, 20)) // we are using Println function of package fmt

	fmt.Println(add2(12, 20))

	a, b := swap("Hello", "World")
	fmt.Println(a, b)

	fmt.Println(split(27))

}