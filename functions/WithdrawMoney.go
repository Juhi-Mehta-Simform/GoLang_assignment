package functions

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func WithdrawMoney(u *[]Users) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			WithdrawMoney(u)
		}
	}()
	fmt.Println(strings.Repeat("-", 60))
	fmt.Print("Please enter the amount you want to withdraw: ")
	input, _ := reader.ReadString('\n')
	amount, err := strconv.ParseFloat(strings.TrimSpace(input), 32)
	if err != nil {
		panic("Please enter the valid input. Input must be a Number.")
	}
	if float64(amount) > usr.AccountBalance {
		fmt.Println("Your account does not have enough balance.")
		fmt.Println(strings.Repeat("-", 60))
		if IsContinue() {
			return
		} else {
			os.Exit(0)
		}
	} else if amount >= 0 && amount < 500 {
		fmt.Println("Withdrawal amount shoulde be greater than 500.")
		fmt.Println(strings.Repeat("-", 60))
		if IsContinue() {
			return
		} else {
			os.Exit(0)
		}
	} else if amount > 20000 {
		fmt.Println("Withdrawal amount should be less than 20000.")
		fmt.Println(strings.Repeat("-", 60))
		if IsContinue() {
			return
		} else {
			os.Exit(0)
		}
	} else if amount > 0 {
		amt := math.Mod(amount, 500)
		if amt == 0 {
			usr.AccountBalance -= amount
			fmt.Println(strings.Repeat("-", 60))
			fmt.Printf("Your Current Balance is: %.2f\n", usr.AccountBalance)
			fmt.Println(strings.Repeat("-", 60))
			if IsContinue() {
				return
			} else {
				os.Exit(0)
			}
		} else {
			fmt.Println("Withdrawal amount should be in multiplication of 500")
			fmt.Println(strings.Repeat("-", 60))
			if IsContinue() {
				return
			} else {
				os.Exit(0)
			}
		}
	} else {
		fmt.Println("Please enter the valid input. Input must be a Positive Number.")
		fmt.Println(strings.Repeat("-", 60))
		if IsContinue() {
			return
		} else {
			os.Exit(0)
		}
	}
}
