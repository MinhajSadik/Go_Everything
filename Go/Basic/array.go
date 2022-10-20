package Basic

import "fmt"

func Array() {
	fmt.Println("Array Learning...")

	programmers := [10]string{} //fixed value
	// var programmers [10]string

	var programmerName string

	for i := 0; i < len(programmers); i++ {
		fmt.Printf("Enter a programmer name:-")
		fmt.Scan(&programmerName)
	}

	fmt.Println(programmers)
}
