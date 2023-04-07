package functions

import (
	"fmt"
	"strings"
)

func IsContinue() bool {
Continue:
	fmt.Printf("Do you want to continue Y/N? ")

	input, _ := reader.ReadString('\n')
	ans := strings.TrimSpace(input)
	fmt.Println(strings.Repeat("-", 60))
	ans = strings.Trim(ans, "\n")
	ans = strings.Trim(ans, " ")
	ans = strings.ToLower(ans)
	switch ans {
	case "yes", "y":
		return true
	case "no", "n":
		return false
	default:
		fmt.Println("Please answer in yes/no")
		fmt.Println(strings.Repeat("-", 60))
		goto Continue
	}

}
