package Basic

import "fmt"

func Switch() {
	fmt.Println("Switch Learning...")

	number := 0

	fmt.Printf("Enter Your Number:")
	fmt.Scanf("%v", &number)

	switch number {
	case 1, 2, 3, 4, 5:
		fmt.Println("You are a primary student")

	case 6, 7, 8, 9, 10:
		fmt.Println("You are a secondary student")
	}
}
