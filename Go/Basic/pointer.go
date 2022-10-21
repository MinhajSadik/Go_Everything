package Basic

import "fmt"

func callByRefrence(ptr *int) {
	*ptr = 131 //3
}

func changeValue(val int) {
	val = 20
}

func Pointer() {
	//(1)it denotes memory address
	//(2)it represents actual value
	//(3)it convey & changed the value by using the memory address
	fmt.Println("Pointer Learning...")
	var pointer *string

	name := "minhaj"
	num := 123

	pointer = &name
	*pointer = "sadik" //3

	fmt.Println(&name)    //1
	fmt.Println(pointer)  //1
	fmt.Println(*pointer) //2

	changeValue(num)

	fmt.Println("change?", num)

	callByRefrence(&num)

	fmt.Println("changed?", num) //3
}
