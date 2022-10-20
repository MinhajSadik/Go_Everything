package Basic

import "fmt"

func Slice() {
	fmt.Println("Slice Learning...")
	programmers := []string{}
	// var programmers []string

	var numOfPro int
	var programmerName string

	fmt.Printf("How many students:?")
	fmt.Scan(&numOfPro)

	for i := 0; i < numOfPro; i++ {
		fmt.Printf("Enter a programmer name:-")
		fmt.Scan(&programmerName)
		programmers = append(programmers, programmerName)
	}

	fmt.Println(programmers)

	for index, programmer := range programmers {
		fmt.Println(index)
		fmt.Println(programmer)
	}
}
