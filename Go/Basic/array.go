package Basic

import "fmt"

func Array() {
	fmt.Println("Array Learning...")

	programmers := [3]string{}
	// var programmers [10]string

	for i := 0; i < len(programmers); i++ {
		fmt.Printf("Enter a programmer name:-")
		fmt.Scan(&programmers[i])
	}
	fmt.Println(programmers)
}
