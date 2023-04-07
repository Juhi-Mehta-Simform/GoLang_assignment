package functions

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func DepositMoney(u *[]Users) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			DepositMoney(u)
		}
	}()
	fmt.Println(strings.Repeat("-", 60))
	fmt.Print("Please enter the amount you want to deposit: ")
	input, _ := reader.ReadString('\n')
	amount, err := strconv.ParseFloat(strings.TrimSpace(input), 32)
	if err != nil {
		panic("Please enter the input amount. The input should be a number.")
	}
	if amount < 0 {
		fmt.Println("Please enter the valid input. Input must be a Positive Number.")
		fmt.Println(strings.Repeat("-", 60))
		if IsContinue() {
			return
		} else {
			os.Exit(0)
		}
	} else if amount >= 0 && amount < 500 {
		fmt.Println("Deposit amount shoulde be greater than 500.")
		fmt.Println(strings.Repeat("-", 60))
		if IsContinue() {
			return
		} else {
			os.Exit(0)
		}
	} else if amount > 20000 {
		fmt.Println("Deposit amount should be less than 20000.")
		fmt.Println(strings.Repeat("-", 60))
		if IsContinue() {
			return
		} else {
			os.Exit(0)
		}
	} else {
		amt := math.Mod(amount, 500)
		if amt == 0 {
			fmt.Println(strings.Repeat("-", 60))
			fmt.Println("Your current balance is: ", usr.AccountBalance)
			usr.AccountBalance += amount
			fmt.Println("Yo have successfully deposited amount: ", amount)
			fmt.Printf("Your new balance is: %.2f\n", usr.AccountBalance)
			fmt.Println(strings.Repeat("-", 60))
			if IsContinue() {
				return
			} else {
				os.Exit(0)
			}
		} else {
			fmt.Println("Deposit amount should be in multiplication of 500")
			fmt.Println(strings.Repeat("-", 60))
			if IsContinue() {
				return
			} else {
				os.Exit(0)
			}
		}

	}
}
