package Basic

import (
	"fmt"
)

func add(one, two float32) float32 {
	return one + two
}

func sub(one, two float32) float32 {
	return one - two
}

func mul(one, two float32) float32 {
	return one * two
}

func div(one, two float32) float32 {
	return one / two
}

func Switch() {
	fmt.Println("Switch Learning...")
	var number1, number2, result float32

	input := true
	var option string

	for input {
		fmt.Printf("First Number:")
		fmt.Scan(&number1)

		fmt.Printf("Second Number:")
		fmt.Scan(&number2)

		fmt.Printf("Choose an Option: (+, -, *, /)")
		fmt.Scan(&option)

		switch option {
		case "+":
			result = add(number1, number2)
		case "-":
			result = sub(number1, number2)
		case "*":
			result = mul(number1, number2)
		case "/":
			result = div(number1, number2)
		default:
			fmt.Println("Invalid Option")
			continue
		}

		fmt.Printf("Your Result: %v\n", result)
	}

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
