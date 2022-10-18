package Basic

import "fmt"

func callByRefrence(ptr *int, val int) {
	*ptr = val
}

func Pointer() {
	var ptr *int

	x := 10
	newX := 32
	y := 2
	ptr = &y

	*ptr = *ptr * 2

	callByRefrence(&x, newX)
	fmt.Println(x, y)
}
