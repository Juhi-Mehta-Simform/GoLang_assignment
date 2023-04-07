package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Users struct {
	Username       string
	CardNo         int
	PinNo          int
	AccountType    string
	AccountBalance float64
}

var usr Users
var reader = bufio.NewReader(os.Stdin)

func UserDetails(u *[]Users) {
CardNo:
	fmt.Println(strings.Repeat("-", 60))
	fmt.Print("Please enter your Card Number: ")
	input, _ := reader.ReadString('\n')
	cardNo, err := strconv.Atoi(strings.TrimSpace(input))
	input = strings.Trim(input, "\n")
	input = strings.Trim(input, " ")
	x := len(input)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			UserDetails(u)
		}
	}()
	if err != nil {
		panic("Please enter valid input. Enter digits only")
	}
	if x != 8 {
		fmt.Println("Please enter 8 digit valid card number")
		goto CardNo
	}
	for _, value := range *u {
		if cardNo == value.CardNo {
			usr = value
			// fmt.Println(value)
			break
		}
	}
	if (usr == Users{}) {
		fmt.Println("User not found")
		fmt.Println(strings.Repeat("-", 60))
		os.Exit(0)
	}
	attemps := 3
PinNo:
	fmt.Println(strings.Repeat("-", 60))
	fmt.Print("Enter your Pin Number: ")
	input, _ = reader.ReadString('\n')
	pinNo, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Please enter valid input. Enter digits only")
		goto PinNo
	}
	input = strings.Trim(input, "\n")
	input = strings.Trim(input, " ")
	x = len(input)
	if pinNo != usr.PinNo {
		attemps--
		if attemps == 0 {
			fmt.Println("You have exceeded the maximum number of attempts. Please try again after sometime.")
			fmt.Println(strings.Repeat("-", 60))
			os.Exit(0)
		} else if x != 4 {
			fmt.Println("Please enter the valid 4 digit Pin Number")
			fmt.Printf("%d attempts left\n", attemps)
			goto PinNo
		} else {
			fmt.Println("Please enter the correct Pin Number")
			fmt.Printf("%d attempts left\n", attemps)
			goto PinNo
		}
	} else {
		fmt.Println("You are successfully logged in.")
	}

}
