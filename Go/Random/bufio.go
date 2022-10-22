package Random

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Bufio() {
	fmt.Println("Input output using go pkg module...")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please give us a rating between 1 and 5:")

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for watching", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("Added 1 to your rating:", numRating+1)
	}
}
